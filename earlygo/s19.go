package earlygo

import (
	"fmt"
	"reflect"
)

type Intdata struct {
	a,b int
}
func S19(){
	var num int = 1
	fmt.Println(reflect.TypeOf(num))
	data := Intdata{1,2}
	ret := reflect.TypeOf(data.a)
	rev := reflect.ValueOf(data.a)
	fmt.Println("reflect.TypeOf ", ret.Name(), ret.Size(), ret.Kind() )
	fmt.Println("reflect.ValueOf ", rev.Type(), rev.Kind())

	var a *int = new(int)
	*a = 1
	retp := reflect.TypeOf(a)
	revp := reflect.ValueOf(a)
	fmt.Println("reflect.TypeOf ", retp )
	fmt.Println("reflect.ValueOf ", revp.Elem(), revp.Elem().Int())


}