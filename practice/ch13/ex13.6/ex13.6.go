package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age   int32   // 4바이트
	Score float64 // 8바이트
}

func main() {
	user := User{23, 55.7}
	fmt.Println(unsafe.Sizeof(user))
}
