package main

import "fmt"

func main() {
	var x int8 = 16   // 0001 0000
	var y int8 = -128 // 1000 0000
	var z int8 = -1   // 1111 1111
	var w uint8 = 128 // 1111 1111

	fmt.Printf("x:%08b x>>2: %08b x>>2: %d\n", x, x>>2, x>>2) // 0000 0100 : 4
	fmt.Printf("y:%08b y>>2: %08b y>>2: %d\n", y, y>>2, y>>2) // 1110 0000 : -32
	fmt.Printf("z:%08b z>>2: %08b z>>2: %d\n", z, z>>2, z>>2) // 1111 1111 : -1
	fmt.Printf("w:%08b w>>2: %08b w>>2: %d\n", w, w>>2, w>>2) // 0011 1111 : 32
}
