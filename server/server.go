package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var (
	const page = `
Keys:
{{.Xi}}:
`
)

// hello world, the web server
func NewKeys(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func DecodeKeys(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/newkeys", NewKeys)
	http.HandleFunc("/decoder", DecodeKeys)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
