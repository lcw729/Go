package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{Name: name, Age: age}
	return &u // 탈출 분석으로 u 메모리가 사라지지 않음
}

func main() {
	user := NewUser("chaen", 24)

	fmt.Println(user.Name, user.Age)
}
