package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)  ///usr/lib/golang/src/cook

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	mess := make(chan string, 5)
	wg := new(sync.WaitGroup)
	f("direct")

	wg.Add(1)
	go func() {
		time.Sleep(time.Second*1)
		f("goroutine")
        defer wg.Done()
	}()

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(msg int) {
            fmt.Println(msg)
			mess <- "test" + strconv.Itoa(msg)
			time.Sleep(1000)
			defer wg.Done()
		}(i)

	}
	for i := 1; i <= 5; i++ {
        fmt.Println(<-mess )
	}
	wg.Wait()
	fmt.Println("done")
}
