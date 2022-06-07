package main

import "fmt"

/*
type Human interface {
	Walk() string
	Talk() string
}

type User struct {
	Name string
}

func (u User) Walk() string {
	return fmt.Sprintln("Walking")
}

func (u User) Talk() string {
	return fmt.Sprintln("Talking")
}
*/
type Student struct {
	Name string
	Age  int
}

func (s Student) Walk() string {
	return fmt.Sprintf("%s can walk", s.Name)
}

func (s Student) GetAge() int {
	return s.Age
}

// Embedded interface
type ReadWriter interface {
	Reader
	Writer
}

type Reader interface {
	Read() string
	// Close() error
}

type Writer interface {
	Write(str string)
	// Close() error
}

type memo struct {
	str string
}

func (m *memo) Write(str string) {
	m.str = str
}

func (m *memo) Read() string {
	return m.str
}

// 타입 변환
type Human interface {
	Learn()
}

type Teacher struct {
	Name string
}

func (m *Teacher) Learn() {
	fmt.Println("Teacher can learn")
}

func (m *Teacher) Teach() {
	fmt.Println("Teacher can teach")
}

func (s *Student) Learn() {
	fmt.Println("Student can learn")
}

func Study(h Human) {
	if h != nil {
		h.Learn()
	}
	var s *Student
	s = h.(*Student)
	fmt.Printf("Name: %v\n", s.Name)
}

func main() {
	m := &memo{str: "NULL"}
	var rw ReadWriter = m
	fmt.Println(rw.Read())

	rw.Write("Hello World")
	fmt.Println(rw.Read())
}
