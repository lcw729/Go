package main

import "fmt"

func main() {
	const PI1 float64 = 3.14159
	var PI2 float64 = 3.14159

	PI2 = 4
	fmt.Println(PI1)
	fmt.Println(PI2)
}
