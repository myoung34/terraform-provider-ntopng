package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"terraform-provider-ntopng/ntopng"
)

// Generate the Terraform provider documentation using `tfplugindocs`:
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	err := providerserver.Serve(
		context.Background(),
		ntopng.New,
		providerserver.ServeOpts{
			Address: "github.com/myoung34/ntopng",
		},
	)

	if err != nil {
		panic(fmt.Sprintf("Error serving provider: %v", err))
	}
}
