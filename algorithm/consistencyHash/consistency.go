package main

import (
	"hash/crc32"
)

func main() {
}

// 一致性hash算法实现

// CustomHash 自定义哈希算法
type CustomHash func(data []byte) uint32

// HashRing 哈希环
type HashRing struct {
	hash     CustomHash
	replicas int            // 虚拟节点数量
	keys     []int          // 哈希环
	hashMap  map[int]string // 节点的哈希映射关系,key是哈希值取模,val是节点
}

func NewHashRing(fn CustomHash, replicas int) *HashRing {
	r := &HashRing{
		hash:     fn,
		replicas: replicas,
		hashMap:  make(map[int]string),
	}

	// 默认的哈希算法:CRC32
	if r.hash == nil {
		r.hash = crc32.ChecksumIEEE
	}
	return r
}
