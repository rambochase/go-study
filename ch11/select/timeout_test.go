package channel

import (
	"testing"
	"time"
	"fmt"
)

func service()  string {
	time.Sleep(time.Millisecond*300)//模拟耗时操作 300ms
	return "Done"
}

//异步的
func asyncService() chan string {
	retCh := make(chan string,1) //声明一个buffer通道
	go func () {
		ret := service()
		fmt.Println("get service result")
		retCh <- ret //把结果放到通道里面去
		fmt.Println("service exited")
	}()
	return retCh
}

//超时控制
func TestSelect(t *testing.T) {
	select {
	case ret := <-asyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond*100): //设置100ms超时
		t.Error("timeout")
	}
}