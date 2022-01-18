package main

import "fmt"

func main() {
	str := "Hello 월드"
	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		fmt.Printf("타입명: %T 값: %d 문자값: %c\n", runes[i], runes[i], runes[i])
	}
}
