package data

import (
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/kupferstich/datatool/helpers"
)

// LoadPictures loads the pictures from data folder. For loading the pictures
// the data.LoadType function is used.
func LoadPictures(root string) []Picture {
	var pics []Picture
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}
		var pic Picture
		pic.ID = filepath.Base(info.Name())
		err = LoadType(&pic, root)
		// Skip entry if data can not be loaded
		if err == ErrFileNotFound {
			return nil
		}
		if err != nil {
			log.Println(err)
			//return err
		}
		pics = append(pics, pic)
		return nil
	})
	return pics
}

// SortPictures by YearIssued takes a slice of stings with the ids and returns
// the values sorted by year.
func SortPictures(picIDs []string, root string) []string {
	pics := LoadPictures(root)
	sort.Sort(ByYearIssued(pics))
	var sorted []string
	for _, p := range pics {
		if helpers.StrInSlice(p.ID, picIDs) {
			sorted = append(sorted, p.ID)
		}
	}
	return sorted
}
