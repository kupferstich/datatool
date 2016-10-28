package hugoexport

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/kupferstich/datatool/data"
)

// PageFrontMatter defines the front matter of the hugo page.
type PageFrontMatter data.PageFrontMatter

// NewPageFrontMatterFromPicture maps the structure to the Picture
// type.
func NewPageFrontMatterFromPicture(p *data.Picture, posts *data.Posts) *PageFrontMatter {
	var pfm PageFrontMatter
	pfm.ID = p.ID
	pfm.Title = p.Title
	pfm.Description = fmt.Sprintf("**Erstellt**: %s", p.YearIssued)
	//p.Topic
	pfm.Tags = p.Tags
	pfm.Posts = posts.GetPostsForPicture(p.ID)
	pfm.Draft = true
	if p.Status == "fertig" {
		pfm.Draft = false
	}
	pfm.Artists = p.Persons
	pfm.ImageBase = fmt.Sprintf("img/artwork/%s/", p.ID)
	pfm.ImageFull = fmt.Sprintf("img/artwork/%s/%s_full.jpg", p.ID, p.ID)
	pfm.ImageMedium = fmt.Sprintf("img/artwork/%s/%s_medium.jpg", p.ID, p.ID)
	pfm.ImageCard = fmt.Sprintf("img/artwork/%s/%s_square.jpg", p.ID, p.ID)
	pfm.ImageThumb = fmt.Sprintf("img/artwork/%s/%s_thumb.jpg", p.ID, p.ID)
	pfm.Date = p.BlogDate
	pfm.PublishDate = p.PublishDate
	return &pfm
}

// ContentFromPicture creates a content page from a picture.
func ContentFromPicture(p *data.Picture, posts *data.Posts, w io.Writer) {
	pfm := NewPageFrontMatterFromPicture(p, posts)
	WritePage(pfm, p.Text, w)
}

// NewPageFrontMatterFromPerson maps the person structur to the hugo content.
func NewPageFrontMatterFromPerson(p *data.Person, posts *data.Posts) *PageFrontMatter {
	var pfm PageFrontMatter
	pfm.ID = p.GetID()
	pfm.Title = p.FullName
	pfm.Description = fmt.Sprintf(
		"* %d in %s; † %d in %s",
		p.YearBirth,
		p.CityBirth,
		p.YearDeath,
		p.CityDeath,
	)
	pfm.Links = p.Links
	pfm.Draft = true
	if p.Status == "fertig" {
		pfm.Draft = false
	}
	pfm.Posts = posts.GetPostsForPerson(p.GetID())
	pfm.Date = p.BlogDate
	pfm.PublishDate = p.PublishDate
	pfm.ImageBase = fmt.Sprintf("img/artist/%s/", p.GetID())
	if len(p.ProfilePics) > 0 {
		pfm.ImageFull = fmt.Sprintf("img/artist/%s/profilepic_01_full.jpg", p.GetID())
		pfm.ImageCard = fmt.Sprintf("img/artist/%s/profilepic_01_square.jpg", p.GetID())
		pfm.ImageMedium = fmt.Sprintf("img/artist/%s/profilepic_01_medium.jpg", p.GetID())
		pfm.ImageThumb = fmt.Sprintf("img/artist/%s/profilepic_01_thumb.jpg", p.GetID())
		pfm.ImageSource = p.ProfilePics[p.PostImage]
	}
	return &pfm
}

// ContentFromPerson creates a content page from a person.
func ContentFromPerson(p *data.Person, posts *data.Posts, w io.Writer) {
	pfm := NewPageFrontMatterFromPerson(p, posts)
	WritePage(pfm, p.Text, w)
}

// NewPageFrontMatterFromPost maps the person structur to the hugo content.
func NewPageFrontMatterFromPost(p *data.Post) *PageFrontMatter {
	pfm := PageFrontMatter(p.PageFrontMatter)
	if p.Image == "" {
		for k, pp := range p.PostPics {
			p.Image = fmt.Sprintf("img/post/%s", k)
			p.ImageSource = pp
			break
		}
	}
	pfm.Draft = true
	if p.Status == "fertig" {
		pfm.Draft = false
	}
	pfm.ImageCard = p.Image
	pfm.ImageMedium = p.Image
	pfm.ImageSource = p.ImageSource
	return &pfm
}

func ContentFromPost(p *data.Post, w io.Writer) {
	pfm := NewPageFrontMatterFromPost(p)
	WritePage(pfm, p.Content, w)
}

// WritePage writes the page into the io.Writer. First the FrontMatter and then the
// content.
func WritePage(pfm *PageFrontMatter, content string, w io.Writer) {
	head, err := json.MarshalIndent(pfm, "", "  ")
	if err != nil {
		log.Println(err)
	}
	w.Write(head)
	w.Write([]byte("\n\n\n"))
	w.Write([]byte(content))
}
