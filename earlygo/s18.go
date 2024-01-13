package earlygo

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func S18() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var data int32 = 0
	wg := new(sync.WaitGroup)
	// mutex := new(sync.Mutex) //서로 상호보안하나 오차발생-> atomic사용

	for i := 0; i < 200000; i++ {
		// mutex.Lock()
		wg.Add(1)
		go func() {
			// data += 1
			atomic.AddInt32(&data, 1)
			wg.Done()
		}()
		// mutex.Unlock()
		// runtime.Gosched()
	}

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			// data -= 1
			atomic.AddInt32(&data, -1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(data)
}
