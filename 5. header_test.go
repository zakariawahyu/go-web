package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/**
Header
- Selain Query Parameter, dalam HTTP ada juga yang namanya Header
- Header adalah informasi tambahan yang biasa dikirim dari client ke server atau sebaliknya
- Jadi dalam Header, tidak hanya ada pada HTTP Request, pada HTTP Response pun kita bisa menambahkan informasi header
- Saat kita menggunakan browser, biasanya secara otomatis header akan ditambahkan oleh broeser, sepeeti infomasi browser,
jenis tipe content yang dikirim dan diterima oleh browser, dan lain-lain

Request Header
- Untuk menangkap request header yang dikirim oleh client, kita bisa mengambilnya di Request.Header
- Header mirip seperti Query Parameter, isinya adalah map[string][]string
- Berbeda dengan query parameter yang case sensitive, secara spesifikasi, Header key tidaklah case sensitive
*/

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type")

	fmt.Fprint(writer, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(request.RequestURI)
	fmt.Println(string(body))
}

/**
Response Header
- Sedangkan jika kita ingin menambahkan header pada response, kita bisa menggunakan function ResponseWriter.Header()
*/

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-Powered-By", "Zakaria Wahyu")

	fmt.Fprint(writer, "OK")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.Header)
	fmt.Println(response.Header.Get("X-Powered-By"))
	fmt.Println(string(body))
}
