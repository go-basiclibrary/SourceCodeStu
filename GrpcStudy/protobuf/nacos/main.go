package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"time"
)

// nacos服务测试
func main() {
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "121.37.232.8",
			Port:   8848,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         "23931cd2-a974-493e-9187-29f1e38f46ad",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}

	cfg, err := client.GetConfig(vo.ConfigParam{
		DataId: "user-web.json",
		Group:  "dev",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)

	//监听
	err = client.ListenConfig(vo.ConfigParam{
		DataId: "user-web.yaml",
		Group:  "dev",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println(namespace)
			fmt.Println(group)
			fmt.Println(dataId)
			fmt.Println(data)
		},
	})
	if err != nil {
		panic(err)
	}

	time.Sleep(3000e9)
}
