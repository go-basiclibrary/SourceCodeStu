package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func Register(address string, port int, name string, tags []string, id string) {
	cfg := api.DefaultConfig()
	addr := fmt.Sprintf("%s:%d", "121.37.232.8", 8500)
	cfg.Address = addr

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// 生成注册对象
	client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		Name:    name,
		ID:      id,
		Port:    port,
		Address: addr,
		Tags:    tags,
	})
}
