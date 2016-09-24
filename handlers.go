package main

import (
	"html/template"
	"log"
	"net/http"
)

var homeTemplate = template.Must(template.ParseFiles(
	"./static/index.html",
	"./static/tmpl_header.html"))

func init() {

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	p := &struct {
		Title string
		Test  string
	}{
		"hi",
		"ho",
	}
	homeTemplate.Execute(w, p)
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"./static/tmpl_list.html",
		"./static/tmpl_header.html",
	))
	pics, _ := sourceData.List()
	tmpl.Execute(w, pics)

}
