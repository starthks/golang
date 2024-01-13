package main

import "fmt"

func main() {
	st := "012"
	fmt.Printf("%#t %#t\n", int(st[1]-'0'), '0')
}
