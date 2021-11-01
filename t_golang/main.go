package main

import (
	// "bytes"
	// "io/ioutil"
	// "net/http"

	// "github.com/rpcxio/rpcx-gateway"

	// "github.com/smallnest/rpcx/codec"
	// "time"
	"sync"	
	// "reflect"
	// "unsafe"
	// "sort"

	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
	// "context"
	"fmt"
	// "os"
   // "github.com/bwmarrin/snowflake"	// "github.com/robfig/cron"
	// "time"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	pb "github.com/t_golang/api"
)

func main()  {
	wg := &sync.WaitGroup{}

    conn, err := grpc.Dial("192.168.199.88:10029", grpc.WithInsecure())
    if err != nil {
        fmt.Println("err 31",err)
    }
    defer conn.Close()

    client := pb.NewTerminalMessageServiceClient(conn)
    reply, err := client.GetTerminalStatus(context.Background(), &pb.GetTerminalStatusRequest{TerminalId: uint32(1)})
    if err != nil {
		fmt.Println("41  line:",err)
    }
    fmt.Println(reply)

	wg.Wait()
}

// func task(wg *sync.WaitGroup)  {
// 	i := 0
// 	c := cron.New()
// 	spec := "*/1 * * * * ?"
// 	c.AddFunc(spec, func() {
// 		i++
// 		fmt.Println(i)
// 		wg.Done()
// 	})
// 	c.Start()
// }

// func random(n int) <-chan int {
//     c := make(chan int)
//     go func() {
//         defer close(c)
//         for i := 0; i < n; i++ {
//             select {
//             case c <- 0:
//             case c <- 1:
//             }
//         }
//     }()
//     return c
// }


// type QueryRuleList struct{
// 	Id int `json:"id"`
// 	CurPage int `json:"cur_page"`
// 	PageSize int `json:"page_size"`
// 	PlateNo string `json:"plate_no"`
// 	PlateName string `json:"plate_name"`
// 	RuleName string `json:"rule_name"`
// 	MainUserId int `json:"main_user_id"`
// }
// type RuleListReply struct{
// 	Count int64 `json:"count"`
// 	RuleList []RuleList
// }
// type RuleList struct{
// 	Id int `json:"id"`
// 	MainUserId int `json:"main_user_id"`
// 	RuleList string `json:"rule_list"`
// 	RuleListType int `json:"rule_list_type"`
// 	RuleName string `json:"rule_name"`
// 	RuleType int `json:"rule_type"`
// 	ListContent []Content `json:"list_content" gorm:"-"`
// }
// type Content struct{
// 	Id int `json:"id"`
// 	PlateNo string `json:"plate_no"`
// 	PlateMobile string `json:"plate_mobile"`
// 	PlateRealname string `json:"plate_realname"`
// 	PlateType string `json:"plate_type"`
// 	RuleId int `json:"rule_id"`
// }

// type Test struct{
// 	Id int `json:"id"`
// 	TopCpu int `json:"top_cpu"`
// 	IsSuccess int `json:"is_success"`
// 	Duration int64 `json:"duration"`
// 	StartTime int64 `json:"start_time"`
// 	EndTime int64 `json:"end_time"`
// }



func mistakes()  {
	
}

func routinetest()  {
	// dsn := "root:pprt123@tcp(39.106.10.158:3306)/parkingpay-vnloco-test?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	fmt.Println("db err",err)
	// }

	// wg := &sync.WaitGroup{}
	// start_time := time.Now().Unix()
	// fmt.Println("start time",time.Now().Unix())
	// var is_success = 1
	// var cxt context.Context
	// for j:=0; j<100; j++ {
	// 	for i:=0; i<1000; i++{
	// 		go func ()  {
	// 			// request(wg)
	// 			wg.Add(1)

	// 			cc := &codec.MsgpackCodec{}

	// 			// request 
	// 			args := &QueryRuleList{}
	// 			data, _ := cc.Encode(args)

	// 			req, err := http.NewRequest("POST", "http://192.168.199.88:8333/", bytes.NewReader(data))
	// 			if err != nil {
	// 				is_success = 2
	// 				fmt.Println("failed to create request: ", err)
	// 			}
				
	// 			// set extra headers
	// 			h := req.Header
	// 			h.Set(gateway.XSerializeType, "3")
	// 			h.Set(gateway.XServicePath, "Rsv")
	// 			h.Set(gateway.XServiceMethod, "GetRuleList")


	// 			// send to gateway
	// 			res, err := http.DefaultClient.Do(req)
	// 			if err != nil {
	// 				is_success = 2
	// 				fmt.Println("failed to call: ", err)
	// 			}
	// 			defer res.Body.Close()

	// 			// handle http response
	// 			replyData, err := ioutil.ReadAll(res.Body)
	// 			if err != nil {
	// 				is_success = 2
	// 				fmt.Println("failed to read response: ", err)
	// 			}

	// 			// parse reply
	// 			reply := &RuleListReply{}
	// 			err = cc.Decode(replyData, reply)
	// 			if err != nil {
	// 				is_success = 2
	// 				fmt.Println("failed to decode reply: ", err)
	// 			}
	// 			wg.Done()
	// 			db.Table("test").WithContext(cxt).
	// 			Create(&Test{IsSuccess: is_success})
	// 		}()
	// 	}
	// 	wg.Wait()
	// 	end_time := time.Now().Unix()
	// 	duration := end_time-start_time
	// 	db.Table("test").WithContext(cxt).
	// 	Create(&Test{Duration:duration, StartTime:start_time,EndTime:end_time})
	// 	fmt.Println("end time",time.Now().Unix())
	// }
}


func test()  {
	// done := make(chan int, 10) // 带 10 个缓存

    // // 开N个后台打印线程
    // for i := 0; i < cap(done); i++ {
    //     go func(){
    //         fmt.Println("你好, 世界")
    //         done <- 1
    //     }()
    // }

    // // 等待N个后台线程完成
    // for i := 0; i < cap(done); i++ {
    //     <-done
    // }

	    // Ctrl+C 退出
		// sig := make(chan os.Signal, 1)
		// signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		// fmt.Printf("quit (%v)\n", <-sig)
}