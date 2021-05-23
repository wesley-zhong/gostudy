package mymodule

import (
	log "github.com/sirupsen/logrus"
)

type Pool struct {
	Queue         chan func() int
	RuntineNumber int
	Total         int

	Result         chan int
	FinishCallback func()
}

//初始化
func (self *Pool) Init(runtineNumber int, total int) {
	self.RuntineNumber = runtineNumber
	self.Total = total
	self.Queue = make(chan func() int, total)
	self.Result = make(chan int, total)
}

func (self *Pool) Start() {
	//开启 number 个goruntine
	for i := 0; i < self.RuntineNumber; i++ {
		go func() {
			for {
				task, ok := <-self.Queue
				if !ok {
					break
				}
				err := task()
				self.Result <- err
			}
		}()
		log.Printf("-------------process result index %d", i)
	}

	//获取每个任务的处理结果
	for j := 0; j < self.RuntineNumber; j++ {
		log.Printf("===================== process result index %d waited ", j)
		res, ok := <-self.Result
		if !ok {
			log.Error(" channel closed")
			break
		}
		log.Printf("===================== process result index %d waited result  res=%d ", j, res)
	}

	//结束回调函数
	if self.FinishCallback != nil {
		self.FinishCallback()
	}
}

//关闭
func (self *Pool) Stop() {
	close(self.Queue)
	close(self.Result)
}

func (self *Pool) AddTask(task func() int) {
	self.Queue <- task
}

func (self *Pool) SetFinishCallback(fun func()) {
	self.FinishCallback = fun
}
