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
	Artists     []string  `json:"artists"`
	ImageBase   string    `json:"imagebase"` // Basepath the the Images
	ImageFull   string    `json:"imagefull"`
	ImageMedium string    `json:"imagemedium"`
	ImageCard   string    `json:"imagecard"`
	ImageThumb  string    `json:"imagethumb"`
}

// NewPageFrontMatterFromPicture maps the structure to the Picture
// type.
func NewPageFrontMatterFromPicture(p *data.Picture) *PageFrontMatter {
	var pfm PageFrontMatter
	pfm.ID = p.ID
	pfm.Title = p.Title
	pfm.Description = fmt.Sprintf("<b>Erstellt</b>: %s", p.YearIssued)
	//p.Topic
	pfm.Tags = p.Tags
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
	pfm.ImageBase = fmt.Sprintf("img/artist/%s/", p.GetID())
	if len(p.ProfilePics) > 0 {
		pfm.ImageFull = fmt.Sprintf("img/artist/%s/profilepic_01_full.jpg", p.GetID())
		pfm.ImageCard = fmt.Sprintf("img/artist/%s/profilepic_01_square.jpg", p.GetID())
		pfm.ImageMedium = fmt.Sprintf("img/artist/%s/profilepic_01_medium.jpg", p.GetID())
		pfm.ImageThumb = fmt.Sprintf("img/artist/%s/profilepic_01_thumb.jpg", p.GetID())
	}
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
