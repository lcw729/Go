package main

import "fmt"

type Student struct {
	Age   int
	No    int
	Score float64
}

func PrintStudent(s Student) {
	fmt.Printf("나이: %d 번호: %d, 점수: %f\n", s.Age, s.No, s.Score)
}

func main() {
	var student = Student{15, 23, 64.3}
	student2 := student    // student 구조체 모든 필드가 student2로 복사된다.
	PrintStudent(student)  // 나이: 15 번호: 23, 점수: 64.300000
	PrintStudent(student2) // 나이: 15 번호: 23, 점수: 64.300000

	student2.No = 34
	PrintStudent(student)  // 나이: 15 번호: 23, 점수: 64.300000
	PrintStudent(student2) // 나이: 15 번호: 34, 점수: 64.300000
}
