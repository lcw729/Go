package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	User     // 필드명 생략
	VIPLevel int
	Price    int
}

func main() {
	var user = User{"chaewon", "chean", 24}
	vip := VIPUser{
		user,
		3,
		280,
	}

	fmt.Printf("[일반 유저] 이름: %s, 닉네임: %s, 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("[VIP 유저] 이름: %s, 닉네임: %s, 나이: %d, 레벨: %d, 가격: %d만원",
		vip.Name, // . 하나로 접근 가능
		vip.ID,
		vip.Age,
		vip.VIPLevel,
		vip.Price)
}
