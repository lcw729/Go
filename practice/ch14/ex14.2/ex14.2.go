package main

import "fmt"

func main() {
	var a = 1
	var b = 1

	var p1 *int = &a
	var p2 *int = &a
	var p3 *int = &b

	fmt.Printf("p1 == p2 : %v\n", p1 == p2)
	fmt.Printf("p1 == p3 : %v\n", p1 == p3)
}
