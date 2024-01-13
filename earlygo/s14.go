package earlygo

import (
	"fmt"
	"runtime"
	"sync"
)

func S14() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var mutex = new(sync.Mutex)    // 뮤텍스 생성
	var cond = sync.NewCond(mutex) // 뮤텍스를 이용하여 조건 변수 생성

	c := make(chan bool, 3) // 비동기 채널 생성

	for i := 0; i < 3; i++ {
		go func(n int) {                    
			mutex.Lock()                    
			c <- true                       
			fmt.Println("wait begin : ", n)
			cond.Wait()                     // 조건 변수 대기
			fmt.Println("wait end : ", n)
			mutex.Unlock()                  

		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c // 고루틴 3개가 모두 실행될 때까지 기다림
	}

	for i := 0; i < 3; i++ {
		mutex.Lock()                
		fmt.Println("signal : ", i)
		cond.Signal()               
		mutex.Unlock()             
	}

	fmt.Scanln()
}
