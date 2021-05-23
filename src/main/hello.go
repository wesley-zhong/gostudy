package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	log "github.com/sirupsen/logrus"
	"netease.com/mymodule"
)

func main1() {
	var p mymodule.Pool
	log.SetOutput(os.Stdout)
	// //启动trace goroutine
	// err = trace.Start(f)
	// if err != nil {
	// 	panic(err)
	// }
	//	defer trace.Stop()
	i := new(int)
	*i = 100
	log.Info("i = ", *i)

	fmt.Println("hell----")
	message := mymodule.Hello("gladays")
	fmt.Println(message)

	log.Println("init started")

	url := []string{"11111", "22222", "33333", "444444", "55555", "66666", "77777", "88888", "999999"}
	p.Init(9, len(url))

	for i := range url {
		u := url[i]
		var index = i
		p.AddTask(func() int {
			return Download(u, index)
		})
	}

	p.SetFinishCallback(DownloadFinish)
	p.Start()
	p.Stop()

	log.Info("start tcp server--------")
	//1 创建一个server句柄
	s := znet.NewServer()

	//2 配置路由
	s.AddRouter(0, &PingRouter{})

	//3 开启服务
	s.Serve()

}

func Download(url string, index int) int {
	time.Sleep(1 * time.Second)
	log.Printf("Download= %s  index =%d ", url, index)
	return index
}

func DownloadFinish() {
	log.Println("Download finsh")
}

//ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

//Ping Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	//先读取客户端的数据
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	//再回写ping...ping...ping
	err := request.GetConnection().SendBuffMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}
