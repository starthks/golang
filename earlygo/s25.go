package earlygo

import (
	"fmt"
	"regexp"
)

func S25() {
	re1 := "[0-9]"
	st1 := "Hello 100 foo"
	ma, _ := regexp.MatchString(re1, st1 )
	fmt.Println(ma)

	re2 := regexp.MustCompile(`foo.?`)
	fmt.Printf("%q\n", re2.Find([]byte(`sefood foo1`)))
	fmt.Printf("%q\n", re2.Find([]byte(st1)))
	
}
