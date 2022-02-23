//ch30/ex30.1/ex30_1_test.go
package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
 "strings"
	"github.com/stretchr/testify/assert"
)

func TestJsonHandler3(t *testing.T) {
	assert := assert.New(t)

	var student Student
	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/students",
		strings.NewReader(`{"Id":0,"Name":"ccc","Age":15,"Score":78}`)) // ❶ 새로운 학생 데이터

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusCreated, res.Code) // ❷ 응답 코드 검사

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/1", nil) // ❸ 추가된 학생 데이터
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("ccc", student.Name)
}