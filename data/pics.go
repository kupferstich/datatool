package data

import (
	"log"
	"os"
	"path/filepath"
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
