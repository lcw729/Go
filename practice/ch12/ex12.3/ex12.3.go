package main

import "fmt"

func main() {
	nums := [...]int{1, 2, 3, 4, 5}
	nums[3] = 6
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}
