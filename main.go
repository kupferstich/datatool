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

var sourceData data.Lister

func init() {
	flag.Parse()
	loadConf()
	sourceData = stabi.NewData(Conf.SourceFolder)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/list/", ListHandler).Methods("GET")
	router.HandleFunc("/form/{id}", FormHandler).Methods("GET")
	router.HandleFunc("/pic/{id}", PicHandler).Methods("GET")
	router.PathPrefix(`/files/`).
		Handler(http.StripPrefix("/files/", http.FileServer(http.Dir(Conf.FilesFolder))))

	log.Println("Starting server", Conf.ServerPort)
	err := http.ListenAndServe(Conf.ServerPort, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
