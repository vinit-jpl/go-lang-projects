package main

import (
	"context"
	"log"
	"terraform-provider-mypg/internal/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	err := providerserver.Serve(context.Background(), provider.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/local/mypg",
	})

	if err != nil {
		log.Fatal(err.Error())
	}
}
