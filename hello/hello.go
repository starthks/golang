package main

import "fmt"

func main() {
   st := [][]int{{1,2,3},{4,5,6}}
   for _, s := range st {
	fmt.Print(s,"\t")
   }
   fmt.Println()
}
