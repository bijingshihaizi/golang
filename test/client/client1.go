package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

//注意字段必须是导出
type Args struct {
	A, B int
}

type Reply struct {
	C int
}

func main() {
	//连接远程rpc服务
	//这里使用jsonrpc.Dial
	rpc, err := jsonrpc.Dial("tcp", "172.28.134.7:8099")
	if err != nil {
		log.Fatal(err)
	}

	reply:=0

	//调用远程方法
	//注意第三个参数是指针类型
	err2 := rpc.Call("adds", Args{1, 2}, &reply)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(reply)

}
