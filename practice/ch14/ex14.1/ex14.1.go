package main

import "fmt"

func main() {
	var n int = 5
	var nPtr *int

	nPtr = &n
	fmt.Println("n의 값: ", n)
	fmt.Println("n의 메모리 주소: ", &n)
	fmt.Println("nPtr의 값: ", nPtr)
	fmt.Println("nPtr이 가리키는 값: ", *nPtr)

	*nPtr = 30
	fmt.Println("n의 변경된 값: ", *nPtr)
}
