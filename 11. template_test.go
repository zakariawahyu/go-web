package go_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/**
Web Dinamis
- Sampai saat ini kita hanya membahas tentang membuat response menggunakan String dan juga static file
- Pada kenyataannya, saat kita membuat web, kita pasti akan membuat halaman yang dinamis, bisa berubah-ubah sesuai dengan data ang diakses oleh user
- Di Go-Lang terdapat fitur HTML template, yaitu fitur template yang bisa kita gunakan untuk membuat HTML yang dimanis

HTML Template
- Fitur HTML Template terdapat di package html/template
- Sebelum menggunakan HTML template, kita perlu terlebih dahulu membuat templatenya
- Template bisa berupa file atau string
- Bagian dinamis pada HTML Template adalah bagian yang menggunakan tanda {{ }}

Membaut Template
- Saat membuat template dengan string, kita perlu memberi tahu nama templatenya
- Dan untuk membuat text template, cukup buat text html dan untuk konten yang dinamis kita bisa menggunakan tanda {{.}}
- contoh : <h1>Nama saya {{.}}</h1>
*/

func SimpleHTML(writer http.ResponseWriter, request *http.Request) {
	//cara manual
	templateText := `<html><body>{{.}}</body></html>`
	//t, err := template.New("SIMPLE").Parse(templateText)
	//if err != nil {
	//	panic(err)
	//}

	//cara lebih mudah
	t := template.Must(template.New("SIMPLE").Parse(templateText))

	t.ExecuteTemplate(writer, "SIMPLE", "Hello World")
}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/**
Template Dari File
- Selain membuat template dari string, kita juga bisa membuat template lansung dari file
- Hal ini mempermudah kita, karena bisa langsung membuat file html
- Saat membuat template menggunakan file, secara otomatis nama file akan menjadi nama templatenya. misal kita punya file simple.html, maka nama templatenya adalah simple.html
*/

func SimpleHTMLFile(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello World Test Template File")
}

func TestSimpleHTMLFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/**
Template Directory
- Kadang biasanya kita jarang sekali menyebutkan file template satu persatu, bayangkan jika terdapat ratusan bahkan hampir ribuan file
- Alangkan baiknya untuk template kita simpan di satu directory
- Go-Lang template mendukung proses load template dari directory
- Hal ini memudahkan kita, sehingga tidak perlu menyebutkan nama filenya satu per satu2
*/

func TemplateDirectory(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello World Test Template Directory")
}

func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/**
Template dari Go-lang Embed
- Sejak Go-Lang 1.16 karena sudah ada Go-Lang Embed, jadi direkomendasikan menggunakan Go-lang embed untuk menyimpan data template
- Menggunakan Go-Lang embed menjadi kita tidak perlu ikut meng-copy template file lagi, karena sudah otomatis di embed di dalam distribution file
*/

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello World Test Template Embed")
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/**
Template Data
- Saat kita membaut template, kadang kita ingin menambahkan banyak data dinamis
- Hal ini bisa kita lakukan dengancara menggunakan data struct atau map
- Namun perlu dilakukan perubahan di dalam text templatenya, kita perlu memberi tahu Field atau Key mana yang akan kita gunakan
untuk mengisi data dinamis di template
- Kita bisa menyebutkan dengan cara seperti ini {{.NamaField}}
*/

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name": map[string]interface{}{
			"Name":     "Zakaria",
			"FullName": "Zakaria Wahyu",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Name struct {
	Name     string
	FullName string
}
type Page struct {
	Title string
	Name  Name
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name: Name{
			Name:     "Zakaria",
			FullName: "Zakaria Wahyu Nur Utomo",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/**
Template Action
- Go-Lang template mendukung perintah action, seperti percabangan, perulangan dan lain-lain

If Else
- {{if.Value}}T1{{end}}, Jika value tidak kosong, maka T1 akan dieksekusi, jika kosong tidak ada yang di eksekusi
- {{if.Value}}T1{{else}}T2{{end}}, jika value tidak kosong maka T1 akan dieksekusi dan jika kosong maka T2 yang akan dieksekusi
- {{if.Value1}}T1{{else if.Value2}}T2{{else}}T3{{end}}, jika value1 tidak kosong maka T1 akan dieksekusi, jika Value2 tidak kosong, maka T2 akan dieksekusi
dan jika tidak semuanya maka T3 akan dieksekusi
*/

func TemplateActionIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(writer, "if.gohtml", map[string]interface{}{
		"Title": "Template Action If",
		//"Name":  "Zakaria",
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/**
Operator Perbandinigan
- Go-Lang template juga mendukung operator perbandingan, ini cocok ketika butuh melakukan perbandingan number di if statement, berikut adalah operatornya
-	eq (equal)				=> artinya arg1 == arg2
-	ne (not equal)			=> artinya arg1 != arg2
- 	lt (less than)			=> artinya arg1 < arg2
- 	le (less than equal)	=> artinya arg1 <= arg2
-	gt (greater than)		=> artinya arg1 > arg2
- 	ge (greater than equal)	=> artinya arg1 >= arg2

Kenapa Operatornya di Depan?
- Hal ini dikarenakan, sebenarnya operator perbandingan tersebut adalah sebuah function
- Jadi saat kita menggunakan {{eq First Secod}}, sebenarnya dia akan memanggil sebuah function eq dengan parameter First dan Second : eq(First, Second)
*/

func TemplateActionOperator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(writer, "comparator.gohtml", map[string]interface{}{
		"Title":      "Template Action Operator",
		"FinalValue": 70,
	})
}

func TestTemplateActionOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/**
Range
- Range digunakan untuk melakukan iterasi data template
- Tidak ada perulanagan biasa seperti menggunakan for di Go-Lang template
- Yang kita bisa lakukan adalah menggunakan range untuk mengiterasi tiap data array, slice, map datau channel
- {{range $index, $element := .Value}} T1 {{end}}, jika value memiliki data, maka T1 akan dieksekusi sebanyak element value
dan kita bisa menggunakan $index untuk mengakses index/key dan $element untuk mengakses element/value
- {{range $index, $element := .Value}} T1 {{else}} T2 {{end}}, sama seperti sebelumnya, namun jika value tidak memiliki element apapun, maka T2 akan dieksekusi
*/

func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Title": "Template Action Range",
		"Hobbies": []string{
			"Game", "Read", "Code",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/**
With
- Kadang kita sering membuat nested struct
- Jika menggunakan template, kita bisa mengaksesnya menggunakan .Value.NestedValue
- Di template terdapat action with, yang bisa digunakan mengubah scope dot menjadi object yang kita mau
- {{with .Value}} <h1>{{.NestedValue}}</h1>{{end}}, jadi tidak pelu menuliskan .Value.NestedValue lagi
- {{with .Value}} T1 {{end}}, jika value tidak kosong, di T1 semua dot akan merefer ke value
- {{with .Value}} T1 {{else}} T2 {{end}}, sama seperti sebelumnya, namun jika value T1 kosong maka T2 akan dieksekusi
*/

func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))
	t.ExecuteTemplate(writer, "with.gohtml", map[string]interface{}{
		"Title": "Template Action With",
		"Name":  "Zakaria Wahyu",
		"Address": map[string]interface{}{
			"Street": "Mangga Besar",
			"City":   "Jakarta Barat",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/**
Comment
- Template juga mendukung komentar
- Komentar secara otomatis akan hilang ketika template text di parsing
- Untuk membuat komentar sangat sederhana, kita bisa menggunakan {{/* Contoh Komentar * /}}
*/
