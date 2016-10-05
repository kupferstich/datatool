package data

import "fmt"

type Picture struct {
	ID         string `xml:"id,attr" json:"ID"`
	File       string `xml:"file" json:"File"`
	Title      string `xml:"title" json:"Title"`
	Topic      string `xml:"topic" json:"Topic"`
	Text       string `xml:"text" json:"Text"`
	Areas      []Area `xml:"areas" json:"Areas"`
	Captured   int    `xml:"captured" json:"Captured"` //Year, when picture was digitalized
	Place      string `xml:"place" json:"Place"`       //Place where the picture was issued
	YearIssued string `xml:"yearIssued" json:"YearIssued"`
	Persons    []int  `xml:"persons" json:"Persons"`
	Links      []Link `xml:"links" json:"Links"`
	Status     string `xml:"status" json:"Status"`
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
	ID         int      `xml:"id,attr" json:"personID"`
	MasterID   int      `xml:"master,attr" json:"masterID"`
	Type       string   `xml:"type,attr" json:"Type"`
	NameFamily string   `xml:"name>family" json:"FamilyName"`
	NameGiven  string   `xml:"name>given" json:"GivenName"`
	FullName   string   `xml:"fullName" json:"FullName"`
	YearBirth  int      `xml:"yearBirth" json:"YearBirth"`
	YearDeath  int      `xml:"yearDeath" json:"YearDeath"`
	CityBirth  string   `xml:"cityBirth" json:"CityBirth"`
	CityDeath  string   `xml:"cityDeath" json:"CityDeath"`
	ProfilePic string   `xml:"profilePic" json:"ProfilePic"`
	GND        int      `xml:"gnd" json:"GND"`
	Text       string   `xml:"text" json:"Text"`
	Pictures   []string `xml:"pictures" json:"Pictures"`
	Links      []Link   `xml:"links" json:"Links"`
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
	ID      string `xml:"id,attr" json:"areaID"`
	Rect    Fabric `xml:"rect" json:"rect"`
	Shape   string `xml:"shape" json:"Shape"`
	Coords  string `xml:"coords" json:"Coords"`
	Persons []int  `xml:"persons" json:"Persons"`
	Text    string `xml:"text" json:"Text"`
	Links   []Link `xml:"links" json:"Links"`
}

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
	CanvasWidth      int     `xml:"canvasWidth" json:"canvasWidth"`
	CavasHeight      int     `xml:"canvasHeight" json:"canvasHeight"`
}

type Link struct {
	URL   string `xml:"url" json:"Url"`
	Text  string `xml:"text" json:"Text"`
	Title string `xml:"title" json:"Title"`
}
