package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	cpuStatFile = "/proc/stat"
)

// 采样结果对象
type result struct {
	used uint64 // CPU 使用时间
	idle uint64 // CPU 闲置时间
}

// CPU 指标采样函数
func sample() (*result, error) {
	data, err := os.ReadFile(cpuStatFile)
	if err != nil {
		return nil, err
	}

	res := &result{}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		// 为了简化演示
		// 这里只取所有 CPU 总的统计数据
		if len(fields) == 0 || fields[0] != "cpu" {
			continue
		}

		// 将第一行数据分割为数组
		n := len(fields)
		for i := 1; i < n; i++ {
			if i > 8 {
				continue
			}

			// 解析每一列的数值
			val, err := strconv.ParseUint(fields[i], 10, 64)
			if err != nil {
				return nil, err
			}

			// 第 4 列表示 CPU 空闲时间
			// 第 5 列表示 等待 I/O 的 CPU 时间
			if i == 4 || i == 5 {
				res.idle += val
			} else {
				res.used += val
			}
		}

		return res, nil
	}

	return res, nil
}

func main() {
	// 获取第一次采样结果
	first, err := sample()
	if err != nil {
		log.Fatal(err)
	}

	// 模拟一些 CPU 密集型任务
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10000; i++ {
		_ = math.Sqrt(rand.Float64())
	}

	// 获取第二次采样结果
	second, err := sample()
	if err != nil {
		log.Fatal(err)
	}

	// 计算两次采样期间 CPU 的空闲时间
	idle := float64(second.idle - first.idle)
	// 计算两次采样期间 CPU 的使用时间
	used := float64(second.used - first.used)
	// CPU 利用率 = CPU 使用时间 / (CPU 闲置时间 + CPU 使用时间)
	var usage float64
	if idle+used > 0 {
		usage = used / (idle + used) * 100
	}

	fmt.Printf("CPU usage is %f%%\n", usage)
}
