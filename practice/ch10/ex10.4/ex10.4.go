package main

import "fmt"

func main() {
	day := "sunday"

	switch day {
	case "sunday", "saturday":
		fmt.Println("주말입니다!")
	default:
		fmt.Println("평일입니다. 출근하세요")
	}
}
