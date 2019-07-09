package netSource

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"testing"
)

var p = &a
var a = 9

const HOST = "0.0.0.0"
const PORT = "8000"

func Test6(tt *testing.T) {
	fmt.Println(p, `----------'`)

	createHttpServer()
	fmt.Println(p, `---777777777777777777-------'`)

}

func createHttpServer() {

	http.HandleFunc("/", Hello)
	http.ListenAndServe(HOST+":"+PORT, nil)

}

type M struct {
	name string
}

func (m *M) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	fmt.Fprintf(w, "----- hello ")

	//headers := r.Header
	//for k, v := range headers {
	//	fmt.Println(k, "k")
	//	fmt.Println(v, "v")
	//
	//}
}
func Hello(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("404.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}
