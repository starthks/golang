package earlygo

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func S23() {
	fmt.Println("마크문자:\u17c9, \u002a, \U0001f602") //마크문자:៉, *, 😂
	// https://charbase.com/1f602-unicode-face-with-tears-of-joy
	fmt.Println(unicode.IsDigit('1'))
	fmt.Println(unicode.IsLetter('1')) //false
	fmt.Println(unicode.IsGraphic('\u17c9')) //화면에 표시

	var s string = "한a글"
	fmt.Println(len(s), utf8.RuneCountInString(s)) // 7 3

	b := []byte("한a글")
	r, size := utf8.DecodeRune(b)
	fmt.Println(len(b), utf8.Valid(b), "rune:", utf8.RuneLen(r), size)  
	//7 true rune: 3 3
	r, size = utf8.DecodeRune(b[4:]) //3, 1, 3
	fmt.Printf("%c %d\n", r, size)

}
