package main

import "fmt"

func getMyAge() int {
	return 24
}

func main() {

	switch age := getMyAge(); true {
	case age < 10:
		fmt.Println("child")
	case age < 20:
		fmt.Println("Teenager")
	case age < 30:
		fmt.Println("20s")
	default:
		fmt.Println("my age is", age)
	}
}
