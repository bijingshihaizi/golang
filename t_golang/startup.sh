#!/bin/bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

go get google.golang.org/grpc@v1.37.0
go get golang.org/x/net/context@v0.0.0-20210614182718-04defd469f4e
go get github.com/golang/protobuf/protoc-gen-go/generator@v1.5.2
go get github.com/golang/protobuf/ptypes/struct@v1.5.2
go get "github.com/robfig/cron"
go mod init


# 安装全部包

# 安装热加载调试
go get -v -u github.com/pilu/fresh
#go run *.go
fresh