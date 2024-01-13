package earlygo

import (
	"fmt"
	"runtime"
	"sync"
)

func S17() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println("1st go",n)
			for j :=0; j<3;j++{
				// wg.Add(1)
				// go func(n,j int){
					fmt.Printf("\t%d * %d = %d\n", n,j,n*j)
					// wg.Done()
				// }(n,j)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("the end")
}
