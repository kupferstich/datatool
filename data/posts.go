package data

import (
	"log"
	"os"
	"path/filepath"
)

type Posts struct {
	RootPath string `json:"-"`
	Posts    []Post `json:"posts"`
}

func NewPosts(rp string) *Posts {
	return &Posts{RootPath: rp}
}

func (p *Posts) Load() {
	filepath.Walk(p.RootPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}
		var post Post
		post.ID = filepath.Base(info.Name())
		err = LoadType(&post, p.RootPath)
		// Skip entry if data can not be loaded
		if err == ErrFileNotFound {
			return nil
		}
		if err != nil {
			log.Println(err)
			//return err
		}
		p.Posts = append(p.Posts, post)
		return nil
	})
}
