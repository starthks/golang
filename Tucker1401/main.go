package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(n string, a int) *User {
	var u = User{n, a}
	return &u //escape analysis로 힙 메모리에 할당
}
func main() {
	userPointer := NewUser("AAA", 23)
	fmt.Println(userPointer)
}
