package channel

import (
	"testing"
	"sync"
	"fmt"
)

//通道广播与关闭

//生产者
func dataProducer(n int,retCh chan int,wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < n; i++ {
			retCh <- i
		}
		close(retCh)//做完后关闭channel
		wg.Done();
	}()
	
}

func dataReceiver(retCh chan int,wg *sync.WaitGroup) {
	go func () {
		for {
			if ret,ok := <-retCh;ok {
				fmt.Println(ret)
			} else {//检测到channel关闭，跳出循环
				fmt.Println(ok)
				break;
			}
		}
		wg.Done()
	}()
}

func TestBroadcast(t *testing.T) {
	var wg sync.WaitGroup
	n := 11
	retCh := make(chan int,1)//buffer channel
	wg.Add(1)
	dataProducer(n,retCh,&wg)//生产者
	
	wg.Add(1)
	dataReceiver(retCh,&wg)//消费者
	wg.Add(1)
	dataReceiver(retCh,&wg)//消费者
	wg.Add(1)
	dataReceiver(retCh,&wg)//消费者

	wg.Wait()
	t.Log("Done")
}

