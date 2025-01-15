package ntopng

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"os"
)

func New() provider.Provider {
	return &NtopngProvider{}
}

type NtopngProvider struct{} //revive:disable-line:exported

type ntopngProviderData struct {
	Host  types.String `tfsdk:"host"`
	Token types.String `tfsdk:"token"`
}

func (p *NtopngProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "ntopng"
}

func (p *NtopngProvider) Schema(_ context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host":  schema.StringAttribute{Optional: true},
			"token": schema.StringAttribute{Optional: true},
		},
	}
}

func (p *NtopngProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Ntopng client")

	// Retrieve provider data from configuration
	var config ntopngProviderData
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var token string
	if config.Token.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as token",
		)
		return
	}

	if config.Token.IsNull() {
		token = os.Getenv("NTOPNG_TOKEN")
	} else {
		token = config.Token.ValueString()
	}

	if token == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find api key",
			"token cannot be an empty string",
		)
		return
	}

	// User must specify a host
	var host string
	if config.Host.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Cannot use unknown value as host",
		)
		return
	}

	if config.Host.IsNull() {
		host = os.Getenv("NTOPNG_HOST")
	} else {
		host = config.Host.ValueString()
	}

	if host == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find host",
			"Host cannot be an empty string",
		)
		return
	}

	// Create a new ntopng client and set it to the provider.client
	client, err := NewClient(Config{
		Host:  host,
		Token: token,
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Unable to create ntopng client:\n\n"+err.Error(),
		)
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client

	tflog.Info(ctx, "Configured Ntopng client", map[string]any{"success": true})

}

func (p *NtopngProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		UserResource,
	}
}

func (p *NtopngProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}
