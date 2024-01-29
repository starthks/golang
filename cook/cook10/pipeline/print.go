package pipeline

import (
	"context"
	"fmt"
)

func (w *Worker) Print(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case val := <-w.in:
			fmt.Println("p",val)
			w.out <- val
		}
	}
}
