package main

import "fmt"

func main() {
	var score [5]int = [5]int{90, 98, 76, 78, 34}
	var avg float64
	total := 0

	for i := 0; i < 5; i++ {
		total += score[i]
	}

	avg = float64(total) / 5.0
	fmt.Println("average: ", avg)
}
