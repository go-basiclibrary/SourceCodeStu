//server.go
package main

import (
	"fmt"
	"gobasic/protobuf/proto"
	"google.golang.org/grpc"
	"net"
	"time"
)

type Server struct{
	proto.UnimplementedStreamServiceServer
}

//GetStream 服务端流模式:服务端不断的给客户端发数据
func (server  *Server ) GetStream(request *proto.StreamRequestData,s proto.StreamService_GetStreamServer)  (err error) {
	response := &proto.StreamResponsetData{}
	if request.Data == "time" {

		for {
			response.Msg = fmt.Sprintf("当前时间为: %v",time.Now().Unix())
			err = s.Send(response)
			time.Sleep(time.Second)
		}

	}
	return
}


// PostStream 客户端流模式: 客户端不断的给服务端发数据，服务端接收
func (server *Server) PostStream(s proto.StreamService_PostStreamServer) error {
	for {
		//可以发现服务端和客户端的流辅助接口均定义了Send和Recv方法用于流数据的双向通信。
		//服务端在循环中接收客户端发来的数据，如果遇到io.EOF表示客户端流被关闭，
		//如果函数退出表示服务端流关闭。生成返回的数据通过流发送给客户端，双向流数据的发送和接收都是完全独立的行为。
		//需要注意的是，发送和接收的操作并不需要一一对应，用户可以根据真实场景进行组织代码。
		requestData,err := s.Recv()
		if err != nil {
			panic(err)
		}
		fmt.Println(requestData)
	}
	return nil
}


func (server *Server) AllStream(s proto.StreamService_AllStreamServer) error {
	return nil
}

func main() {
	server := grpc.NewServer()
	proto.RegisterStreamServiceServer(server, new(Server))
	lis, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		panic("fail to listen port:" + err.Error())
	}
	fmt.Println("启动grpc服务成功:0.0.0.0:9999")
	err = server.Serve(lis)
	if err != nil {
		panic("fail to start grpc:" + err.Error())
	}
}

