terraform {
  required_providers {
    ntopng = {
      source  = "myoung34/ntopng"
      version = "0.0.1"
    }
  }
}

provider "ntopng" {
  host  = "http://test.foo.com"
  token = "234444444444433333333333222222221111111"
}

resource "ntopng_user" "test" {
  username               = "testuser"
  full_name              = "Test User"
  password               = "securepassword123"
  user_role              = "unprivileged"        # Or "administrator"
  allowed_networks       = ["0.0.0.0/0", "::/0"] # Example networks
  allowed_interface      = ""                    # Leave empty for all interfaces
  user_language          = "en"                  # Example: en, it, de, jp, pt, cz
  allow_pcap_download    = true
  allow_historical_flows = true
  allow_alerts           = true
}

output "new_user_id" {
  value = resource.ntopng_user.test.username
}
