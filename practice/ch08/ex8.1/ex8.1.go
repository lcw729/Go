package main

import "fmt"

func main() {
	const C int = 10
	var b int = C * 20
	C = 20          // 에러 발생 - 상수는 대입문 좌변에 올 수 없음
	fmt.Println(&C) // 상수 메모리 주솟값에 접근 불가
}
