package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kupferstich/datatool/data"
	"github.com/kupferstich/datatool/stabi"
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

//ListHandler is for listing all availiable pictures
func ListHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"./static/tmpl_list.html",
		"./static/tmpl_header.html",
	))
	pics, _ := sourceData.List()
	tmpl.Execute(w, pics)

}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)
	/*tmpl := template.Must(template.ParseFiles(
		"./static/tmpl_header.html",
	))
	tmpl.Execute(w, id)*/
	staticFile, _ := ioutil.ReadFile("./static/tmpl_form.html")
	w.Write(staticFile)

}

// PicHandler sends the data of a picture in JSON format
func PicHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)
	pic, err := stabi.LoadPicture(id, Conf.DataFolder)

	if err != nil {
		log.Println(err)
	}
	b, err := json.Marshal(*pic)
	if err != nil {
		log.Println(err)
	}
	w.Write(b)
}

// PicSaveHandler sends the data of a picture in JSON format
func PicSaveHandler(w http.ResponseWriter, r *http.Request) {
	rbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	var pic data.Picture
	err = json.Unmarshal(rbody, &pic)
	if err != nil {
		log.Println(err)
	}
	err = data.SaveType(&pic, Conf.DataFolder)
	if err != nil {
		log.Println(err)
	}
}
