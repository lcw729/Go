package main

import "fmt"

type User struct {
	Name  string
	ID    string
	Age   int
	Level int // User의 Level 필드
}

type VIPUser struct {
	User  // Level 필드를 갖는 구조체
	Price int
	Level int // VIPUser의 Level 필드
}

func main() {
	var user = User{"chaewon", "chean", 24, 10}
	vip := VIPUser{
		user,
		250,
		3,
	}

	fmt.Printf("[일반 유저] 이름: %s, 닉네임: %s, 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("[VIP 유저] 이름: %s, 닉네임: %s, 나이: %d, 가격: %d만원, VIP 레벨: %d, 유저 레벨: %d",
		vip.Name, // . 하나로 접근 가능
		vip.ID,
		vip.Age,
		vip.Price,
		vip.Level,      // VIPUser의 Level
		vip.User.Level, // 포함된 구조체명을 쓰고 접근
	)
}
