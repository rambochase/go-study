package duck_type2

import (
	"testing"
	"fmt"
)

/**
	通过接口实现多态
*/

type code string

//接口定义
type Programmer interface {
	WriteHelloWorld() code
}

//go编程的结构体
type GoProgrammer struct {

}

//php编程结构体
type PhpProgrammer struct {

}

//各自实现对应的接口方法
//go
func(g *GoProgrammer) WriteHelloWorld() code {
	return "fmt.Println(\"Hello World!\")"
}

//php
func(p *PhpProgrammer) WriteHelloWorld() code {
	return "echo \"Hello World!\";"
}

//创建实例
func Instance(p Programmer) {
	str := p.WriteHelloWorld() 
	fmt.Printf("type is %T,%v\n",p,str)
}

//多态
func TestDuoTai(t *testing.T) {
	php_prog := new(PhpProgrammer)
	go_prog  := new(GoProgrammer)
	Instance(php_prog)
	Instance(go_prog)
}