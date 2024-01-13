package interfaces

import (
	"io"
	"os"
)

func PipeExample() error {
	r, w := io.Pipe()
	// 블록방식으로 동작하기 때문에 별도의 go루틴에서 실행
	go func() {
		fi, err := os.Open("/home/start/go/src/cook/cook01/interfaces/example/exam.txt")
		if err != nil {
			panic(err)
		}
		defer fi.Close()
		filebyte, err := io.ReadAll(fi)
		if err != nil {
			panic(err)
		}
		
		// w.Write([]byte("test\n"))
		w.Write(filebyte)
		w.Close()
	}()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return nil
}
