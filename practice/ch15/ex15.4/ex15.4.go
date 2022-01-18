package main

import "fmt"

func main() {
	str := "Hello 월드"
	runes := []rune(str)
	var char rune = 46300

	fmt.Println(str)
	fmt.Println("len(str):", len(str))
	fmt.Println(runes)
	fmt.Println("len(runes):", len(runes))
	fmt.Printf("%c\n", char)
}
