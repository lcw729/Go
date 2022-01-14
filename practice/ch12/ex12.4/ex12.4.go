package main

import "fmt"

func main() {
	myfamily := [4]string{"alice", "james", "justin", "jenny"}

	for i, name := range myfamily {
		fmt.Println(i, " : ", name)
	}
}
