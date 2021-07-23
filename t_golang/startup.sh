#!/bin/bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

# 安装rpcx 
go get -u -v github.com/smallnest/rpcx/...
go get github.com/rpcxio/rpcx-gateway
go get github.com/smallnest/rpcx/codec
go get gorm.io/driver/mysql
go get gorm.io/gorm

go mod init


# 安装全部包

# 安装热加载调试
go get -v -u github.com/pilu/fresh
#go run *.go
fresh