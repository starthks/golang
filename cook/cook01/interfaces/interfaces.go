package interfaces

import (
	"fmt"
	"io"
	"os"
)

func Copy(in io.Reader, out io.Writer) error {
	w := io.MultiWriter(out, os.Stdout)

	// 표준복사, 매개변수 in에 대량의 데이터가 있는 경우 이 방법은 위험할 수 있다.
	if _, err := io.Copy(w, in); err != nil {
		return err
	}

	// 64바이트 청크(chunk)를 사용
	// buf := make([]byte, 64)
	// if _, err := io.CopyBuffer(w, in, buf); err != nil {
	// 	return err
	// }
	fmt.Println()
	return nil
}
