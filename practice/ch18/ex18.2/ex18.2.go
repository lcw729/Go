package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3}

	fmt.Println(slice)
	for i, v := range slice {
		slice[i] = v * 2
	}

	fmt.Println(slice)

	slice2 := append(slice, 4)
	fmt.Println(slice2)
}
