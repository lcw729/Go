package main

import "fmt"

func getMyAge() int {
	return 24
}

func main() {

	switch age := getMyAge(); age {
	case 10:
		fmt.Println("Teenage")
	case 33:
		fmt.Println("Pair 3")
	default:
		fmt.Println("my age is", age)
	}

	// fmt.Println("my age is", age) // Error - age 변수는 사라졌다.
}
