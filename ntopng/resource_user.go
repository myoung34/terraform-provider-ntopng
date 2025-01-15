package ntopng

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/myoung34/terraform-plugin-framework-utils/modifiers"
)

func UserResource() resource.Resource {
	return &userResource{}
}

type userResource struct {
	client *Client
}

func (r *userResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (r *userResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*Client)
}

func (r *userResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				Required: true,
			},
			"full_name": schema.StringAttribute{
				Optional: true,
			},
			"password": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
			},
			"user_role": schema.StringAttribute{
				Required: true,
			},
			"allowed_networks": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
			},
			"allowed_interface": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"user_language": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"allow_pcap_download": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
					modifiers.DefaultBool(false),
				},
			},
			"allow_historical_flows": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
					modifiers.DefaultBool(false),
				},
			},
			"allow_alerts": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
					modifiers.DefaultBool(false),
				},
			},
		},
	}
}

func (r *userResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// not implemented yet
}

func (r *userResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan User
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	userResp, err := r.client.CreateUser(User{
		Username:             plan.Username,
		FullName:             plan.FullName,
		Password:             plan.Password,
		UserRole:             plan.UserRole,
		AllowedInterface:     plan.AllowedInterface,
		UserLanguage:         plan.UserLanguage,
		AllowPCAPDownload:    plan.AllowPCAPDownload,
		AllowHistoricalFlows: plan.AllowHistoricalFlows,
		AllowAlerts:          plan.AllowAlerts,
		//TODO AllowedNetworks:      toStringSlice(plan.AllowedNetworks.Elements()),
	})

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating",
			"Could not create user, unexpected error: "+err.Error(),
		)
		return
	}

	var result = User{
		Username:             userResp.Username,
		FullName:             userResp.FullName,
		UserRole:             userResp.UserRole,
		Password:             userResp.Password,
		AllowedInterface:     userResp.AllowedInterface,
		UserLanguage:         userResp.UserLanguage,
		AllowPCAPDownload:    userResp.AllowPCAPDownload,
		AllowHistoricalFlows: userResp.AllowHistoricalFlows,
		AllowAlerts:          userResp.AllowAlerts,
		//AllowedNetworks: types.ListValueMust(
		//	types.StringType,
		//	toStringSlice(userResp.AllowedNetworks),
		//),

		//TODO AllowedNetworks:   makeListStringAttributeFn(userResp.AllowedNetworks, func(g ntopng.ObjectGroup) (string, bool) { return g.ObjectGroupID, true }),
	}

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *userResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *userResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *userResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
