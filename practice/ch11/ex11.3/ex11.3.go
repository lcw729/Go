package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("입력하세요.")
		var num int
		_, err := fmt.Scanln(&num)
		if err != nil { // 숫자가 아닌 경우
			fmt.Println("숫자를 입력하세요.")

			stdin.ReadString('\n') // 키보드 버퍼를 지워준다.
			continue               // 다시 무한 루프 초기로 돌아간다.
		}

		fmt.Printf("입력하신 숫자는 %d입니다.\n", num)
		if num%2 == 0 { // 짝수 검사
			break // for문을 종료한다.
		}
	}
	fmt.Println("for문이 종료됐습니다.")
}
