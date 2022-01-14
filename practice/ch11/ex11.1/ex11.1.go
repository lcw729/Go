package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ { // 초기문; 조건문; 후처리
		fmt.Print(i, ", ")
	}
	fmt.Println()

	var i = 1
	for ; i < 5; i++ {
		fmt.Print(i, ", ")
	}
	fmt.Println()

	for j := 0; j < 5; {
		fmt.Print(i, ", ")
		j++
	}
}
