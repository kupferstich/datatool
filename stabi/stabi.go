//Package stabi contains all the implementations for the pictures of
//Staats- und Universitaets Bibliothek Hamburg
//It links the XML representation to the data structures.
package stabi

import (
	"encoding/xml"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/kupferstich/datatool/data"
	"github.com/kupferstich/datatool/mods"
)

type Data struct {
	Folder   string
	Pictures []data.Picture
}

//NewData creates a pointer to a Data element with a given path
func NewData(p string) *Data {
	return &Data{Folder: p}
}

//List creates a list of pictures with the original data given by the stabi
func (d *Data) List() (*[]data.Picture, error) {
	d.Pictures = nil
	filepath.Walk(d.Folder, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if !strings.EqualFold(filepath.Ext(path), ".xml") {
			return nil
		}
		pic, err := GetPicture(path)
		if err != nil {
			log.Println(err)
			return err
		}
		d.Pictures = append(d.Pictures, *pic)
		return nil
	})
	return &d.Pictures, nil
}

//Save stores the list into the root folder
func (d *Data) Save(root string) error {
	for _, p := range d.Pictures {
		err := data.SaveType(&p, root)
		if err != nil {
			return err
		}
	}
	return nil
}

//SaveTiffAsJpg takes the tiff pictures and saves them inside the data folder as
//jpg
func (d *Data) SaveTiffAsJpg(root string) error {
	for _, p := range d.Pictures {
		dst := data.MakePath(&p, root)
		dst = filepath.Join(
			filepath.Dir(dst),
			"00000001.jpg",
		)
		log.Println("Create:", dst)
		cmd := exec.Command("magick", p.File, dst)
		err := cmd.Run()
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

// GetPicture maps the data of the mods xml to the data.Picture and returns a
// pointer to the created pic
func GetPicture(fpath string) (*data.Picture, error) {
	file, err := os.Open(fpath)
	defer file.Close()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var mets mods.Mets
	if err := xml.NewDecoder(file).Decode(&mets); err != nil {
		return nil, err
	}
	pic := NewDataPicture(&mets)
	fp := filepath.Dir(fpath)
	pic.File = filepath.Join(fp, "00000001.tif")
	return pic, nil
}
