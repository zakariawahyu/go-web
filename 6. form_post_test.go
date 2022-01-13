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
Form Post
- Saat kita belajar HTML, kita tahu bahwa saat kita membuat form, kita bisa submit datanya dengan method GET atau POST
- Jika menggunakan method GET, maka hasilnya semua data di form akan menjadi Query Parameter
- Sedangkan jika menggunakan POST, maka semua data di form akan dikirim via boddy HTTP request

Request.PostForm
- Smua data form post yang dikirim dari client, secara otomatis akan disimpan dalam attribute Request.PostForm
- Namun sebelum kita bisa mengambil data di attribute PostForm, kita wajib memanggil method Request.ParseForm() terlebih dahulu,
method ini digunakan untuk melakukan parsing data body, apakah bisa di parsing menjadi form atau tidak, jika tidak bisa di parsing
maka akan menyebabkan error
*/

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.PostForm
	if err != nil {
		panic(err)
	}

	firstName := request.PostFormValue("firstName")
	lastName := request.PostFormValue("lastName")

	fmt.Fprintf(writer, "My name %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("firstName=Zakaria&lastName=Wahyu")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
