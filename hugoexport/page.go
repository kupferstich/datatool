package hugoexport

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/kupferstich/datatool/data"
)

// PageFrontMatter defines the front matter of the hugo page.
type PageFrontMatter struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Tags        []string  `json:"tags"`
	PublishDate time.Time `json:"publishdate"`
	Draft       bool      `json:"draft"`
	ID          string    `json:"id"`
	ImageFull   string    `json:"imagefull"`
	ImageCard   string    `json:"imagecard"`
	ImageThumb  string    `json:"imagethumb"`
}

// NewPageFrontMatterFromPicture maps the structure to the Picture
// type.
func NewPageFrontMatterFromPicture(p *data.Picture) *PageFrontMatter {
	var pfm PageFrontMatter
	pfm.ID = p.ID
	pfm.Title = p.Title
	pfm.Description = p.Topic
	pfm.Tags = p.Tags
	pfm.Draft = false
	pfm.ImageFull = fmt.Sprintf("img/artwork/%s/%s_big.jpg", p.ID, p.ID)
	pfm.ImageCard = fmt.Sprintf("img/artwork/%s/%s_small.jpg", p.ID, p.ID)
	pfm.ImageThumb = fmt.Sprintf("img/artwork/%s/%s_thumb.jpg", p.ID, p.ID)
	return &pfm
}

// ContentFromPicture creates a content page from a picture.
func ContentFromPicture(p *data.Picture, w io.Writer) {
	pfm := NewPageFrontMatterFromPicture(p)
	WritePage(pfm, p.Text, w)
}

// NewPageFrontMatterFromPerson maps the person structur to the hugo content.
func NewPageFrontMatterFromPerson(p *data.Person) *PageFrontMatter {
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
	pfm.Draft = false
	if len(p.ProfilePics) > 0 {
		var ppic string
		for p := range p.ProfilePics {
			ppic = p
		}
		pfm.ImageThumb = fmt.Sprintf("img/artist/%s/%s", p.GetID(), ppic)
	}
	//pfm.ImageFull = fmt.Sprintf("img/artwork/%s/%s_big.jpg", p.ID, p.ID)
	//pfm.ImageCard = fmt.Sprintf("img/artwork/%s/%s_small.jpg", p.ID, p.ID)
	//pfm.ImageThumb = fmt.Sprintf("img/artwork/%s/%s_thumb.jpg", p.ID, p.ID)
	return &pfm
}

// ContentFromPerson creates a content page from a person.
func ContentFromPerson(p *data.Person, w io.Writer) {
	pfm := NewPageFrontMatterFromPerson(p)
	WritePage(pfm, p.Text, w)
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
