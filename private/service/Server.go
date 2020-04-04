
package service

import (
	"context"
	"errors"
	"fmt"
	"gRPC-Server/private/pb"
)

type MyServer struct {}	// 名称自定义，但需要实现proto中定义的全部方法

func (s *MyServer) Echo(ctx context.Context, in *pb.ReqMsg) (*pb.RspMsg, error) {
	var rsp = new(pb.RspMsg)
	var err error

	if in.Name == "" {
		err = errors.New("the field of name is nil")
	} else {
		rsp.Age_Name = fmt.Sprintf("name:%s, age:%d", in.Name, in.Age)
		fmt.Printf("rsp=%+v\n", rsp)
	}

	return rsp, err
}
