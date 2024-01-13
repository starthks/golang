package main
// 구조체 슬라이스 정렬
import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}
type Students []Student //[]Student의 별칭 타입 Students
// sort함수 사용 Len(), Less(), Swap() 세 메서드가 필요
func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func main() {
	s := Students{
		{"백두산", 30}, {"나무", 100}, {"사과", 50}, {"귤", 10},
	}
	sort.Sort(Students(s))
	fmt.Println(s)
}
