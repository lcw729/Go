package main

import "fmt"

var cnt = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {
	if false && IncreaseAndReturn() { // 함수가 호출되지 않는다.
		fmt.Println("1 증가")
	}

	if true && IncreaseAndReturn() { //  함수가 호출되지 않는다.
		fmt.Println("2 증가")
	}
}
