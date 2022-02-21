package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

func main() {
	student := Student{Id: 0, Name: "ccc", Age: 15, Score: 78}
	jsonData, _ := json.Marshal(student)
	buff := bytes.NewBuffer(jsonData)
	req, err := http.NewRequest("POST", "http://localhost:3000/students", buff)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(response)
		println(str)
	}
}
