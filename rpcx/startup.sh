#!/bin/bash

# go get -u -v -tags "reuseport quic kcp zookeeper etcd consul ping" github.com/smallnest/rpcx/...
go get -u github.com/nacos-group/nacos-sdk-go
go get -u -v github.com/smallnest/rpcx/...
go mod init
go run main.go 

