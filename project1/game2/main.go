package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	rand.Seed(time.Now().UnixNano())
	money := 1000
	for {
		fmt.Print("[Input] : ")
		num, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자를 입력하세요.")
		} else if num < 1 || num > 5 {
			fmt.Println("1~5사이의 숫자를 입력하세요.")
		} else {
			n := rand.Intn(5) + 1
			if num == n {
				money += 500
				fmt.Printf("축하합니다. 잔액 : %d\n", money)
				if money >= 5000 {
					fmt.Println("게임 오버")
					break
				}
			} else {
				money -= 100
				fmt.Printf("아쉽습니다. 잔액 : %d\n", money)
				if money <= 0 {
					fmt.Println("게임 오버")
					break
				}
			}
		}
	}
}
