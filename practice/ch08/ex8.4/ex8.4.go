package main

import "fmt"

func main() {
	const PI = 3.14 // 타입없는 상숫값
	const FloatPI float64 = 3.14

	var a int = PI * 100
	// var b int = int(FloatPI) * 100 // 타입 오류 발생

	fmt.Println(a)
}
