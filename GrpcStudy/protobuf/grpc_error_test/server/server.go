package main

import (
	"GrpcStudy/protobuf/grpc_error_test/proto"
	context "context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net"
	"os"
	"time"
)

type Service struct {
	proto.UnimplementedGreeterServer
}

func (s *Service) SayHello(ctx context.Context, person *proto.Person) (*proto.Person, error) {
	fmt.Println(ctx)
	time.Sleep(1.1e9)
	fmt.Println(ctx.Err())
	go DB(ctx, 1)
	return &proto.Person{
		Id: 32,
	}, nil
}

func DB(ctx context.Context, id int) error {
	fmt.Println("我会入库么")
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)
	dsn := "root:root@tcp(127.0.0.1:3306)/db_sns_admin?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	create := db.WithContext(ctx).Create(&TbUser{
		UserName:    "wangShao",
		ServiceType: "AMS",
	})
	fmt.Println(create)
	return nil
}

// TbUser 用户表
type TbUser struct {
	UserName    string `gorm:"column:user_name" db:"user_name" json:"user_name" form:"user_name"`             //用户名
	ServiceType string `gorm:"column:service_type" db:"service_type" json:"service_type" form:"service_type"` //当前选中的业务
	UpdateTime  int64  `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"`     //更新时间
	CreateTime  int64  `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"`     //创建时间
}

func (u *TbUser) TableName() string {
	return "tb_user"
}

type Validator interface {
	Validate() error
}

func main() {
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if r, ok := req.(Validator); ok {
			err = r.Validate()
			if err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}

		ctx, cancel := context.WithDeadline(ctx, time.Now())
		defer cancel()

		return handler(ctx, req)
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(interceptor), grpc.ConnectionTimeout(1e9))
	server := grpc.NewServer(opts...)
	proto.RegisterGreeterServer(server, &Service{})
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	server.Serve(listen)
}
