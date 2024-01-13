package main

import (
	"cook/cook05/dns" 
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <address>\n", os.Args[0])
		os.Exit(1)
	}
	// os.Args[1] = []string{"golang.org"}[0]
	address := os.Args[1]
	lookup, err := dns.LookupAddress(address)
	if err != nil {
		log.Panicf("failed to lookup: %s", err.Error())
	}
	fmt.Println(lookup.String()) //객체를 문자열로 출력할때 자동호출
}
