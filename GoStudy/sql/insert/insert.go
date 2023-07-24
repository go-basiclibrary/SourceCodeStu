package main

import (
	"context"
	"fmt"
	"git.tencent.com/intl/intl_comm/intldb"
	"git.tencent.com/trpc-go/trpc-go/client"
	"git.tencent.com/trpc-go/trpc-go/log"
	"github.com/spf13/cast"
	"math/rand"
	"strings"
	"time"
)

func main() {
	c, err := NewLocalDB()
	if err != nil {
		panic(err)
	}
	now := time.Now()
	var billings []*TbServiceSonitemBilling
	for j := 0; j < 1800; j++ {
		billings = append(billings, &TbServiceSonitemBilling{
			ServiceItemID:    "121",
			ServiceSonitemID: "73",
			CrosName:         "英雄联盟手游",
			KpiCode:          randIntRunes(5),
			OtherSystemsID:   "202209231717",
			ProjectName:      fmt.Sprintf("优惠礼包 %d", j),
			Price:            "1500",
			Unit:             "月",
			Quantity:         cast.ToString(rand.Intn(10) + 1), // 生成1到10之间的随机数量
			DiscountFactor:   "0.75",                           // 生成0到1之间的随机折扣系数
			Remark:           "无呵呵大大饿饿饿",
		})
		billings[j].SsbID = GenerateSsbID(billings[j].KpiCode)
		billings[j].RecordKey = GetCurrencyRecordKey(billings[j].ServiceItemID, billings[j].ServiceSonitemID, billings[j].KpiCode, billings[j].ProjectName)
	}

	err = c.BatchInsertBillings(billings)
	if err != nil {
		log.Errorf("batch insert billings err: %s", err)
	}
	since := time.Since(now)
	fmt.Println(since)
}

func GetCurrencyRecordKey(srvID, itemSonID string,
	productCode, projectName string) string {
	return fmt.Sprintf("%s:%s:%s:%s:%s:%s",
		srvID, itemSonID, GetCurrentMonthCors(), productCode, "", projectName)
}

// GetCurrentMonthCors 获取当月账期格式 202306
func GetCurrentMonthCors() string {
	tm := time.Unix(time.Now().Unix(), 0)
	return tm.Format("200601")
}

// Client store
type Client struct {
	dbCli *intldb.MysqlClient
}

var GlobalDB *Client

// NewLocalDB 本地DB初始化
func NewLocalDB() (c *Client, err error) {
	GlobalDB = &Client{}
	GlobalDB.dbCli = intldb.NewMysqlClient("", client.WithTarget("dsn://root:123456@tcp(10.250.221.54:3306)/bill_admin?timeout=1s"))

	return GlobalDB, nil
}

