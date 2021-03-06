package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreateAt  time.Time `json:"create_at"`
}

type fooHandler struct{} // instance 생성

func indexHandler(rw http.ResponseWriter, r *http.Request) { // handler 등록
	fmt.Fprint(rw, "Hello world")
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { // 인터페이스 등록
	user := new(User)
	// Decode r.Body to struct
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	user.CreateAt = time.Now()
	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(data))
}

func barHandler(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World!"
	}
	fmt.Fprintf(rw, "Hello %s", name)
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/bar", barHandler)

	mux.Handle("/foo", &fooHandler{})
	return mux
}
