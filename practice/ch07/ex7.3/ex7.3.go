package main

import "fmt"

func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func main() {
	c, success := Divide(3, 1)
	fmt.Println(c, success) // 3, true
	c, success = Divide(2, 0)
	fmt.Println(c, success) // 0, false
}
