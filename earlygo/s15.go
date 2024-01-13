package earlygo
import (
	"fmt"
	"runtime"
	"sync"
)
func S15() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	once := new(sync.Once)
	for i:=0 ;i<3;i++{
		go func(n int){
			fmt.Println("goroutine: ",n)
			once.Do(hello)
		}(i)
	}
	fmt.Scanln()
}
func hello(){
	fmt.Println("Hello")
}
