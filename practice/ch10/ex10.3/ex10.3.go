package main

import "fmt"

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("첫째날입니다.")
	case 2:
		fmt.Println("둘째 날입니다.")
	case 3:
		fmt.Println("셋째 날입니다.")
	default:
		fmt.Println("마지막 날입니다.")
	}
}
