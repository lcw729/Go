package main

import "fmt"

// go언어는 강타입 언어
func main() {
	var a int = 3
	var b float64 = 3.5
	c := a * int(b)
	d := float64(a) * b
	// c := a * b -- a와 b는 타입이 다르기 떄문에 연산이 불가능하다.

	fmt.Println(c)
	fmt.Println(d)
}
