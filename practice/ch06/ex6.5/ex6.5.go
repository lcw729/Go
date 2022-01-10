package main

import "fmt"

func main() {
	var x int8 = 127 // 01111111

	fmt.Printf("%d < %d + 1: %v\n", x, x+1, x < x+1) // false
	fmt.Printf("x\t = %4d,  %08b\n", x, x)
	fmt.Printf("x+1\t = %4d,  %08b\n", x+1, x+1) // -10000000 => -128
	fmt.Printf("x+2\t = %4d,  %08b\n", x+2, x+2) // -10000000 + 1 => -127
	fmt.Printf("x+3\t = %4d,  %08b\n", x+3, x+3) // -10000000 + 2 => -126

	var y int8 = -128                                // 10000000
	fmt.Printf("%d > %d - 1: %v\n", y, y+1, y > y-1) // 01111111 = 127
	fmt.Printf("y\t = %4d,  %08b\n", y, y)
	fmt.Printf("y-1\t = %4d, %08b\n", y-1, y-1) // 10000000 - 1 = 01111111
}
