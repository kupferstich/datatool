package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kupferstich/datatool/data"
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

var sourceData data.Lister

//var collection data.Lister

func init() {
	flag.Parse()
	loadConf()
	sourceData = stabi.NewData(Conf.SourceFolder)
	//collection = stabi.NewData(Conf.DataFolder)
}

func main() {
	if *Init != "" {
		d := stabi.NewData(Conf.SourceFolder)
		_, err := d.List()
		if err != nil {
			log.Println(err)
		}
		if *Init == "CreateList" {
			err = d.Save(Conf.DataFolder)
			if err != nil {
				log.Println(err)
			}
			log.Println("List is created...")
		} else if *Init == "ImportTiff" {
			err = d.SaveTiffAsJpg(Conf.DataFolder)
			if err != nil {
				log.Println(err)
			}
		}
		return
	}
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/list/", ListHandler).Methods("GET")
	router.HandleFunc("/form/{id}", FormHandler).Methods("GET")
	router.HandleFunc("/pic/{id}", PicHandler).Methods("GET")
	router.HandleFunc("/pic/{id}", PicSaveHandler).Methods("POST")
	router.PathPrefix(`/files/`).
		Handler(http.StripPrefix("/files/", http.FileServer(http.Dir(Conf.FilesFolder))))

	log.Println("Starting server", Conf.ServerPort)
	err := http.ListenAndServe(Conf.ServerPort, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
