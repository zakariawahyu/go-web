package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/**
Query Parameter
- Query parameter adalah salah satu fitur yang bisa kita gunakan ketika membuat web
- Query parameter biasanya digunakan untuk mengirim data dari client ke server melalui URL
- Query parameter ditempatkan pada URL
- Untuk menambahkan Query Paramter, kita bisa menggunakan ?namaQuery=valueQuery pada URLnya

url.URL
- Dalam parameter Request, terdapat atribut URL yang berisi data url.URL
- Dari data URL ini, kita bisa mengambil data query parameter yang dikirim dari client dengan menggunakan method Query() yang kan mengembalikan map
*/

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Zakaria", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(request.RequestURI)
	fmt.Println(string(body))
}

/**
Multiple Query Parameter
- Dalam spesifikasi URL kita bisa menambahkan lebih dari satu query parameter
- Ini cocok sekali jika kita memang ingin mengirim banyak data ke server, cukup tambahkan query parameter lainnya
- Untuk menambahkan query parameter, kita bisa gunakan tanda & lalu diikuti dengan query parameter berikutnya
*/

func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("firstname")
	lastName := request.URL.Query().Get("lastname")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?firstname=Zakaria&lastname=Wahyu", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(request.RequestURI)
	fmt.Println(string(body))
}

/**
Multiple Value Query Parameter
- Sebernarnya URL melakukan parsing query parameter dan menyimppan dalam map[string][]string
- Artinya dalam satu key query parameter, kita bisa memasukkan beberapa value
- Caranya kita bisa menambahkan query parameter dengan nama yang sala tapi value yang berbeda
- misal : name=Zakaria&nameWahyu
- Akan tetapi jika menggunakan query Get() akan mengambil data pertamanya saja
- Jika kita ingin mengambil semua value kita harus menggunakan map []string
*/

func MultipleParameterValues(writer http.ResponseWriter, request *http.Request) {
	result := request.URL.Query()

	name := result["name"]
	fmt.Fprintf(writer, strings.Join(name, " "))
}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Zakaria&name=Wahyu&name=Nur", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(request.RequestURI)
	fmt.Println(string(body))
}
