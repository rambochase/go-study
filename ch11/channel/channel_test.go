package channel

import (
	"testing"
	"time"
	"fmt"
)

func service()  string {
	time.Sleep(time.Millisecond*100)//模拟耗时操作
	return "Done"
}

//异步的
func asyncService() chan string {
	retCh := make(chan string,1) //声明一个通道
	go func () {
		ret := service()
		fmt.Println("get service result")
		retCh <- ret //把结果放到通道里面去
		fmt.Println("service exited")
	}()
	return retCh
}

func otherService() {
	fmt.Println("do something")
	time.Sleep(time.Millisecond*200)//模拟耗时操作
	fmt.Println("do something ok")
}

//串行
func TestSerial(t *testing.T) {
	ret := service()
	otherService()
	t.Log(ret)
}

//使用了channel实现异步
func TestAsync(t *testing.T) {
	retCh := asyncService() // 获取这个通道
	otherService()
	t.Log(<-retCh)
}

