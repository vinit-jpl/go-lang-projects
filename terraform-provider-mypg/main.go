package main

import (
	"context"
	"terraform-provider-mypg/internal/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	providerserver.Serve(context.Background(), provider.New, providerserver.ServeOpts{
		Address: "example/mypg",
	})
}
