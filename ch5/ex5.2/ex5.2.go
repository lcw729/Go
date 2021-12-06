package main

import "fmt"

// 최소 출력 너비 지정
func main() {
	var a = 123
	var b = 456
	var c = 123456789

	fmt.Printf("%5d, %5d\n", a, b)    // 최소 5칸을 사용해서 정숫값 출력
	fmt.Printf("%05d, %05d\n", a, b)  // 최소 5칸을 사용하고 공란에 0을 채움
	fmt.Printf("%-5d, %-05d\n", a, b) // 마이너스 -를 붙이면 왼쪽을 기준삼아 정렬
	// 출력 값 123  , 456
	// 왼쪽 정렬의 경우 공란에 0이 들어가면 값이 바뀌기 때문에
	// 45600으로 출력되지 않았다.

	fmt.Printf("%5d, %5d\n", c, c)
	fmt.Printf("%05d, %05d\n", c, c)
}
