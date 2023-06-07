//client.go
package main

import (
	"GrpcStudy/protobuf/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

type Client struct {
	Url string
}

func (c *Client) dial() *grpc.ClientConn {

	conn, err := grpc.Dial(c.Url, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return conn

}

func (c *Client) GetStream() {
	conn := c.dial()
	defer conn.Close()

	client := proto.NewStreamServiceClient(conn)
	stream, err := client.GetStream(context.Background(), &proto.StreamRequestData{
		Data: "time",
	})
	if err != nil {
		return
	}

	for {
		data, _ := stream.Recv()
		fmt.Println(data)

	}

}

func (c *Client) PostStream() {
	conn := c.dial()
	defer conn.Close()

	client := proto.NewStreamServiceClient(conn)
	stream, err := client.PostStream(context.Background())
	if err != nil {
		return
	}

	request := &proto.StreamRequestData{}
	for {
		request.Data = fmt.Sprintf("当前时间: %v", time.Now().Unix())
		stream.Send(request)
		time.Sleep(time.Second)
	}
}

func main() {
	client := Client{"127.0.0.1:9999"}
	client.GetStream()
	//client.PostStream()
}
