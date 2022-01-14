package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{6, 7, 8, 9, 10}

	for i, num := range a {
		fmt.Printf("a[%d] = %d\n", i, num)
	}

	for i, num := range b {
		fmt.Printf("b[%d] = %d\n", i, num)
	}

	b = a
	for i, num := range b {
		fmt.Printf("b[%d] = %d\n", i, num)
	}

}
