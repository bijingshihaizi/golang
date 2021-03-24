package utils

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"syscall"
)

var cfg *goconfig.ConfigFile

func GetConfigIni() (err error) {
	v, _ := syscall.Getenv("ENV_NAME")
	fmt.Println(v)
	config, err := goconfig.LoadConfigFile("./config/config." + v + ".ini")
	if err != nil {
		fmt.Println("配置文件读取错误,找不到配置文件", err)
		return err
	}
	cfg = config
	return nil
}

func  GetDatabase() (mysql string, err error) {
	GetConfigIni()
	if mysql, err = cfg.GetValue("database", "mysql"); err != nil {
		fmt.Println("配置文件中不存在mysql", err)
		return mysql, err
	}
	return mysql, nil
}

func GetSelf() (port string, err error) {
	GetConfigIni()
	if port, err = cfg.GetValue("self", "port"); err != nil {
		fmt.Println("配置文件中不存在port", err)
		return port, err
	}
	return port, nil
}

func GetGrpc() (addr string, err error)  {
	GetConfigIni()
	if addr, err = cfg.GetValue("grpc", "addr"); err != nil {
		fmt.Println("配置文件中不存在 addr", err)
		return addr, err
	}
	return addr, nil
}

func GetLink() (link string, err error)  {
	GetConfigIni()
	if link, err = cfg.GetValue("link", "link"); err != nil {
		fmt.Println("配置文件中不存在 link", err)
		return link, err
	}
	return link, nil
}

func GetFile() (prefix string, err error)  {
	GetConfigIni()
	if prefix, err = cfg.GetValue("file", "prefix"); err != nil {
		fmt.Println("配置文件中不存在 prefix", err)
		return prefix, err
	}
	return prefix, nil
}