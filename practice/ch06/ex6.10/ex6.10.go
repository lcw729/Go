package main

import "fmt"

func main() {
	var a int
	var b int

	a, b = 3, 4
	a, b = b, a
	fmt.Println(a, b)
}
