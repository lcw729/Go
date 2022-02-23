package main

import "fmt"

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) { // 일반 함수 표현
	a.balance -= amount
}

func (a *account) withdrawMethod(amount int) { // 메서드 표현
	a.balance -= amount
}

func main() {
	a := &account{100}

	withdrawFunc(a, 30) // 함수 형태 호출

	a.withdrawMethod(30) // 메서드 형태 호출

	fmt.Printf("%d \n", a.balance)

}
