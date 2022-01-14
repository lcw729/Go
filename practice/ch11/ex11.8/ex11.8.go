package main

import "fmt"

func main() {
	a := 1
	b := 1
OuterFor: // 레이블 정의
	for ; a < 10; a++ {
		for b = 1; b < 10; b++ {
			if a*b == 45 {
				fmt.Println(a, b)
				break OuterFor // 레이블에 가장 먼저 포함된 for문까지 종료
			}
		}
	}
}
