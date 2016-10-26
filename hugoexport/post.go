package hugoexport

import (
	"fmt"
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
