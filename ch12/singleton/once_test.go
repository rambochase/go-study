package singleton

import (
	"testing"
	"fmt"
	"sync"
	"unsafe"
)

type SingleTon struct {

}

var once sync.Once
var Obj *SingleTon

//单例模式（懒汉式，线程安全）
func GetSingleTonInstance() *SingleTon {
	once.Do(func() {//这段方法只会运行一次
		fmt.Println("create Object")
		Obj = new(SingleTon)
	})
	return Obj
}

func TestOnce(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			sgle := GetSingleTonInstance()
			fmt.Printf("%x\n",unsafe.Pointer(sgle))
			wg.Done()
		}()
	}
	wg.Wait()
}