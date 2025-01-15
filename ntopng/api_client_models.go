package ntopng

type APIUser struct {
	Username             string   `json:"username"`
	FullName             string   `json:"full_name,omitempty"`
	Password             string   `json:"password,omitempty"`
	ConfirmPassword      string   `json:"confirm_password,omitempty"`
	UserRole             string   `json:"user_role,omitempty"`
	AllowedNetworks      []string `json:"allowed_networks,omitempty"`
	AllowedInterface     string   `json:"allowed_interface,omitempty"`
	UserLanguage         string   `json:"user_language,omitempty"`
	AllowPCAPDownload    bool     `json:"allow_pcap_download,omitempty"`
	AllowHistoricalFlows bool     `json:"allow_historical_flows,omitempty"`
	AllowAlerts          bool     `json:"allow_alerts,omitempty"`
}
