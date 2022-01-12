package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

/**
Handler
- Server akan bertugas sebagai Web Server, sedangkan untuk menerima HTTP Request yang masuk ke server, kita butuh yang namanya Handler
- Handler di Go-lang di representasikan dalam intreface, dimana dalam kontraknya terdapat sebuah function bernama ServerHTTP()
yang digunakan sebagai function yang akan dieksekusi ketika menerima HTTP Request

HandlerFunc
- Salah satu implementasi dari interface Handler adalah HandlerFunc. tidak perlu membaut manual
- Kita bisa menggunakan HandlerFunc untuk membuat function handler HTTP
*/

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic webnya
		fmt.Fprint(writer, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/**
ServeMux
- Saat kita membaut web, kita biasanya ingin membuat bnyak sekali endpoint URL (tidak cuma nama domainya saja, ingin endpoint login, register)
- HandlerFunc sayangnya tidak mendukung itu (cuma 1 domain saja)
- Alternative implementasi dari Handler adalah ServeMux
- ServeMux dalah implementasi Handler yang bisa mendukung multiple endpoint. sebenarnya ServeMux adalah handler yang digabungkan, kita bisa menentukan handler a untuk endpoitn a
- ServeMux anggap saja sebagai router di berbagai bahasa pemograman
*/

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	})

	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hi")
	})

	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/**
URL Pattern
- URL Pattern dalam ServeMux sederhana, kita tinggal menabahkan string yang ingin kita gunakan sebagai endpoint, tanpa perlu memasukkan domain web kita (localhost, xxx.com dll)
- Jika URL Pattern dalam ServeMux kita tambahkan di akhirnya dengan garis miring, artinya semua url tersebut akan menerima path dengan awalan tersebut.
- Misal /images/ artinya akan dieksekusi jika endpoint nya /images/, /images/jpg, /images/contoh/lagi
- Namun jika terdapat URL Pattern yang lebih panjang, maka akan di perioritaskan adalah yang lebih panjang, misal jika kita terdapat URL /images/ dan /images/thumbnails/, maka jika
kita mengakses /images/thumbnails/ akan dituju ke URL /images/thumbnails/ bukan ke URL /images/
*/

func TestUrlPattern(t *testing.T) {
	mux := http.NewServeMux()

	// jika kita mengakses url unik dibelakang path /images/ akan tetap mengarah ke handler ini
	// contohnya /images/xxx atau /images/xxx/xxx
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Images")
	})

	// akan tetapi jika terdapat kesamaan url maka yang diperioritaskan yang lebih panjang
	// jika mengakses /images/thumbnails/ kan mengarah ke handler ini bukan ke /images/
	mux.HandleFunc("/images/thumbnails/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Thumbnails")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/**
Request
- Request adalah struct yang merepresentasikan HTTP Request yang dikirim oleh browser
- Semua informasi request yang dikirim bisa kita dapatkan di Request
- Contohnya seperti request url, method, http header, http body dan lain lain
*/

func TestRequest(t *testing.T) {
	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.Header)
		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.RequestURI)
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
