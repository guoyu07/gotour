// httpserver
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Hello struct {
}

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.RequestURI)
	if r.RequestURI == "/stop" {
		log.Println("Http Server stoped.")
		os.Exit(0)
	}
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func main() {
	var h Hello
	http.Handle("/string", String("Hello String"))
	http.Handle("/struct", &Struct{"Good", "job", "John"})
	http.Handle("/stop", &h)

	err := http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Http Server start up on port 4000.")
}
