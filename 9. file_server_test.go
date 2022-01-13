package go_web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

/**
FileServer
- Go-Lang memiliki sebuah fitur yang bernama FileServer
- Dengan ini, kita bisa membuat Handler di Go-Lang web yang digunakan sebagai static file server
- Dengan menggunakan FileServer, kita tidak perlu manual me-load file lagi
- FileServer adalah Handler, jadi bisa kita tambahkan ke dalam http.Server atau http.ServerMux

404 Not Found File Server
- Jika kita coba jalankan, saat kita membuka misal /static/index.html makan akan dapat error 404 Not FOund
- Hal ini dikarenakan FIleServer akan membaca URL, lalu mencari file berdasarkan url nya
- Jadi ketika kita membuat /static/index.js maka FileServer akan mencari ke dile /resource/static/index.js
- Hal ini menyebabkan 404 Not Found karena memang file tidak bis ditemukan
- Oleh karena itu kita bisa menggunakan function http.StripPrefix untuk menghapus prefix di url
*/

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	//mux.Handle("/static/", fileServer) // tidak menambahkan StripPrefix maka akan not found
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

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
Go-Lang Embed
- Go-Lang 1.16 terdapat fitur baru yang bernama Go-Lang embed
- Dalam Go-Lang embed kita bisa embed file ke dalam binary distribution file, hal ini mempermudah sehingga kita tidak perlu mengcopy static file lagi
- Go-Lang embed juga memiliki fitur yang bernama embed.FS, fitur ini bisa di intergrasikan dengan FileServer

404 Not Found Golang Embed
- Jika kita coba jalankan dan coba buka /static/index.js, maka kita akan mendapatkan error 404 Not Found
- Hal ini dikarenakan di Go-Lang embed, nama folder ikut menjadi nama resourcenya. misal url /static/index.js, jadi untuk mengaksesnya kita perlu menggunakan
url /static/resource/index.js
- Jika kita ingin langsung mengakses file index.js tanpa menggunakan resource, kita bisa menggunakan function fs.Sub untuk mendapatkan dub directory
*/

//go:embed resources
var resources embed.FS

func TestFileServerGolangEmbed(t *testing.T) {
	directory, err := fs.Sub(resources, "resources")
	if err != nil {
		panic(err)
	}
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
