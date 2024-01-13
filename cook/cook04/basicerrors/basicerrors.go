package basicerrors

import (
	"errors"
	"fmt"
)

var ErrorValue = errors.New("this is a typed error")

type TypedError struct {
	error
}

func BasicErrors() {
	err := errors.New("this is a quick and easy way to create an error")
	fmt.Println("errors.New: ", err)

	fmt.Println("ErrorValue occur: ", ErrorValue)
	err = TypedError{errors.New("typed error")}
	fmt.Println("typed error: ", err)
}
