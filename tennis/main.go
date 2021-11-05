package main

import (
	"encoding/base64"
	"os"
)

func main() {
	GetTarPic()
}

func GetInfo() {

}

func GetBackPic() {
	str64 := ""

	dist, _ := base64.StdEncoding.DecodeString(str64)
	f, _ := os.OpenFile("back.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.Write(dist)

}

func GetTarPic() {
	str64 := ""
	dist, _ := base64.StdEncoding.DecodeString(str64)
	f, _ := os.OpenFile("tar.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.Write(dist)
}
