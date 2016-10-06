package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kupferstich/datatool/data"
	"github.com/kupferstich/datatool/stabi"
	"github.com/nfnt/resize"
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
	/*vars := mux.Vars(r)
	type := vars["type"]*/
	staticFile(w, "./static/tmpl_list.html")
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	staticFile(w, "./static/tmpl_form.html")
}

func EditPersonHandler(w http.ResponseWriter, r *http.Request) {
	staticFile(w, "./static/tmpl_edit_person.html")
}

// PicAllHandler is for listing all availiable pictures
func PicAllHandler(w http.ResponseWriter, r *http.Request) {
	collection := stabi.NewData(Conf.DataFolderPictures, personDB)
	collection.LoadPictures()
	b, err := json.Marshal(collection.Pictures)
	if err != nil {
		log.Println(err)
	}
	w.Write(b)
}

// PicHandler sends the data of a picture in JSON format
func PicHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)
	var pic data.Picture
	pic.ID = id
	err := data.LoadType(&pic, Conf.DataFolderPictures)

	if err != nil {
		log.Println(err)
	}
	b, err := json.Marshal(pic)
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
	err = data.SaveType(&pic, Conf.DataFolderPictures)
	if err != nil {
		log.Println(err)
	}
}

func PersonAllHandler(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(personDB.GetAll())
	if err != nil {
		log.Println(err)
	}
	w.Write(b)
}

func PersonHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}
	p, _ := personDB.GetPerson(id)
	b, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
	}
	w.Write(b)
}

// PersonSaveHandler sends the data of a picture in JSON format
func PersonSaveHandler(w http.ResponseWriter, r *http.Request) {
	rbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	var person data.Person
	err = json.Unmarshal(rbody, &person)
	if err != nil {
		log.Println(err)
	}
	err = personDB.SavePerson(&person)
	if err != nil {
		log.Println(err)
	}
}

func PersonImgHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["file"]
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}
	maxWidth := 290
	maxHeight := 400

	p, _ := personDB.GetPerson(id)

	picturePath := filepath.Join(
		filepath.Dir(data.MakePath(p, Conf.DataFolderPersons)),
		filename,
	)

	file, err := os.Open(picturePath)
	defer file.Close()
	if err != nil {
		log.Println(err)
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Println(err)
	}
	t := resize.Thumbnail(uint(maxWidth), uint(maxHeight), img, resize.NearestNeighbor)
	jpeg.Encode(w, t, nil)

}
func ImgHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	maxWidth, err := strconv.Atoi(vars["maxWidth"])
	if err != nil {
		log.Println(err)
	}
	maxHeight, err := strconv.Atoi(vars["maxHeight"])
	if err != nil {
		log.Println(err)
	}
	picturePath := filepath.Join(
		Conf.DataFolderPictures,
		id,
		"00000001.jpg",
	)
	file, err := os.Open(picturePath)
	defer file.Close()
	if err != nil {
		log.Println(err)
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Println(err)
	}
	t := resize.Thumbnail(uint(maxWidth), uint(maxHeight), img, resize.NearestNeighbor)
	jpeg.Encode(w, t, nil)

}

func staticFile(w http.ResponseWriter, filename string) {
	sf, _ := ioutil.ReadFile("./static/tmpl_header.html")
	w.Write(sf)
	sf, _ = ioutil.ReadFile(filename)
	w.Write(sf)
	sf, _ = ioutil.ReadFile("./static/tmpl_footer.html")
	w.Write(sf)
}
