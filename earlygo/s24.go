package earlygo

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func S24() {
	fmt.Println(strings.Count("hello hey", "he")) //2
	s := "hello heyyou bye"
	in := strings.Index(s, "hey")
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
	fmt.Printf("%#v\t %#v\n", num1, s1) //100      "100"

	var sb []byte = make([]byte, 0)
	sb = strconv.AppendFloat(sb, 1.2, 'f', -1, 32)
	sb = strconv.AppendInt(sb, 3, 10)
	for _, v := range sb {
		fmt.Printf("%c ", v) //1 . 2 3  
	}
}
