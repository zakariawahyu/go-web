package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/**
Response Code
- Dalam HTTP, terdapat yang namanya response code
- Response code merupakan representasi kode response
- Dari response code ini kita bisa melihat apakah semua request yang kita kirim itu sukses di proses oleh server atau gagal
- Ada banyak sekali response code yang bisa kita gunakan saat membuat web
- https://developer.mozilla.org/en-US/docs/Web/HTTP/Status

HTTP Response Status Code
- Informational Response (100 - 199)
- Successful Response (200 - 299)
- Redirects (300 - 399)
- Client Errors (400 - 499)
- Server Errors (500 - 599)

Mengubah Response Code
- Secara default, jika kita tidak menyebutkan response code maka response codenya adalah 200 OK
- Jika kita ingin mengubahnya, kita bisa menggunakan function ResponseWriter.WriteHeader(int)
- Semua data status code juga sudah disediakan di Go-Lang, jadi jika kita ingin kita bisa menggunakan yang sudah disediakan
https://github.com/golang/go/blob/master/src/net/http/status.go
*/

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, "Name is empty")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestResponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/helo", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}

func TestResponseCodeValid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/helo?name=Zakaria", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}
