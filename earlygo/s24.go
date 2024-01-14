package earlygo

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func S24() {
	fmt.Println(strings.Count("hello hey","he")) //2
	s := "hello heyyou bye"
	in := strings.Index(s,"hey")
	fmt.Println(s[in:]) //heyyou bye

	f := func(r rune) bool {
		return unicode.Is(unicode.Hangul, r)
	}
	s = "go 언어"
	in = strings.IndexFunc(s, f)
	fmt.Println(s[in:]) //언어

	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")              // 바꿀 문자열을 지정
	fmt.Println(r.Replace("<div><span>Hello, world!</span></div>")) // HTML에서 < >를 &lt; &gt;로 바꿈

	num1, _ := strconv.Atoi("100")
	s1 := strconv.Itoa(100)
	fmt.Printf("%#v\t %#v\n", num1, s1)

	
}
