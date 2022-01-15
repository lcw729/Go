package main

import "fmt"

const (
	Computer int = iota
	English
	Art
	Soccer
)

type student struct { // student 구조체 정의
	Name  string
	Class int
	No    int
	Score float64
}

func ClassToString(class int) string {
	if class == Computer {
		return "Computer"
	} else if class == English {
		return "English"
	} else if class == Art {
		return "Art"
	} else if class == Soccer {
		return "Soccer"
	}
	return ""
}

func main() {
	var jenny student    // 구조체 변수 선언
	jenny.Name = "jenny" // 각 필드값 초기화
	jenny.Class = 0
	jenny.No = 12345000
	jenny.Score = 100

	fmt.Printf("이름 : %s\n", jenny.Name) // 필드값을 출력
	fmt.Printf("전공 : %s\n", ClassToString(jenny.Class))
	fmt.Printf("학번 : %d\n", jenny.No)
	fmt.Printf("점수 : %f\n", jenny.Score)
}
