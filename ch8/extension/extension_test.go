package extension

import (
	"testing"
	"fmt"
)

//相当PHP的父类(没有继承，其实也不能说是父类，不过可以这样理解)
type Pet struct {

}

//对应结构体的方法实现
func (p *Pet) Speak() {
	fmt.Println("...")
}

//对应结构体的方法实现
func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ",host)
}

//相当PHP的子类
type Dog struct {
	p *Pet
}

func (d *Dog) Speak() {
	d.p.Speak()
}

func (d *Dog) SpeakTo(host string) {
	d.p.SpeakTo(host)	
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.Speak()
	dog.SpeakTo("ric")
}


