package struct_test

import (
	"testing"
	 "fmt"
)

//定义结构体
type Employee struct {
	Id string
	Name string
	Age int
}

//行为方法的定义
//传递指针就可以减少内存拷贝
func(e *Employee) String() string {
	return fmt.Sprintf("ID:%s,Name:%s,Age:%d",e.Id,e.Name,e.Age)
}
//行为方法的定义
func (e Employee) Span() string {
	return fmt.Sprintf("ID:%s,Name:%s,Age:%d",e.Id,e.Name,e.Age)
}

func TestStruct(t *testing.T) {
	e := Employee{"0","ric",28}
	e1 := Employee{Name:"Rambo",Age:27}


	str := e.String()//调用行为方法
	str2 := e.Span()//调用行为方法
	str_e1 := e1.Span()
	str_e12 := e1.String()

	t.Log(str)
	t.Log(str2)
	t.Log(str_e1)
	t.Log(str_e12)
}