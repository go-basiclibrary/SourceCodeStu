package main

import (
	"fmt"
	"strings"
)

func main() {
	// 假设你已经获取了多个ID值的列表
	idList := []int{1, 2, 3, 4, 5}

	// 构建IN查询条件
	placeholders := make([]string, len(idList))
	args := make([]interface{}, len(idList))
	for i, id := range idList {
		placeholders[i] = "?"
		args[i] = id
	}
	inCondition := fmt.Sprintf("IN (%s)", strings.Join(placeholders, ", "))

	// 构建查询语句
	query := fmt.Sprintf("SELECT * FROM 表名 WHERE id %s", inCondition)

	fmt.Println(query)
}
