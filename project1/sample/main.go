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

	rand.Seed(time.Now().UnixNano()) // 시간 값을 랜덤 시드로 설정

	n := rand.Intn(100)
	fmt.Println(n)

	in, _ := InputIntValue()
	fmt.Println(in)

}
