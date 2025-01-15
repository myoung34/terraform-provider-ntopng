package ntopng

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type User struct {
	Username             types.String `tfsdk:"username"`
	FullName             types.String `tfsdk:"full_name"`
	Password             types.String `tfsdk:"password"`
	UserRole             types.String `tfsdk:"user_role"`
	AllowedNetworks      types.List   `tfsdk:"allowed_networks"` // List of strings
	AllowedInterface     types.String `tfsdk:"allowed_interface"`
	UserLanguage         types.String `tfsdk:"user_language"`
	AllowPCAPDownload    types.Int64  `tfsdk:"allow_pcap_download"`
	AllowHistoricalFlows types.Int64  `tfsdk:"allow_historical_flows"`
	AllowAlerts          types.Int64  `tfsdk:"allow_alerts"`
}
