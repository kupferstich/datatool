package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path"

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
	router.HandleFunc("/pic/{id}", PicSaveHandler).Methods("POST")
	router.PathPrefix(`/files/`).
		Handler(http.StripPrefix("/files/", http.FileServer(http.Dir(Conf.FilesFolder))))

	log.Println("Starting server", Conf.ServerPort)
	err := http.ListenAndServe(Conf.ServerPort, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//loadPicture gets the data of a picture with a given id.
//If there is no data availiable inside the DataFolder the
//Data is loaded from the sourceFolder.
//The logik for the filepath is:
//DataFolder/[id]/data.json
func loadPicture(id string) (*data.Picture, error) {
	var pic data.Picture
	pic.ID = id
	err := data.LoadType(&pic, Conf.DataFolder)
	if err == data.ErrFileNotFound {

		// If there is no data saved, the meta data is used
		spath := path.Join(
			Conf.SourceFolder,
			id,
			fmt.Sprintf("%s.xml", id),
		)
		fmt.Println(spath)
		pic, err := stabi.GetPicture(spath)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return pic, nil
	}
	/*file, err := os.Open(fpath)
	defer file.Close()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if err := xml.NewDecoder(file).Decode(&pic); err != nil {
		return nil, err
	}*/

	return &pic, err
}
