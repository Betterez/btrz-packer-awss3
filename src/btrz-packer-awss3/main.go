package main

import (
	"btrz/provisioner"
	"github.com/hashicorp/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(new(provisioner.S3Loader))
	server.Serve()
}
