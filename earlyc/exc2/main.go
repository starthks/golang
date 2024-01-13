package main

/*
#include <stdio.h>
int sum(int a, int b)
{
	return a + b;
}
void hello()
{
	printf("Hello, world \n");
}
*/
import "C"
import "fmt"

func main() {
	var a, b int = 1, 2
	s := C.sum(C.int(a), C.int(b))
	fmt.Println(s)

	C.hello()
}
