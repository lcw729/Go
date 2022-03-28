package main

import "fmt"

func changeArray(array2 [5]int) { // array를 복사한다. 40바이트
	array2[2] = 200
}

func changeSlice(slice2 []int) { // slice struct를 복사한다. 24바이트
	slice2[2] = 200
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArray(array)
	changeSlice(slice)

	fmt.Println(array)
	fmt.Println(slice)
}
