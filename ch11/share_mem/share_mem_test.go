package share_mem

import (
	"testing"
	"sync"
	"time"
)

//非线程安全
func TestShareMem(t *testing.T) {
	counter := 0;
	for i := 0; i < 5000; i++ {
		go func() {
			counter++;
		}()	
	}
	time.Sleep(1*time.Second)
	t.Logf("counter=%d",counter)
}

//线程安全的做法
//加锁
func TestShareMemThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0;
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {//防止挂了，这个锁得不到释放导致整个程序挂起
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()	
	}
	time.Sleep(1*time.Second)
	t.Logf("counter=%d",counter)
}

//线程安全的做法
//加锁
func TestShareMemWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0;
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {//防止挂了，这个锁得不到释放导致整个程序挂起
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()	
	}
	wg.Wait()
	t.Logf("counter=%d",counter)
}
