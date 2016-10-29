package hugoexport

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/kupferstich/datatool/data"
	"github.com/kupferstich/datatool/data/pdb"
)

// Artists exports all the artists into the Artist content folder
func Artists(artistRootFolder, pictureRootFolder, postsRootFolder, exportRootPath string) {
	artists, err := pdb.Load(artistRootFolder, "")
	posts := data.NewPosts(postsRootFolder)
	posts.Load()
	if err != nil {
		log.Println(err)
	}
	for _, artist := range artists.Persons {
		log.Println(artist.Identify())
		ExportArtistProfilePics(&artist, artistRootFolder, exportRootPath)
		ExportArtistContent(&artist, posts, exportRootPath)
		ExportArtistData(&artist, pictureRootFolder, exportRootPath)
	}
}

// ExportArtistContent exports the picture data into the content folder
func ExportArtistContent(p *data.Person, posts *data.Posts, exportRootPath string) {
	dstPath := filepath.Join(
		exportRootPath,
		ContentArtistSubfolder,
		fmt.Sprintf("%s.md", p.Identify()),
	)
	os.MkdirAll(filepath.Dir(dstPath), 0777)
	f, err := os.Create(dstPath)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	ContentFromPerson(p, posts, f)
}

// ExportArtistProfilePics exports the profile pics into the hugo image folder. Adds also
// the first image to the imagesource. Because of that the pic export should be called
// **before** the content export.
func ExportArtistProfilePics(p *data.Person, artistRootFolder, exportRootPath string) {
	artistPath := filepath.Dir(data.MakePath(p, artistRootFolder))
	counter := 0
	for pic := range p.ProfilePics {
		counter++
		// First picture is used es picture for the hugo post
		if counter == 1 {
			p.PostImage = pic
			err := data.SaveType(p, artistRootFolder)
			if err != nil {
				log.Println(err)
			}
		}
		img := openPic(filepath.Join(artistPath, pic))
		for key, size := range ResizeSizes {
			dst := filepath.Join(
				exportRootPath,
				ImgArtistSubfolder,
				p.GetID(),
				fmt.Sprintf("profilepic_%02d_%s.jpg", counter, key),
			)
			if skipImageCreation(dst) {
				continue
			}
			var rType = ResizeFit
			if key == "thumb" || key == "square" {
				rType = ResizeThumbnail
			}
			resizePic(img, size, dst, rType)
		}

	}
}

func ExportArtistData(p *data.Person, picRoot, exportRootPath string) {
	dataPath := filepath.Join(
		exportRootPath,
		JSONArtistSubfolder,
		fmt.Sprintf("%s.json", p.GetID()),
	)
	p.ExtID = p.Identify()
	p.Pictures = data.SortPictures(p.Pictures, picRoot)
	b, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Println(err)
	}
	os.MkdirAll(filepath.Dir(dataPath), 0777)
	err = ioutil.WriteFile(dataPath, b, 0777)
	if err != nil {
		log.Println(err)
	}
}
