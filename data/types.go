package data

import (
	"fmt"
	"image"
	"time"
)

// Picture stores all the entities for the histblogger
type Picture struct {
	ID           string    `xml:"id,attr" json:"ID"`
	SrcPath      string    `xml:"-" json:"-"`
	File         string    `xml:"-" json:"-"`
	Title        string    `xml:"title" json:"Title"`
	Topic        string    `xml:"topic" json:"Topic"`
	Text         string    `xml:"text" json:"Text"`
	CanvasWidth  int       `xml:"canvasWidth" json:"canvasWidth"`
	CanvasHeight int       `xml:"canvasHeight" json:"canvasHeight"`
	Areas        []Area    `xml:"areas" json:"Areas"`
	Captured     int       `xml:"captured" json:"Captured"` //Year, when picture was digitalized
	Place        string    `xml:"place" json:"Place"`       //Place where the picture was issued
	YearIssued   string    `xml:"yearIssued" json:"YearIssued"`
	Persons      []string  `xml:"persons" json:"Persons"`
	Tags         []string  `xml:"tags" json:"Tags"`
	Links        []Link    `xml:"links" json:"Links"`
	Status       string    `xml:"status" json:"Status"`
	BlogDate     time.Time `xml:"-" json:"BlogDate"`
	PublishDate  time.Time `xml:"-" json:"PublishDate"`
}

// Identify implements the Identifier interface for loading and saving
func (p *Picture) Identify() string {
	return p.ID
}

// TypeName implements the Identifier inteface for loading and saving
func (p *Picture) TypeName() string {
	return "picture"
}

// Person represents the entities of a person
type Person struct {
	ID          int               `xml:"id,attr" json:"intID"`
	GND         string            `xml:"gnd" json:"GND"`
	ExtID       string            `xml:"gnd" json:"personID"`
	Idents      []Ident           `xml:"idents" json:"idents"`
	MasterID    int               `xml:"master,attr" json:"masterID"`
	Type        string            `xml:"type,attr" json:"Type"`
	NameFamily  string            `xml:"name>family" json:"FamilyName"`
	NameGiven   string            `xml:"name>given" json:"GivenName"`
	FullName    string            `xml:"fullName" json:"FullName"`
	YearBirth   int               `xml:"yearBirth" json:"YearBirth"`
	YearDeath   int               `xml:"yearDeath" json:"YearDeath"`
	CityBirth   string            `xml:"cityBirth" json:"CityBirth"`
	CityDeath   string            `xml:"cityDeath" json:"CityDeath"`
	Text        string            `xml:"text" json:"Text"`
	ProfilePics map[string]Source `xml:"profilePics" json:"ProfilePics"`
	Pictures    []string          `xml:"pictures" json:"Pictures"`
	Links       []Link            `xml:"links" json:"Links"`
}

// GetID returns the ID of the person
func (p *Person) GetID() string {
	if p.GND != "" {
		return fmt.Sprintf("GND%s", p.GND)
	}
	return fmt.Sprintf("%04d", p.ID)
}

// Identify implements the Identifier interface for loading and saving
func (p *Person) Identify() string {
	return fmt.Sprintf("%s%s_%s", p.NameFamily, p.NameGiven, p.GetID())
}

// TypeName implements the Identifier inteface for loading and saving
func (p *Person) TypeName() string {
	return "personData"
}

// Area defines parts of the picture
type Area struct {
	ID      string   `xml:"id,attr" json:"areaID"`
	Rect    Fabric   `xml:"rect" json:"rect"`
	Shape   string   `xml:"shape" json:"Shape"`
	Coords  string   `xml:"coords" json:"Coords"`
	Persons []string `xml:"persons" json:"Persons"`
	Text    string   `xml:"text" json:"Text"`
	Links   []Link   `xml:"links" json:"Links"`
}

// ImageRect returns an image.Rectangle for the area. This method is
// used to cut the area out of the image. The input scale factor scales
// the size of the source pic to the canvas size, where the area was
// defined
func (a Area) ImageRect(scale float32) image.Rectangle {
	topLeft := image.Point{
		int(a.Rect.Left * scale),
		int(a.Rect.Top * scale),
	}
	botRight := image.Point{
		int((a.Rect.Left + (float32(a.Rect.Width) * a.Rect.ScaleX)) * scale),
		int((a.Rect.Top + (float32(a.Rect.Height) * a.Rect.ScaleY)) * scale),
	}
	return image.Rectangle{topLeft, botRight}
}

// Fabric defines rects of the fabricjs library
type Fabric struct {
	Type             string  `xml:"type" json:"type"`
	Left             float32 `xml:"left" json:"left"`
	Top              float32 `xml:"top" json:"top"`
	Width            int     `xml:"width" json:"width"`
	Height           int     `xml:"height" json:"height"`
	Fill             string  `xml:"fill" json:"fill"`
	Opacity          float32 `xml:"opacity" json:"opacity"`
	ScaleX           float32 `xml:"scaleX" json:"scaleX"`
	ScaleY           float32 `xml:"scaleY" json:"scaleY"`
	HasRotatingPoint bool    `xml:"hasRotatingPoint" json:"hasRotatingPoint"`
}

// Ident is a unique identification
type Ident struct {
	Value string `xml:"value" json:"Value"`
	Type  string `xml:"type,attr" json:"Type"`
}

// Source defines a source with copyright and source information
type Source struct {
	Value       string `xml:"value" json:"Value"`
	Text        string `xml:"text" json:"Text"`
	Title       string `xml:"title" json:"Title"`
	Attribution string `xml:"attribution" json:"Attribution"`
}

// Link for a part of the blog
type Link struct {
	URL   string `xml:"url" json:"Url"`
	Text  string `xml:"text" json:"Text"`
	Title string `xml:"title" json:"Title"`
}
