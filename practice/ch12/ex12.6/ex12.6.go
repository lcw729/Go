package main

import "fmt"

func main() {
	a := [2][5]int{ // 이중 배열 선언
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	for _, arr := range a { // arr값은 순서대로 a[0]의 배열 a[1]의 배열
		for _, v := range arr { // v의 값은 순서대로 a[0]와 a[1]배열의 각 요소
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}
