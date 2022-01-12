package go_web

import (
	"net/http"
	"testing"
)

/**
Go-Lang Web
- Go-lang saat ini populer dijadikan salah satu pilihan bahasa pemograman untuk membuat Web, terutama WEB API (backend)
- Selain itu, di Go-Lang juga sudah disediakan package untuk membuat Web, bahkan disertakan pula package untuk implementasi unit testingnya
- Hal ini menjadikan pembuatan Web menggunakan Go-Lang lebih mudah, karena tidak butuh menggunakan library atau framework

Cara kerja Go-Lang Web
- Web Browser akan melakukan HTTP Request ke Web Server
- Golang menerima HTTP Request, lalu mengeksekusi request tersebut sesuai dengan yang diminta
- Setelah melakukan eksekusi request, Golang akan mengembalikan data dan di render sesuai dengan kebutuhan misal HTML, CSS, JavaScript dll
- Golang akan mengembalikan content hasil render tersebut sebagai HTTP Response ke Web Browser
- Web Browser menerima content dari Web Server, lalu me render content tersebut sesuai dengan tipe contentnya
*/

/**
Package net/http
- Pada beberapa bahasa pemograman lain, seperti Java misalnya. untuk membuat web biasanya dibutuhkan tambahan library atau framework
- Sedangkan di Go-Lang sudah disediakan pacakage untuk membuat web bernama package net/http
- Sehingga untuk membuat web menggunakan Go-Lang, kita tidak butuh lagi library tambahan, kita bisa menggunakan package yang sudah tersedia
- Walaupun memang saat kita akan membuat web dalah skala besar, direkomendasikan menggunakan framework karena beberapa hal sudah dipermudah oleh web framework
- Namun pada course ini, kita akan fokus menggunakan package net/http untuk membuat web nya, karena semua framework Go-Lang akan menggunakan net/http sebagai basis dasar frameworknya

Server
- Server adalah struct yang terdapat di package net/http yang digunakan sebagai represetasi Web Server di Go-Lang
- Untuk membuat web, kita wajib membuat Server
- Saat membuat data Server, ada beberapa hal yang perlu kita tentukan, seperti host dan juga port tempat dimana Web kita berjalan
- Rekomendasi untuk belajar jangan menggunakan port 80, karena port 80 di beberapa sistem operasi diwajibkan administrator
- Contoh menggunakan port 4 digit 8080, 8181, 8283 untuk belajar
- Setelah membuat Server, kita bisa menjalankan Server tersebut menggunakan function ListenAndServe()
*/

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
