package earlygo

import (
	"fmt"
	"regexp"
)

func S25() {
	re1 := "[0-9]"
	st1 := "Hello 100"
	ma, _ := regexp.MatchString(re1, st1 )
	fmt.Println(ma)
}
