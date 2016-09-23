package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var homeTemplate = template.Must(template.ParseFiles("./static/index.html"))

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
	fmt.Fprintf(w, "%#v\n%v", sourceData, Conf)
	sourceData.List()
}
