package main

import "fmt"

type account struct {
	balance int
	name    string
}

func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

func (a3 account) withdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}

func main() {
	var mainA *account = &account{10000, "chaewon"}
	mainA.withdrawPointer(2000)
	fmt.Println(mainA.balance) // 8000

	mainA.withdrawValue(3000)
	fmt.Println(mainA.balance) // 8000

	var mainB account = mainA.withdrawReturnValue(2000)
	fmt.Println(mainB.balance) // 6000

	mainB.withdrawPointer(1000)
	fmt.Println(mainB.balance) // 5000
}
