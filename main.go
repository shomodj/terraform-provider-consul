package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-provider-consul/consul"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: consul.Provider})
}
