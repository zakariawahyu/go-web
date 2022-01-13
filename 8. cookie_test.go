package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/**
Stateless
- HTTP merupakan stateless antara client dan server, artinya sever tidak akan menyimpan data apapun untuk mengigat setiap request dari client
- Hal ini bertujuan agar mudah melakukan scalability di sisi server
- Scalable disini ketika runing aplikasi web tidak mungkin di satu server, pasti di beberapa server. artinya request pertama bisa saja masuk ke server 1
tapi request kedua belum tentu masuk ke server 1 lagi, bisa ke server 2, 3 atau 4
- Jika tidak stateles maka request akan di simpan terus menerus dalam 1 server dan bikin penuh
- Lantas bagaimana cara agar server bisa mengingat sebuah client? Misal ketika kita sudah login di website, server otomatis harus tahu
jika client tersebut sudah login, sehinga request selanjutnya tidak perlu diminta untuk login lagi
- Untuk melakukan hal ini, kita bisa memanfaatkan Cookie (datanya bukan  disimpan di server, melainkan di client)

Cookie
- Cookie adalah fitur HTTP dimana server bisa memberi response cookie(key - value) dan client akan menyimpan cookie tersebut di web browser
- Request selanjutnya, client akan selalu membawa cookie tersebut secara otomatis
- Dan server secara otomatis akan selalu menerima data cookie yang dibawa oleh client setiap client mengirimkan request

Membuat Cookie
- Cookie merupakan data yang dibuat di server dan sengaja agar disimpan di web browser
- Untuk membuat cookie di server, kita bisa menggunakan function http.SetCookie()
- Ketika membuat cookie, harus menentukan nama cookie, value cookie, dan cookie akan nempel di URl apa
- Cookie boleh lebih dari 1
*/

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprintf(writer, "Sukses create cookie")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("X-PZN-Name")
	if err != nil {
		fmt.Fprintf(writer, "No Cookie")
	} else {
		name := cookie.Value
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=Zakaria", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s:%s", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = "Zakaria"
	request.AddCookie(cookie)
	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
