package ntopng

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func makeProviderFactoryMap(name string, prov *NtopngProvider) map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		name: providerserver.NewProtocol6WithError(prov),
	}
}

func BasicTestNtopng(t *testing.T) {
	const testConfig = // language=hcl
	`
provider "ntopng" {
  host  = "https://test.example.com"
  token = "234444444444433333333333222222221111111"
}

resource "ntopng_user" "test" {
  username               = "testuser"
  full_name              = "Test User"
  password               = "securepassword123"
  user_role              = "unprivileged" # Or "administrator"
  allowed_networks       = ["0.0.0.0/0", "::/0"] # Example networks
  allowed_interface      = "" # Leave empty for all interfaces
  user_language          = "en" # Example: en, it, de, jp, pt, cz
  allow_pcap_download    = 1 # 1 to allow, 0 to deny
  allow_historical_flows = 1 # 1 to allow, 0 to deny
  allow_alerts           = 1 # 1 to allow, 0 to deny
}
`

	var (
		prov = new(NtopngProvider)
	)

	resource.Test(
		t,
		resource.TestCase{
			ProtoV6ProviderFactories: makeProviderFactoryMap("ntopng", prov),
			Steps: []resource.TestStep{
				{
					Config: testConfig,
				},
			},
		},
	)
}
