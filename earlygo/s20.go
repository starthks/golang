package earlygo

import (
	"fmt"
	"reflect"
)

func S20() {
	var intsum func(int, int) int64
	var floatsum func(float32, float32) float64
	var stringSum func(string, string) string

	makeSum(&intsum)
	valueint := intsum(2, 3)

	makeSum(&floatsum)
	valuefloat := floatsum(2.1, 3.5)

	makeSum(&stringSum)
	valuestring := stringSum("Hello", "world")

	fmt.Println(valueint, valuefloat, valuestring)    //5 5.599999904632568 Helloworld
	fmt.Println(float64(float32(2.1) + float32(3.5))) //5.6

}
func makeSum(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()
	v := reflect.MakeFunc(fn.Type(), rsum)
	fn.Set(v)
}
func rsum(args []reflect.Value) []reflect.Value {
	a, b := args[0], args[1]
	if a.Kind() != b.Kind() {
		fmt.Println("타입이 다릅니다..")
		return nil
	}
	switch a.Kind() {
	case reflect.Int, reflect.Int64:
		return []reflect.Value{reflect.ValueOf(a.Int() + b.Int())}
	case reflect.Float32, reflect.Float64:
		return []reflect.Value{reflect.ValueOf(a.Float() + b.Float())}
	case reflect.String:
		return []reflect.Value{reflect.ValueOf(a.String() + b.String())}
	default:
		return []reflect.Value{}
	}
}
