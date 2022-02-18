package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()                                          // ServeMux 인스턴스 생성
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // 인스턴스에 핸들러 등록
		fmt.Fprint(w, "Hello world")
		http.FileServer(http.Dir("static"))
	})
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello bar")
	})

	mux.HandleFunc("/student", StudentHandler)
	return mux
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values.Get("name")
	age, _ := strconv.Atoi(values.Get("age"))
	score, _ := strconv.Atoi(values.Get("score"))
	student := Student{
		name,
		age,
		score,
	}
	data, _ := json.Marshal(student)                   // Student 객체를 []byte로 변환
	w.Header().Add("content-type", "application/json") // JSON 포맷임을 표시
	w.Write(data)                                      // 결과 전송
	fmt.Fprint(w, string(data))
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler()) // mux 인스턴스 사용
}
