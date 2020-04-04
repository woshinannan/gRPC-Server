
package main

import (
	"fmt"
	"gRPC-Server/private/pb"
	"gRPC-Server/private/service"
	"google.golang.org/grpc"
	"net"
)

// 定义服务端监听地址和端口：一般地址就是"127.0.0.1"，代表本机；有的机器有多个ip地址，可以填其他ip。端口自定义，不与系统端口冲突就行
const Address = ":50052"

func main() {
	// 创建监听对象
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Printf("net.Listen failed, err=%+v\n", err)
	} else {
		fmt.Printf("net.Listen success\n")
	}

	// 实现gRPC Server
	grpc_server := grpc.NewServer()

	// 注册service.MyServer类到gRPC，为客户端提供服务
	pb.RegisterMyServiceNameServer(grpc_server, new(service.MyServer))

	// 监听
	fmt.Printf("start Serve(listen)\n")
	err = grpc_server.Serve(listen)
	if err != nil {
		fmt.Printf("Serve(listen) failed, err=%+v\n", err)
	}
}
