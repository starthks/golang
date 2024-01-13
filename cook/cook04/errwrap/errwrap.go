package errwrap

import (
	"fmt"

	"github.com/pkg/errors"
)

func WrappedError(e error) error {
	return errors.Wrap(e, "An error") // /home/start/go/src
}
type ErrorTyped struct{
	error
}
func Wrap() {
	e := errors.New("starndard error")
	fmt.Println("Regula Error -", WrappedError(e))
	fmt.Println("Typed Error -", WrappedError(ErrorTyped{errors.New("typed error")}))
}
