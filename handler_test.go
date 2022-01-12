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
