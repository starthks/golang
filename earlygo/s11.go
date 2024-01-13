package earlygo

import (
	"fmt"
	"time"
)

func sum(a, b int, c chan<- int) {
	for i := 0; i < 5; i++ {
		i := i //loopclosure go1.22개선
		go func() {
			c <- a + b + i
		}()
	}
}

func S11() {
	c := make(chan int)
	sum(1, 2, c)

	for {
		select {
		case t := <-c:
			fmt.Println("receive ", t)
		case <-time.After(3 * time.Second):
			fmt.Println("3 time over")
			return
		}
	}
}
