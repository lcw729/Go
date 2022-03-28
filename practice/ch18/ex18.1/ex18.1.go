package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	var slice2 = []int{1, 5: 3, 10: 7}
	var slice3 = make([]int, 3)

	if len(slice) == 0 {
		fmt.Println("slice is empty")
	}

	slice[1] = 10
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
