package main

import (
	client "mysql/myclient"
)

type Test struct {
	ID     int `gorm:"bigint unsigned;primaryKey;not null" column:"id"`
	Val    int `gorm:"int unsigned;index;not null" column:"val"`
	Source int `gorm:"int unsigned;not null" column:"source"`
}

func (t *Test) TableName() string {
	return "test"
}

// 测试limit如何不回表,走聚簇索引
func main() {
	// 给定测试数据给到mysql
	cli := client.NewMySQLClient()

	// Create column
	//err := cli.AutoMigrate(&Test{})
	//if err != nil {
	//	panic(err)
	//}

	// Create Virtual Data
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 10000; j++ {
			cli.Create(&Test{
				Val:    i,
				Source: j,
			})
			if cli.Error != nil {
				panic(cli.Error)
			}
		}
	}
}
