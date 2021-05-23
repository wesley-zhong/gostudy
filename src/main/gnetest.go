package main

import (
	"os"

	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"

	"github.com/panjf2000/gnet"
)

type echoServer struct {
	*gnet.EventServer
}

func (es *echoServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	incomeStr := string(frame[:])
	out = []byte(incomeStr + "aaa")
	log.Info("receive msg ", incomeStr)
	return
}
func (es *echoServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	log.Info("new connectd ", c.RemoteAddr().String())
	return
}

func main() {
	log.SetOutput(os.Stdout)
	a := "hello"
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jst, _ := jsoniter.MarshalToString(&a)
	log.Infof("jsonstr= %s", jst)

	echo := new(echoServer)
	err := gnet.Serve(echo, "tcp://:9000", gnet.WithMulticore(true))

	log.Fatal(err)
}
