package main

import "fmt"

func find45(a int) (int, bool) {
	for b := 1; b <= 9; b++ {
		if a*b == 45 {
			return b, true
		}
	}
	return 0, false
}

func main() {

	for a := 1; a <= 9; a++ {
		if b, found := find45(a); found {
			fmt.Println(a, b)
			break
		}
	}
}
