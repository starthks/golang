package earlygo

import (
	"fmt"
	"os"
)

func S22() {
	file1, _ := os.Open("../hello.txt") //earlygo/hello.txt
	defer file1.Close()

	// hello world 12345 한글
	//end
	var s1, s2, s3, s4 string
	var num1 int
	fmt.Fscanf(file1, "%s %s %d %s %s", &s1, &s2, &num1, &s3, &s4)
	fmt.Println(s1, s2, num1, s3, s4)
	fmt.Printf("%T, %#v \n", num1, num1)

}
