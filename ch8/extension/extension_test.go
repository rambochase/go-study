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

//复合基本上都要这样子重写，另外有方法参考extension2内嵌的写法就基本上不用这样写多一次就能拥有父结构体的方法
func (d *Dog) Speak() {
	//d.p.Speak()
	fmt.Println("dog wangwang")
}

func (d *Dog) SpeakTo(host string) {
	d.Speak()
	fmt.Println("dog speak to ",host)
	str := fmt.Sprintf("dog speak to %s",host)
	fmt.Println(str)
	//d.p.SpeakTo(host)	
}

func TestDog(t *testing.T) { 
	dog := new(Dog)
	dog.Speak()
	dog.SpeakTo("ric")
}


