package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{} // instance 생성

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { // 인터페이스 등록
	fmt.Fprintf(w, "Hello Foo!")
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) { // handler 등록
		fmt.Fprint(rw, "Hello world")
	})

	http.HandleFunc("/bar", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello bar")
	})

	http.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3000", nil)
}
