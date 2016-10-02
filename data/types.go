package data

import "fmt"

type Picture struct {
	ID         string `xml:"id,attr" json:"ID" schema:"id"`
	File       string `xml:"file" json:"File" schema:"file"`
	Title      string `xml:"title" json:"Title" schema:"title"`
	Topic      string `xml:"topic" json:"Topic" schema:"topic"`
	Text       string `xml:"text" json:"Text" schema:"text"`
	Areas      []Area `xml:"areas" json:"Areas" schema:"areas"`
	Captured   int    `xml:"captured" json:"Captured" schema:"captured"` //Year, when picture was digitalized
	Place      string `xml:"place" json:"Place" schema:"place"`          //Place where the picture was issued
	YearIssued string `xml:"yearIssued" json:"YearIssued" schema:"yearIssued"`
	Persons    []int  `xml:"persons" json:"Persons" schema:"persons"`
	Links      []Link `xml:"links" json:"Links" schema:"links"`
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
	ID         int      `xml:"id,attr" json:"personID" schema:"personID"`
	MasterID   int      `xml:"master,attr" json:"masterID" schema:"masterID"`
	Type       string   `xml:"type,attr" json:"Type" schema:"type"`
	NameFamily string   `xml:"name>family" json:"FamilyName"`
	NameGiven  string   `xml:"name>given" json:"GivenName"`
	FullName   string   `xml:"fullName" json:"FullName" schema:"fullName"`
	GND        int      `xml:"gnd" json:"GND" schema:"GND"`
	Pictures   []string `xml:"pictures" json:"Pictures" schema:"pictures"`
	Links      []Link   `xml:"links" json:"Links" schema:"links"`
}

// Identify implements the Identifier interface for loading and saving
func (p *Person) Identify() string {
	return fmt.Sprintf("%s%s_%04d", p.NameFamily, p.NameGiven, p.ID)
}

// TypeName implements the Identifier inteface for loading and saving
func (p *Person) TypeName() string {
	return "personData"
}

type Area struct {
	ID      string   `xml:"id,attr" json:"areaID" schema:"areaID"`
	Rect    Fabric   `xml:"rect" json:"rect" schema:"rect"`
	Shape   string   `xml:"shape" json:"Shape" schema:"shape"`
	Coords  string   `xml:"coords" json:"Coords" schema:"coords"`
	Persons []Person `xml:"persons" json:"Persons" schema:"persons"`
	Text    string   `xml:"text" json:"Text" schema:"text"`
	Links   []Link   `xml:"links" json:"Links" schema:"links"`
}

type Fabric struct {
	Type             string  `xml:"type" json:"type"`
	Left             float32 `xml:"left" json:"left" schema:"left"`
	Top              float32 `xml:"top" json:"top" schema:"top"`
	Width            int     `xml:"width" json:"width" schema:"width"`
	Height           int     `xml:"height" json:"height" schema:"height"`
	Fill             string  `xml:"fill" json:"fill"`
	Opacity          float32 `xml:"opacity" json:"opacity"`
	ScaleX           float32 `xml:"scaleX" json:"scaleX" schema:"scaleX"`
	ScaleY           float32 `xml:"scaleY" json:"scaleY" schema:"scaleY"`
	HasRotatingPoint bool    `xml:"hasRotatingPoint" json:"hasRotatingPoint"`
	CanvasWidth      int     `xml:"canvasWidth" json:"canvasWidth" schema:"canvasWidth"`
	CavasHeight      int     `xml:"canvasHeight" json:"canvasHeight" schema:"canvasHeight"`
}

type Link struct {
	URL   string `xml:"url" json:"Url" schema:"url"`
	Text  string `xml:"text" json:"Text" schema:"text"`
	Title string `xml:"title" json:"Title" schema:"title"`
}
