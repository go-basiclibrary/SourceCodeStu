package main

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)

const goodsMapping = `
{
	"mappings":{
		"properties":{
			"name":{
				"type":"text",
				"analyzer":"ik_max_word"
			},
			"id":{
				"type":"integer"
			}
		}
	}
}`

type User struct {
	AccountNumber int    `json:"account_number"`
	Balance       int    `json:"balance"`
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	Age           int    `json:"age"`
	Gender        string `json:"gender"`
	Address       string `json:"address"`
	Employer      string `json:"employer"`
	Email         string `json:"email"`
	City          string `json:"city"`
	State         string `json:"state"`
}

// es 测试
func main() {
	host := "http://43.143.172.37:9200"
	logger := log.New(os.Stdout, "mall", log.LstdFlags)
	// sniff 会将上面的http地址自动转换内网地址或者docker中的ip地址,导致连接不上
	c, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false),
		elastic.SetTraceLog(logger))
	if err != nil {
		panic(err)
	}

	// match方法获取数据
	q := elastic.NewMatchQuery("address", "street")
	res, err := c.Search().Index("user").Query(q).Do(context.Background())
	if err != nil {
		panic(err)
	}
	for _, v := range res.Hits.Hits {
		data, err := v.Source.MarshalJSON()
		if err != nil {
			panic(err)
		}
		var user User
		err = json.Unmarshal(data, &user)
		if err != nil {
			panic(err)
		}
	}

	// 添加数据到es
	//user := User{
	//	AccountNumber: 15468,
	//	Balance:       669933,
	//	Firstname:     "WangShao",
	//	Lastname:      "ShaoWang",
	//	Age:           18,
	//	Gender:        "F",
	//	Address:       "BaoAn Center",
	//	Employer:      "Bezal",
	//	Email:         "123@qq.com",
	//	City:          "Shenzhen",
	//	State:         "MO",
	//}
	//put, err := c.Index().Index("my_user").BodyJson(user).Do(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Index %v Id %v Type %v\n", put.Index, put.Id, put.Type)

	// 添加索引 index - mapping
	//ctI, err := c.CreateIndex("mygoods").BodyString(goodsMapping).Do(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//if !ctI.Acknowledged {
	//	fmt.Println("创建出错")
	//}
}
