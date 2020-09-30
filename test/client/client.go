package main

import (
	"github.com/smallnest/rpcx/client"
	"log"
	//"service"
	"context"
	"flag"
	"fmt"
)

var (
	addr = flag.String("addr", "172.28.134.7:8099", "server address")
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

func main() {
	Peer2Peer()
}
func Peer2Peer() {
	flag.Parse()
	d := client.NewPeer2PeerDiscovery("tcp@" + *addr, "")
	xclient := client.NewXClient("CalculatorService", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	fmt.Println(xclient)
	defer xclient.Close()

	args := &Args{
		A: 10,
		B: 20,
	}

	reply := &Reply{}
	fmt.Println("===================================")
	err := xclient.Call(context.Background(), "Calculator.add", args, &reply)
	fmt.Println(err)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}