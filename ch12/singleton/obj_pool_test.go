package singleton

import (
	"testing"
	"time"
	"errors"
	"unsafe"
	"fmt"
)

//对象
type Obj struct {

}

//对象池
type ObjPool struct {
	bufChan chan *Obj//拿来缓冲可重复使用的对象
}

//创建对象池
func CreateObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *Obj,numOfObj)//numOfObj就是对象池的长度
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &Obj{}
	}
	return &objPool
}

//获取对象
func(pool *ObjPool) GetObj(timeout time.Duration) (*Obj,error) {
	select {
	case ret := <-pool.bufChan:
		return ret,nil
	case <-time.After(timeout):
		return nil,errors.New("get Obj timeout")//超时控制
	}
}

//把对象放回对象池去
func(pool *ObjPool) ReleaseObj(obj *Obj) error {
	select {
	case pool.bufChan <- obj:
		return nil
	default: return errors.New("overflow")
	}
}

func TestObjPool(t *testing.T) {
	NumOfObj := 20
	pool := CreateObjPool(NumOfObj)
	for i := 0; i < 21; i++ {
		if objPointer,err := pool.GetObj(time.Second*1);err == nil{
			fmt.Printf("%x\n",unsafe.Pointer(objPointer))
		} else {
			t.Error(err)
		}
	}
}