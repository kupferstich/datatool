package main

import (
	"encoding/json"
	"fmt"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/gorilla/mux"
	"github.com/kupferstich/datatool/data"
	"github.com/kupferstich/datatool/hugoexport"
	"github.com/kupferstich/datatool/stabi"
	"github.com/nfnt/resize"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/list/pictures", http.StatusPermanentRedirect)
}

//ListHandler is for listing all availiable pictures
func ListHandler(w http.ResponseWriter, r *http.Request) {
	/*vars := mux.Vars(r)
	type := vars["type"]*/
	staticFile(w, "tmpl_list.html")
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	staticFile(w, "tmpl_form.html")
}

func EditPersonHandler(w http.ResponseWriter, r *http.Request) {
	staticFile(w, "tmpl_edit_person.html")
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	staticFile(w, "tmpl_edit_post.html")
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
	id := vars["id"]
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
	id := vars["id"]
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
		w.WriteHeader(http.StatusNotFound)
		return
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Println(err)
	}
	t := resize.Thumbnail(uint(maxWidth), uint(maxHeight), img, resize.NearestNeighbor)
	jpeg.Encode(w, t, nil)

}

func PostImgHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["file"]
	id := vars["id"]
	maxWidth := 290
	maxHeight := 400
	var post data.Post
	post.ID = id
	log.Println(vars)

	data.LoadType(&post, Conf.DataFolderPosts)

	picturePath := filepath.Join(
		filepath.Dir(data.MakePath(&post, Conf.DataFolderPosts)),
		filename,
	)

	file, err := os.Open(picturePath)
	defer file.Close()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Println(err)
	}
	t := resize.Thumbnail(uint(maxWidth), uint(maxHeight), img, resize.Lanczos3)
	jpeg.Encode(w, t, nil)

}

// PicAllHandler is for listing all availiable pictures
func PostAllHandler(w http.ResponseWriter, r *http.Request) {
	posts := data.NewPosts(Conf.DataFolderPosts)
	posts.Load()

	b, err := json.Marshal(posts.Posts)
	if err != nil {
		log.Println(err)
	}
	w.Write(b)
}

// PicHandler sends the data of a picture in JSON format
func PostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	/*
		// The root of the post should handle also the pictures of the
		// post folder.
		ext := filepath.Ext(id)
		if strings.EqualFold(ext, "jpg") || strings.EqualFold(ext, "jpeg") {
			PostImgHandler(w, r)
			return
		}*/

	var post data.Post
	if id == "new" {
		posts := data.NewPosts(Conf.DataFolderPosts)
		posts.Load()
		post.ID = fmt.Sprintf("%d", len(posts.Posts))
	} else {
		post.ID = id
		err := data.LoadType(&post, Conf.DataFolderPosts)
		if err != nil {
			log.Println(err)
		}
		data.GetPostPics(&post, Conf.DataFolderPosts)
	}
	b, err := json.Marshal(post)
	if err != nil {
		log.Println(err)
	}
	w.Write(b)
}

// PicSaveHandler sends the data of a picture in JSON format
func PostSaveHandler(w http.ResponseWriter, r *http.Request) {
	rbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	var post data.Post
	err = json.Unmarshal(rbody, &post)
	if err != nil {
		log.Println(err)
	}
	err = data.SaveType(&post, Conf.DataFolderPosts)
	if err != nil {
		log.Println(err)
	}
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
	//t := resize.Thumbnail(uint(maxWidth), uint(maxHeight), img, resize.Lanczos3)
	t := imaging.Fit(img, maxWidth, maxHeight, imaging.Lanczos)
	jpeg.Encode(w, t, nil)

}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	target := vars["target"]

	fmt.Fprintln(w, id, target)
	log.Printf("%v", vars)
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var fp string
	switch target {
	case "post":
		fp = filepath.Join(
			Conf.DataFolderPosts,
			id,
			handler.Filename)
	case "person":
		p, _ := personDB.GetPerson(id)
		fp = filepath.Join(
			filepath.Dir(data.MakePath(p, Conf.DataFolderPersons)),
			handler.Filename,
		)

	}

	f, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

func ExportHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Export gestartet")
	hugoexport.ImgArtwork(Conf.DataFolderPictures, Conf.DataFolderPosts, Conf.HugoFolder)
	fmt.Fprintln(w, "Export erfolgt")
}

func staticFile(w http.ResponseWriter, filename string) {
	sf, _ := ioutil.ReadFile(
		filepath.Join(
			Conf.StaticFolder,
			"tmpl_header.html",
		))
	w.Write(sf)
	sf, _ = ioutil.ReadFile(
		filepath.Join(
			Conf.StaticFolder,
			filename,
		))
	w.Write(sf)
	sf, _ = ioutil.ReadFile(
		filepath.Join(
			Conf.StaticFolder,
			"footer.html",
		))
	w.Write(sf)
}
