package extension2

import (
	"testing"
	"fmt"
)

type Pet struct {

}

type Dog struct {
	Pet//内嵌
}


/*--------------这是父结构体的方法 start ------------------*/
func(p *Pet) Speak() {
	fmt.Println("Pet speak···")
}

func(p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println("Pet speak to ",host)
}
/*--------------这是父结构体的方法 end ------------------*/

/*-------------重载父结构体的方法 start -----------------*/
func(d *Dog) Speak() {
	fmt.Println("I am dog,say wangwang")
}

func TestDog(t *testing.T) {
	var d *Dog = new(Dog)
	d.Speak()
	d.SpeakTo("ric")
}
