package earlygo

import (
	"fmt"
	"time"
)

func S10() {
	mychanel()

	c := make(chan int)

	for i := 0; i < 5; i++ {
		go input(i, c)
	}

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
func input(i int, c chan<- int) {
	fmt.Println("go func", i)
	c <- i 
}

func mychanel() { //colse ~ range
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) //종료채널
	}()

	for c := range ch { //range 정상종료
		fmt.Println("range ", c)
	}
}
