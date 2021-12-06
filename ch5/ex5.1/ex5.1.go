package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var f float64 = 34209859282.43223

	fmt.Print("a:", a, "b:", b)
	fmt.Println("a:", a, "b:", b, "f:", f)     // 출력값 사이 공란 + 출력이 끝나면 개행
	fmt.Printf("a: %d b: %d f: %f\n", a, b, f) // 주어진 사용자 서식에 맞춰 입력값 출력
}
