package hugoexport

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/kupferstich/datatool/data"
)

func Posts(postRootPath, exportRootPath string) {
	posts := data.NewPosts(postRootPath)
	posts.Load()
	for _, p := range posts.Posts {
		ExportPostPics(&p, postRootPath, exportRootPath)
		ExportPostContent(&p, exportRootPath)
		ExportPostData(&p, postRootPath, exportRootPath)
	}
}

func ExportPostContent(p *data.Post, exportRootPath string) {
	dstPath := filepath.Join(
		exportRootPath,
		ContentPostSubfolder,
		fmt.Sprintf("%s.md", p.ID),
	)
	os.MkdirAll(filepath.Dir(dstPath), 0777)
	f, err := os.Create(dstPath)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	ContentFromPost(p, f)
}

func ExportPostPics(p *data.Post, postRootPath, exportRootPath string) {
	postPath := filepath.Dir(data.MakePath(p, postRootPath))
	for pic := range p.PostPics {
		img := openPic(filepath.Join(postPath, pic))
		dst := filepath.Join(
			exportRootPath,
			ImgPostSubfolder,
			p.ID,
			pic,
		)
		if skipImageCreation(dst) {
			continue
		}
		resizePic(img, PostPicSize, dst, PostPicResizeType)
	}
}

func ExportPostData(p *data.Post, postRootPath, exportRootPath string) {
	dataPath := filepath.Join(
		exportRootPath,
		JSONPostSubfolder,
		fmt.Sprintf("%s.json", p.ID),
	)
	os.MkdirAll(filepath.Dir(dataPath), 0777)
	b, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile(dataPath, b, 0777)
	if err != nil {
		log.Println(err)
	}
}
