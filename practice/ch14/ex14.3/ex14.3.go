package main

import "fmt"

func changeNumber(arg int) {
	arg = 99
}

func main() {
	var a = 9

	fmt.Println("changeNumer() 이전: ", a)
	changeNumber(a)
	fmt.Println("changeNumer() 이후: ", a)
}
