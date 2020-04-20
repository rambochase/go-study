package duck_type

import(
	"testing"
)

//接口定义(不同的结构体通过实现这样一个接口方法来实现多态)
type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {

}

//接口实现，即绑定在结构指针上, 没有显式implements
func(g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

//客户端
func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	s := p.WriteHelloWorld()
	t.Log(s)
}
