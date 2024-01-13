package main

import (
	"bytes"
	"cook/cook01/interfaces"
	"fmt"
)

func main() {
	// fi, err := os.Open("/home/start/go/src/cook/cook01/interfaces/example/exam.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer fi.Close()
	// filebyte, err:=io.ReadAll(fi)
	// if err != nil {
	// 	panic(err)
	// }
	// in := bytes.NewReader(filebyte)
	
	in := bytes.NewReader([]byte("example"))
	out := &bytes.Buffer{}
	fmt.Println("stdout on Copy = ")
	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}
	fmt.Println("out bytes buffer =", out.String())
	fmt.Println("=======================")

	fmt.Println("stdout on PipeExample =")
	if err := interfaces.PipeExample(); err != nil {
		panic(err)
	}
}
