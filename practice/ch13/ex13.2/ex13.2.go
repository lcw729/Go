package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	UserInfo User
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
		vip.UserInfo.Name,
		vip.UserInfo.ID,
		vip.UserInfo.Age,
		vip.VIPLevel,
		vip.Price)
}
