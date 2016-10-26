package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kupferstich/datatool/data"
	"github.com/kupferstich/datatool/data/pdb"
	"github.com/kupferstich/datatool/hugoexport"
	"github.com/kupferstich/datatool/stabi"
)

// ConfFile is the path to the configuration file
var ConfFile = flag.String("conf", "conf.yaml", "Path to the conf (yaml) file")

// Init is used for initial works.
// "GetSource" generates all data into the data folder from the xml data
// Pictures will be transformed to jpg and copied.
var Init = flag.String("init", "", `Initial actions
		CreateList	creates a list from the xml data
		ImportTiff	imports the tiff pics as jpg`)

var Export = flag.Bool("export", false, "Export the data to hugo")

var sourceData data.Lister

var personDB data.PersonDBer

//var collection data.Lister

func init() {
	flag.Parse()
	loadConf()
	var err error
	personDB, err = pdb.Load(Conf.DataFolderPersons, Conf.DataFolderPictures)
	if err != nil {
		log.Println(err)
	}
	//sourceData = stabi.NewData(Conf.SourceFolder, personDB)
	//collection = stabi.NewData(Conf.DataFolder)
}

func main() {
	if *Init != "" {
		d := stabi.NewData(Conf.SourceFolder, personDB)
		_, err := d.List()
		if err != nil {
			log.Println(err)
		}
		if *Init == "CreateList" {
			err = d.Save(Conf.DataFolderPictures)
			if err != nil {
				log.Println(err)
			}
			log.Println("List is created...")
		} else if *Init == "ImportTiff" {
			err = d.SaveTiffAsJpg(Conf.DataFolderPictures)
			if err != nil {
				log.Println(err)
			}
		}
		return
	}
	if *Export == true {
		hugoexport.ImgArtwork(Conf.DataFolderPictures, Conf.HugoFolder)
		hugoexport.Artists(Conf.DataFolderPersons, Conf.DataFolderPictures, Conf.HugoFolder)
		hugoexport.Posts(Conf.DataFolderPosts, Conf.HugoFolder)
		return
	}
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	// Html page for all pictures and persons.
	// For type: persons or pictures is allowed. The logic is handled just
	// by the gui.
	router.HandleFunc("/list/{type}", ListHandler).Methods("GET")
	// Html page for editing a picture
	router.HandleFunc("/form/{id}", FormHandler).Methods("GET")
	// Html page for all persons
	router.HandleFunc("/edit/person/{id}", EditPersonHandler).Methods("GET")
	router.HandleFunc("/edit/post/{id}", EditPostHandler).Methods("GET")

	router.HandleFunc("/pic/all", PicAllHandler).Methods("GET")
	router.HandleFunc("/pic/{id}", PicHandler).Methods("GET")
	router.HandleFunc("/pic/{id}", PicSaveHandler).Methods("POST")

	router.HandleFunc("/person/all", PersonAllHandler).Methods("GET")
	router.HandleFunc("/person/{id}", PersonHandler).Methods("GET")
	router.HandleFunc("/person/{id}", PersonSaveHandler).Methods("POST")

	router.HandleFunc("/post/all", PostAllHandler).Methods("GET")
	router.HandleFunc("/post/{id}", PostHandler).Methods("GET")
	router.HandleFunc("/post/{id}", PostSaveHandler).Methods("POST")

	router.HandleFunc("/img/{id}-{maxWidth}-{maxHeight}", ImgHandler).Methods("GET")
	router.HandleFunc("/img/person/{id}/{size}/{file}", PersonImgHandler).Methods("GET")
	router.HandleFunc("/edit/post/{id}/{file}", PostImgHandler).Methods("GET")
	router.HandleFunc("/action/export", ExportHandler).Methods("GET")
	router.PathPrefix(`/files/`).
		Handler(http.StripPrefix("/files/", http.FileServer(http.Dir(Conf.FilesFolder))))

	log.Println("Starting server", Conf.ServerPort)
	err := http.ListenAndServe(Conf.ServerPort, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
