package main

import "fmt"

func main() {
	var slice []int

	for i := 0; i < 10; i++ {
		slice = append(slice, i)
	}

	slice = append(slice, 10, 11, 12, 13, 14, 15)
	fmt.Println(slice)
}
