package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	. "shamir-secret-sharing/sss"
	"strconv"
)

// var (
// 	page = `
// Keys:
// {{.Xi}}:
// `
// )

// hello world, the web server
func newKeys(w http.ResponseWriter, req *http.Request) {
	p, _ := loadPage("newkeys")

	templ, err := template.ParseFiles("ui/newkeys.html")
	checkError(err, "ParseFiles")

	templ.Execute(w, p)
}

func decodeKeys(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "goodbye, world!\n")
}

func keyGen(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	checkError(err, "ParseForm")

	secret := req.FormValue("secret")
	nString := req.FormValue("n")
	kString := req.FormValue("k")

	n, err := strconv.Atoi(nString)
	checkError(err, "Atoi for n")

	k, err := strconv.Atoi(kString)
	checkError(err, "Atoi for k")

	keys := MakeKeys(secret, k, n)

	for _, s := range keys {
		io.WriteString(w, s)
		io.WriteString(w, "\n")
	}
}

func main() {
	http.HandleFunc("/newkeys", newKeys)
	http.HandleFunc("/keygen", keyGen)
	http.HandleFunc("/decoder", decodeKeys)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	fmt.Println("Listening...")
	err := http.ListenAndServe(":8080", nil)
	checkError(err, "ListenAndServe")

	fmt.Println("Connction closed.")
}

func checkError(err error, message string) {
	if err != nil {
		log.Fatal(message, ": ", err)
	}
}
func loadPage(title string) (*Page, error) {
	filename := "ui/" + title + ".html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

type Page struct {
	Title string
	Body  []byte
}