func (c *Client) BatchInsertBillings(billings []*TbServiceSonitemBilling) error {
	// 存放 (?,?的slice)
	valueStrings := make([]string, 0, len(billings))
	// 存放values的slice
	valuesArgs := make([]interface{}, 0, len(billings)*19)
	// 遍历users准备相关数据
	for _, bill := range billings {
		// 插入值对应的个数
		valueStrings = append(valueStrings, "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
		valuesArgs = append(valuesArgs, bill.SsbID, bill.AccountPeriod, bill.ServiceItemID, bill.ServiceSonitemID,
			bill.Status, bill.AddUser, bill.KpiCode, bill.CrosName, bill.ProjectName, bill.Price, bill.Unit,
			bill.Quantity, bill.TotalAmount, bill.DiscountFactor, bill.AmountAfterDiscount, bill.DeductionAmount,
			bill.Remark, bill.OtherSystemsID, bill.RecordKey)
	}

	sql := fmt.Sprintf("INSERT INTO %s "+
		"(`ssb_id`,`account_period`,`service_item_id`,`service_sonitem_id`,"+
		"`status`,`add_user`,`kpi_code`,`cros_name`,`project_name`,`price`,`unit`,"+
		"`quantity`,`total_amount`,`discount_factor`,`amount_after_discount`,deduction_amount,"+
		"`remark`,`other_systems_id`,`record_key`) VALUES %s", `tb_service_sonitem_billing`, strings.Join(valueStrings, ","))

	_, err := c.dbCli.TExec(context.TODO(), sql, valuesArgs...)
	return err
}

// TbServiceSonitemBilling 计费详情表
type TbServiceSonitemBilling struct {
	ID                  int32  `json:"id" db:"id"`                                       // ID
	SsbID               string `json:"ssb_id" db:"ssb_id"`                               // 计费详情ID的规则为kpicode+年份+月份+xxxx，例如：30077023040001
	AccountPeriod       string `json:"account_period" db:"account_period"`               // 计费日期
	ServiceItemID       string `json:"service_item_id" db:"service_item_id"`             // 服务项ID
	ServiceSonitemID    string `json:"service_sonitem_id" db:"service_sonitem_id"`       // 服务子项ID
	Status              int    `json:"status" db:"status"`                               // 填报进度1填报中；2已提交；3打回；4通过；5：未结算；0删除；
	AddUser             string `json:"add_user" db:"add_user"`                           // 填报人
	UpdateTime          string `json:"update_time" db:"update_time"`                     // 更新时间
	AddDate             string `json:"add_date" db:"add_date"`                           // 添加时间
	KpiCode             string `json:"kpi_code" db:"kpi_code"`                           // kpi代码
	CrosName            string `json:"cros_name" db:"cros_name"`                         // 业务名称
	ProjectName         string `json:"project_name" db:"project_name"`                   // 项目名称
	Price               string `json:"price" db:"price"`                                 // 单价 保留两位小数
	Unit                string `json:"unit" db:"unit"`                                   // 单价单位
	Quantity            string `json:"quantity" db:"quantity"`                           // 数量 保留两位小数
	TotalAmount         string `json:"total_amount" db:"total_amount"`                   // 合计金额
	DiscountFactor      string `json:"discount_factor" db:"discount_factor"`             // 折扣系数 保留两位小数
	AmountAfterDiscount string `json:"amount_after_discount" db:"amount_after_discount"` // 折扣后金额 保留四位小数
	Remark              string `json:"remark" db:"remark"`                               // 备注
	EditUser            string `json:"edit_user" db:"edit_user"`                         // 编辑人
	CallBackReason      string `json:"call_back_reason" db:"call_back_reason"`           // 打回原因
	AdjustmentQuantity  string `json:"adjustment_quantity" db:"adjustment_quantity"`     // 建议调整数量
	ReviewUser          string `json:"review_user" db:"review_user"`                     // 审核人
	ReviewDate          string `json:"review_date" db:"review_date"`                     // 审核时间
	SubmitUser          string `json:"submit_user" db:"submit_user"`                     // 提交人
	SubmitDate          string `json:"submit_date" db:"submit_date"`                     // 提交时间
	OtherSystemsID      string `json:"other_systems_id" db:"other_systems_id"`           // 其他系统ID多个;分号分隔
	DeductionAmount     string `json:"deduction_amount" db:"deduction_amount"`           // 减免金额
	RecordKey           string `json:"record_key" db:"record_key"`                       // 填报记录唯一key, srv_id:item_id:period:product_code:hr_code:project_name
	DeletedAt           int    `json:"deleted_at" db:"deleted_at"`                       // 删除时间戳
}

var numberRunes = []rune("123456789")

// randIntRunes 随机生成数字  并发安全
func randIntRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = numberRunes[rand.Intn(len(numberRunes))]
	}
	return string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateSsbID 生成计费详情ID  规则  kpicode-7位随机数,在库中保持唯一
func GenerateSsbID(kpiCode string) string {
	return fmt.Sprintf("%s-%s", kpiCode, randIntRunes(7))
}
