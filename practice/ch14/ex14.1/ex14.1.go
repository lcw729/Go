package main

import "fmt"

func main() {
	var a int
	var p *int // 포인터 변수 선언

	p = &a // a의 메모리 주소를 포인터 변수에 할당

	*p = 20

	fmt.Printf("a의 값은 %d입니다.\n", a)
	fmt.Printf("p의 값 : %p\n", p)
	fmt.Printf("p가 가리키는 메모리의 값 : %d \n", *p)

	*p = 100
	fmt.Printf("a의 값은 %d입니다.\n", a)
}
