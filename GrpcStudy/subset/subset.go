package main

import (
	"fmt"
	"math/rand"
)

func main() {
	res := subset([]string{"172.25.96.1", "172.25.96.2", "172.25.96.3", "172.25.96.4"}, 1, 2)
	tes := subset([]string{"172.25.96.1", "172.25.96.2", "172.25.96.3", "172.25.96.4"}, 1, 2)
	fmt.Println(res)
	fmt.Println(tes)
}

// backends 全量ip
// subsetSize 期望子集数,比方我想获取50个ip
// rand使用洗牌算法,
// clientId为原client的Ip地址做了CRC处理
func subset(backends []string, clientId int, subsetSize int) []string {
	// 均分后子集的数量
	subSetCount := len(backends) / subsetSize

	// 获取随机round
	round := clientId / subSetCount
	rand.Seed(int64(round))
	rand.Shuffle(len(backends), func(i, j int) {
		backends[i], backends[j] = backends[j], backends[i]
	})

	subsetId := clientId % subSetCount
	start := subsetId * subsetSize
	return backends[start : start+subsetSize]
}
