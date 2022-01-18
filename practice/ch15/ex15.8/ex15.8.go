package main

import "fmt"

func main() {
	// 문자열 합치기
	str1 := "Hello"
	str2 := "world"
	str3 := str1 + " " + str2
	fmt.Println(str3)

	// 문자열 비교하기
	str4 := "world"
	fmt.Println("str1 == str4 :", str1 == str4)
	fmt.Println("str2 == str4 :", str2 == str4)

	// 문자열 대소 비교하기
	str5 := "aaa"
	str6 := "aab"
	fmt.Println("str5 < str6 :", str5 < str6)
}
