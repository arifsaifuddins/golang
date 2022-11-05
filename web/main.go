package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// data
type pageData struct {
	Title string
	Body  string
}

func home(w http.ResponseWriter, r *http.Request) {

	data := pageData{"GO | HTML Template", "HELLO GO TEMPLATE HTML"}
	t, _ := template.ParseFiles("templates/home.gohtml")
	t.Execute(w, data)
}

func about(w http.ResponseWriter, r *http.Request) {

	data := pageData{"GO About | HTML Template", "iam arief saifuddien"}
	t, _ := template.ParseFiles("templates/about.gohtml")
	t.Execute(w, data)
}

func main() {

	// static
	port := flag.String("p", "3000", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	http.Handle("/statics/", http.StripPrefix(strings.TrimRight("/statics/", "/"), http.FileServer(http.Dir(*directory))))

	// routes
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	// run serve
	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
