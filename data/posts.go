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

func GetPostPics(post *Post, pRootPath string) error {
	picExt := []string{".jpg", ".jpeg"}
	err := LoadType(post, pRootPath)
	if err != nil {
		return err
	}
	pics, _ := GetFiles(filepath.Dir(MakePath(post, pRootPath)), picExt)
	// Store the db values in var, because the pic source could be renamed or
	// deleted, then the old value should be not loaded.

	dbPostPics := post.PostPics
	// Delete the old values
	post.PostPics = make(map[string]Source)
	for _, postPic := range pics {
		dbSource, pok := dbPostPics[postPic]
		// If file is already in the db set the old value
		if pok {
			post.PostPics[postPic] = dbSource
		} else {
			post.PostPics[postPic] = Source{Value: postPic}
		}
	}
	return nil
}

// GetPostByID returns a post, when the id is provided
func (p *Posts) GetPostByID(id string) *Post {
	for _, post := range p.Posts {
		if post.ID == id {
			return &post
		}
	}
	return nil
}
