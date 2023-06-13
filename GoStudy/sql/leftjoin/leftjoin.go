package main

import "fmt"

func main() {
	sql := fmt.Sprintf("SELECT s.`service_item_id`,s.`status`,i.`service_item` "+
		"FROM %s AS s "+
		"LEFT JOIN %s AS i ON s.`service_item_id` = i.`service_item_id WHERE `date`=%q`",
		"bill_admin.tb_billing_summary", "bill_admin.tb_service_item", "2023-06")

	fmt.Println(sql)
}
