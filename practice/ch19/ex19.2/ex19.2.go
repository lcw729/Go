package main

import "fmt"

// 사용자 정의 별칭 타입
type myInt int

// myInt 별칭 타입을 리시버로 갖는 메서드
func (a myInt) add(b int) int {
	return int(a) + b
}

func main() {
	var a myInt
	fmt.Println(a.add(30))
	var b int = 20
	fmt.Println(myInt(b).add(50))
}
