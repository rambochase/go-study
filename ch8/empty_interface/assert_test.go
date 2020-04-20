package empty_interface

import (
	"testing"
	"fmt"
)

//interface{}即为空接口
func doSomething(p interface{}) {
	if i,ok := p.(int);ok {
		fmt.Println("int ",i)
		return 
	}
	if s,ok := p.(string);ok {
		fmt.Println("string ",s)
		return 
	}
	fmt.Println("unknow type")
}

//doSomething 还可以通过switch
func doSomething2(p interface{}) {
	switch v := p.(type) {
	case int: fmt.Println("int ",v)
	case string: fmt.Println("string ",v)
	case bool: fmt.Println("bool ",v)
	default: fmt.Println("unknow type")
	}
}

func TestAssert(t *testing.T) {
	doSomething(10)
	doSomething("10")
	doSomething(3.15)

	doSomething2(10)
	doSomething2("10")
	doSomething2(true)
	doSomething2(3.15)
}