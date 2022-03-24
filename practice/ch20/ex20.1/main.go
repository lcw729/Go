package main

import "fmt"

type DuckInterface interface { //  interface도 타입이다.
	Fly() // 구현이 빠진 Method들
	Walk(distance int) int
}

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string { // string 함수
	return fmt.Sprintf("Name: %s, Age: %d", s.Name, s.Age)
}

func (s Student) GetAge() int {
	return s.Age
}

func main() {
	student := Student{"chaewon", 24}
	var stringer Stringer

	stringer = student // student라는 structure가 stringer라는 interface를 implement 구현하고 있다.
	fmt.Println(stringer.String())
}
