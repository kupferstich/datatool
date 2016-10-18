package hugoexport

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/kupferstich/datatool/data"
)

// PageFrontMatter defines the front matter of the hugo page.
type PageFrontMatter struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Date        string   `json:"date"`
	Tags        []string `json:"tags"`
	PublishDate string   `json:"publishdate"`
	Draft       bool     `json:"draft"`
	ID          string   `json:"id"`
	ImageFull   string   `json:"imagefull"`
	ImageCard   string   `json:"imagecard"`
	ImageThumb  string   `json:"imagethumb"`
}

// NewPageFrontMatterFromPicture maps the structure to the Picture
// type.
func NewPageFrontMatterFromPicture(p data.Picture) *PageFrontMatter {
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
func ContentFromPicture(p data.Picture, w io.Writer) {
	pfm := NewPageFrontMatterFromPicture(p)
	head, err := json.MarshalIndent(pfm, "", "  ")
	if err != nil {
		log.Println(err)
	}
	w.Write(head)
	w.Write([]byte("\n\n\n"))
	w.Write([]byte(p.Text))
}
