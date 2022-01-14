package main

import "fmt"

func main() {
	i := 2
	j := 1

	for {
		for {
			if j == 10 {
				break // 안쪽 for문을 종료
			}
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
			j++
		}
		j = 1
		i++
		if i == 10 {
			break
		}
	}
}
