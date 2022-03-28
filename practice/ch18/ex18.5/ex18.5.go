package main

import "fmt"

func main() {
	slice1 := make([]int, 3, 5)
	slice2 := append(slice1, 4)

	fmt.Println(slice1, len(slice1), cap(slice1))
	fmt.Println(slice2, len(slice2), cap(slice2))

	slice1[2] = 200
	fmt.Println(slice1, len(slice1), cap(slice1))
	fmt.Println(slice2, len(slice2), cap(slice2))

	slice2 = append(slice1, 5)

	fmt.Println(slice1, len(slice1), cap(slice1))
	fmt.Println(slice2, len(slice2), cap(slice2))
}
