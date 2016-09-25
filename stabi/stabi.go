//Package stabi contains all the implementations for the pictures of
//Staats- und Universitaets Bibliothek Hamburg
package stabi

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/kupferstich/datatool/data"
	"github.com/kupferstich/datatool/mods"
)

type Data struct {
	Folder   string
	Pictures []data.Picture
}

func NewData(p string) *Data {
	fmt.Println(p)
	return &Data{Folder: p}
}

func (d *Data) List() (*[]data.Picture, error) {
	d.Pictures = nil
	filepath.Walk(d.Folder, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if !strings.EqualFold(filepath.Ext(path), ".xml") {
			return nil
		}
		file, err := os.Open(path)
		defer file.Close()
		if err != nil {
			log.Println(err)
			return err
		}
		var mets mods.Mets
		if err := xml.NewDecoder(file).Decode(&mets); err != nil {
			return err
		}
		pic := NewDataPicture(&mets)
		d.Pictures = append(d.Pictures, *pic)
		return nil
	})
	return &d.Pictures, nil
}
