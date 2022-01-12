package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/**
HTTP Test
- Go-Lang sudah menyediakan package khusus untuk membuat unit test terhadap fitur web yang kita buat
- Semua ada di dalam package net/http/httptest
- https://golang.org/pkg/net/http/httptest/
- Dengan menggunakan package ini, kita bisa melakukan testing handler web di Go-Lang tanpa harus menjalankan aplikasi webnya
- Kita bisa langsung fokus terhadap handler function yang akan kita test

httptest.NewRequest()
- NewRequest(method, url, body) merupakan function yang digunakan untuk membuat http.request
- KIta bisa menentukan method, url dan body yang akan kita kirim sebagai simulasi unit test
- Selain itu, kita juga bisa menambahkan informasi tambahan lainnya pada request yang ingin kita kirim seperti header, cookie dll

http.NewRecorder()
- Ada 2 parameter yang digunakan dalam function handler, parameter writer dan parameter request
- http.NewRecorder() ini direpresentasikan menggantikan writer
- http.NewREcoder() merupakan function yang digunakan untuk membuat ResponseRecorder
- ResponseRecorder merupakan struct bantuan untuk merekam HTTP Response dari hasil testing yang kita lakukan
*/

func HelloHanlder(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHanlder(recorder, request)

	// mengecek hasil test
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	bodyString := string(body)
	if err != nil {
		panic(err)
	}
	fmt.Println(bodyString)
	fmt.Println(response.Status)
}
