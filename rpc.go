package main

import (
	"net/rpc"
	"net"
	"fmt"
	"net/rpc/jsonrpc"
	"context"
)

func main()  {
	rpc.Register(new(Msgs))
	listener, err := net.Listen("tcp", ":8899")
	if err != nil {
		fmt.Println("listen error:", err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//新协程来处理--json
		go jsonrpc.ServeConn(conn)
	}
}

type Msgs string

//pda车位状态通知
func (s *Msgs) SpaceMsg (msg string,reply *int) (error){
	// 连接服务端接口
	conn, err := grpc.Dial("192.168.31.214:42809", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	client := impush.NewParkingPayPDAServiceClient(conn)

	reply_space, err_space := client.UpSpaceStatusChangeData(context.Background(), &impush.ParkingSpaceStatusReportRequest{Msg:msg})
	if err_space != nil {
		fmt.Println(err_space)
	}
	fmt.Println(reply_space)
	return nil
}