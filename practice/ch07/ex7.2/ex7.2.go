package main

import "fmt"

func PrintAvgScore(math int, history int, eng int, name string) {
	avg := (math + history + eng) / 3
	fmt.Printf("%s님의 평균 점수는 %d입니다.\n", name, avg)
}

func main() {
	PrintAvgScore(95, 90, 85, "chaewon")
}
