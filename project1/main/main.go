package main

/*
  숫자 맞추기 게임
  1. 랜덤 값 생성하기
  2. 입력 값 받기
  3. 입력 값과 랜덤 값 비교하기
  4-1. 입력 값과 랜덤 값이 같다면 시도 횟수와 문구 출력 + break
  4-2. 다르다면 입력 값이 랜덤 값보다 큰지 작은지 출력
*/

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

	num := rand.Intn(100) // 랜덤 값 생성

	var count = 1
	for {
		fmt.Print("[Input] : ")
		n, err := InputIntValue() // 숫자 값 입력
		if err != nil {
			fmt.Println("[Error] 숫자를 입력하세요.")
			continue
		} else {
			if n > num { // 숫자 값 비교
				fmt.Println("입력하신 숫자가 랜덤 값보다 큽니다.")
			} else if n < num {
				fmt.Println("입력하신 숫자가 랜덤 값보다 작습니다.")
			} else if n == num {
				fmt.Println("축하합니다! 랜덤 값을 맞췄습니다.")
				fmt.Printf("시도 횟수 : %d", count)
				break
			}
			count++
		}
	}
}
