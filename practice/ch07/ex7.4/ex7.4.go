package main

import "fmt"

func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return // 명시적으로 반환할 값을 지정하지 않은 return문
	}
	result = a / b
	success = true
	return
}

func main() {
	c, success := Divide(3, 1)
	fmt.Println(c, success) // 3, true
	c, success = Divide(2, 0)
	fmt.Println(c, success) // 0, false
}
