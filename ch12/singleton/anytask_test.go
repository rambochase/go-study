package singleton

import (
	"testing"
	"fmt"
	"time"
	"runtime"
)

func runTask(task_id int) string {
	time.Sleep(time.Millisecond*100)
	return fmt.Sprintf("task from %d\n",task_id)
}

func GetFirstRespone() string {
	n := 20
	retCh := make(chan string,num)//防止内存泄漏则使用buffer channel
	for i := 0; i < n; i++ {
		go func(j int) {
			ret := runTask(j)
			retCh <- ret
		}(i)
	}
	return <-retCh//通过通道阻塞来实现，一旦通道有数据就返回了
}

func TestAnyTesk(t *testing.T) {
	t.Log("Before:",runtime.NumGoroutine())//获取协程数
	t.Log(GetFirstRespone())
	time.Sleep(time.Second*1)
	t.Log("After:",runtime.NumGoroutine())
}