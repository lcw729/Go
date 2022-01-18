package main

import "fmt"

func main() {
	str1 := "Hello \t 'world' \n"
	fmt.Println(str1)

	str2 := `Hello \t 'world' \n Testing 
back qoute`

	fmt.Println(str2)
}
