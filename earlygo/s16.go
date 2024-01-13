package earlygo

// 객체(메모리)를 사용한 후 보관, 일종의 캐시로 성능향상
import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

type Data struct {
	tag    string
	buffer []int
}

func S16() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	pool := sync.Pool{
		New: func() interface{} {
			data := new(Data)
			data.tag = "new"
			data.buffer = make([]int, 10)
			fmt.Println("new data", data)
			return data
		},
	}
	for i := 0; i < 10; i++ {
		go func(n int) {
			data := pool.Get().(*Data) //pool에서 *Data타입으로 데이터를 가져옴
			for index := range data.buffer {
				data.buffer[index] = rand.Intn(100)
			}
			fmt.Println("rand: ",n,"data: ", data)
			data.tag = "used"
			pool.Put(data)
		}(i)
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			data := pool.Get().(*Data) //pool에서 *Data타입으로 데이터를 가져옴
			n := 0
			for index := range data.buffer {
				n += i
				data.buffer[index] = n
			}
			fmt.Println("num: ",i,"data: ", data)
			data.tag = "used"
			pool.Put(data)
		}(i)
	}
	Rangetest()
	fmt.Scanln()
}
func Rangetest(){
	sl := []int{10,20,30,40}
	for v := range sl {
		fmt.Println(v)
	}
}
