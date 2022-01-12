package main

import "fmt"

func main() {
	var age = 35

	if age >= 10 && age <= 15 {
		fmt.Println("you are young")
	} else if age > 30 || age < 20 {
		fmt.Println("you are not 20s")
	}
}
