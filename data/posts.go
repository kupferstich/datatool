package data

import (
	"log"
	"os"
	"path/filepath"

	"github.com/kupferstich/datatool/helpers"
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

func (p *Posts) GetPostsForPicture(pictureID string) []string {
	var out []string
	for _, post := range p.Posts {
		if helpers.StrInSlice(pictureID, post.Pictures) {
			out = append(out, post.ID)
		}
	}
	return out
}

func (p *Posts) GetPostsForPerson(personID string) []string {
	var out []string
	for _, post := range p.Posts {
		if helpers.StrInSlice(personID, post.Artists) {
			out = append(out, post.ID)
		}
	}
	return out
}

func GetPostPics(post *Post, pRootPath string) error {
	picExt := []string{".jpg", ".jpeg"}
	err := LoadType(post, pRootPath)
	if err != nil {
		return err
	}
	pics, _ := GetFiles(filepath.Dir(MakePath(post, pRootPath)), picExt)

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
