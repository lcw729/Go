package main

import "fmt"

func HasRichFriend() bool {
	return true
}

func GetFriendCount() int {
	return 3
}

func main() {
	price := 35000

	if price > 50000 {
		if HasRichFriend() {
			fmt.Println("신발끈 묶어야지~")
		} else {
			fmt.Println("n빵")
		}
	} else if price >= 30000 {
		if GetFriendCount() <= 3 {
			fmt.Println("n빵")
		} else {
			fmt.Println("신발끈 묶어야지~")
		}
	} else {
		fmt.Println("내가 쏜다.")
	}
}
