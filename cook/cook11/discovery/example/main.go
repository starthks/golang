package main

import (
	"cook/cook11/discovery"

	consul "github.com/hashicorp/consul/api"
)

func main() {
	config := consul.DefaultConfig()
	config.Address = "localhost:8500"
	cli, err := discovery.NewClient(config, "localhost", "discovery", 8080)
	if err != nil {
		panic(err)
	}
	if err := discovery.Exec(cli); err != nil {
		panic(err)
	}
}
