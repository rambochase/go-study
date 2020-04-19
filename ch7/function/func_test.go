package function

import (
	"testing"
	"time"
	"fmt"
)

// 多返回值
func returnMultiVal() (int,int) {
	return 12,15
}


//计算时长的方法
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:",time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second*1)
	return op
}

func TestFunc(t *testing.T) {
	a,b := returnMultiVal();
	t.Log(a,b)
	c,_ := returnMultiVal();//忽略其中一个返回值
	t.Log(c)

	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

