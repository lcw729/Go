package main

import "fmt"

func getMyAge() (int, bool) {
	return 30, true
}

func main() {

	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("you are young", age)
	} else if ok && age < 30 {
		fmt.Println("Nice age", age)
	} else if ok {
		fmt.Println("You are beautiful", age)
	} else {
		fmt.Println("Error")
	}

	// fmt.Println("Your age is", age) // Error - age는 소멸되었음
}
