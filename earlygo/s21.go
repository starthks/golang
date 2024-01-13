package earlygo

import "fmt"

func S21(){
	fmt.Printf("%b\t%b\n", 1, 1.0)

	a := new(int)
	b := new(int)
	// fmt.Scanf("%d,%d", a, b)
	fmt.Println(*a,*b)

	s1 := fmt.Sprintln(1, 2.0, "Hello")
	fmt.Printf("%#v\n", s1)
}