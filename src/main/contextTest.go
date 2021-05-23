package main

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
)

var key string = "name"

func main1() {
	ctx, cancel := context.WithCancel(context.Background())
	//附加值
	valueCtx := context.WithValue(ctx, key, "【监控1】")
	go watch(valueCtx)

	ctx1, cancel2 := context.WithCancel(context.Background())
	valueCtx1 := context.WithValue(ctx1, key, "【监控2】")
	go watch(valueCtx1)
	time.Sleep(10 * time.Second)
	log.Println("可以了，通知监控停止")
	cancel()
	cancel2()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(10 * time.Second)

}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//取出值
			log.Println(ctx.Value(key), "监控退出，停止了...")
			return
		default:
			//取出值
			log.Println(ctx.Value(key), "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}

}
