package function

import(
	"testing"
	"fmt"
)

func DeLay() {
	fmt.Println("defer function")
}
func TestDefer(t *testing.T) {
	defer DeLay()//函数返回前执行
	fmt.Println("I do sometime")
	fmt.Println("do other thing")
	panic("Fatal error")//报了错误仍然会执行defer函数
}