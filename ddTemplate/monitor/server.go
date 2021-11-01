package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strings"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/monitor/database"
)

func main() {
	rpc.Register(new(Notify))
	listener, err := net.Listen("tcp", ":8899")
	if err != nil {
		fmt.Println("listen error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//新协程来处理--json
		go jsonrpc.ServeConn(conn)
	}
}

type Tmp struct {
	Id         int    `json:"id"`
	Webhook    string `json:"webhook"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	At         string `json:"at"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func getTmp(DB *sql.DB, data map[string]interface{}) {
	var tmp Tmp
	row := DB.QueryRow("select * from `parking_template` where id = " + data["templateId"].(string))
	if err := row.Scan(&tmp.Id, &tmp.Webhook, &tmp.Title, &tmp.Content, &tmp.At, &tmp.CreateTime, &tmp.UpdateTime); err != nil {
		fmt.Println(err)
	}
	title := tmp.Title
	v, _ := syscall.Getenv("ENV_NAME")
	if v == "dev" {
		title = title + "(测试)"
	}

	str := tmp.Content

	data["title"] = title
	for key, val := range data {
		str = strings.Replace(str, key, val.(string)+"\\n", -1)
	}

	times := time.Now().Format("2006-01-02 15:04:05")

	context := `{
			"msgtype": "text",
			"text": {"content":"通知类型：` + title + `\n通知时间：` + times + `\n` + str + `"},
			"at": {
				"atMobiles": [
		    		` + tmp.At + `
				],
					"isAtAll": false
			}}`

	//
	//	context := `{
	//		"msgtype": "markdown",
	//     	"markdown": {
	//         	"title":"`+title+`",
	//         	"text": "#### 杭州天气 @15711258458 \n ### 通知时间 `+times+` \n `+str+` "
	//     	},
	//      	"at": {
	//          	"atMobiles": [
	//              	"15711258458"
	//          	],
	//          	"isAtAll": false
	//      	}
	//	}`

	fmt.Println(context)
	//创建一个请求
	req, err := http.NewRequest("POST", tmp.Webhook, strings.NewReader(context))
	if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{}
	//设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-agent", "firefox")
	//发送请求
	resp, err := client.Do(req)
	//关闭请求
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		fmt.Println("handle error")
	}
}

type Notify string

type Reply struct{}

func (n *Notify) Send(data map[string]interface{}, reply *Reply) error {
	getTmp(database.DB, data)
	return nil
}

func (n *Notify) GetUrl(data map[string]interface{}, reply *Reply) error {
	getTmp(database.DB, data)
	return nil
}
