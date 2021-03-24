package main

import(
    "flag"
    "github.com/user/services"
    "github.com/smallnest/rpcx/server"
)

var (
    addr = flag.String("addr", ":8972", "server address")
)

func main() {
    flag.Parse()
    s := server.NewServer()
    s.RegisterName("User", new(services.User), "")
    s.Serve("tcp", *addr)
}