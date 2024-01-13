package main

import (
	"fmt"
	"unsafe"
)

// 레지스터 크기 4바이트(한번 연산 최소단위)인 32비트
// 레지스터 크기 8바이트인 64비트, 메모리크기가 같게 되면 효율적
// -> 메모리 패딩을 줄이기 위해 8바이트단위로 변수 위치 조정
type User struct {
	Name string
	ID   string
	Age  int
}
type VIPUser struct {
	User     // 포함된 필드
	VIPLevel int32 //int32 4byte, int 8byte
	Price    int32
}

func main() {
	vip := VIPUser{
		User : User{Name: "하나", ID: "hana"}, // 포함된 필드
	    VIPLevel: 3,
		Price: 250, //마지막 쉼표
	}
	vip.Age = 30
	fmt.Println(unsafe.Sizeof(vip.User.Name),unsafe.Sizeof(vip.User), unsafe.Sizeof(vip))
	fmt.Printf("User : %s, %s, %d, VIP Level : %d, %d원\n",
		vip.Name, vip.ID, vip.Age,
		vip.VIPLevel, vip.Price)
}
