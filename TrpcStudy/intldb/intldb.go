package main

import (
	"context"
	"git.tencent.com/intl/intl_comm/intldb"
	"git.tencent.com/trpc-go/trpc-database/mysql"
	"git.tencent.com/trpc-go/trpc-go/client"
)

type TbUser struct {
	Id            int    `db:"id" json:"id"`                           //ID
	User          string `db:"user" json:"user"`                       //用户RTX
	RoleTypeId    int    `db:"role_type_id" json:"role_type_id"`       //角色（1-管理员，2-成员，99-超级管理员）
	ServiceItemId string `db:"service_item_id" json:"service_item_id"` //服务项ID（0-超级管理员）
	AddUser       string `db:"add_user" json:"add_user"`               //添加人
	LoadTime      string `db:"load_time" json:"load_time"`             //更新时间0oll
}

type billadminInterfaceServiceImpl struct {
	cli *Client
}

func main() {

}

var Global *Client

func init() {
	Global = newCli()
}

func newCli() *Client {
	c := &Client{}
	c.dbCli = intldb.NewMysqlClient("",
		client.WithTarget("dsn://root:root@tcp(127.0.0.1:3306)/bill_admin?timeout=1s"))
	return c
}

type Client struct {
	dbCli *intldb.MysqlClient
}

// Transaction 事务执行
func (c *Client) Transaction(ctx context.Context, fn mysql.TxFunc, opts ...mysql.TxOption) error {
	return c.dbCli.Client.Transaction(ctx, fn, opts...)
}

func (c *Client) TExec(ctx context.Context, sql string) error {
	_, err := c.dbCli.TExec(ctx, sql)
	return err
}
