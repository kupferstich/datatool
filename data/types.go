package data

type Picture struct {
	ID         string   `xml:"id,attr" json:"ID" schema:"id"`
	File       string   `xml:"file" json:"File" schema:"file"`
	Title      string   `xml:"title" json:"Title" schema:"title"`
	Topic      string   `xml:"topic" json:"Topic" schema:"topic"`
	Text       string   `xml:"text" json:"Text" schema:"text"`
	Areas      []Area   `xml:"areas" json:"Areas" schema:"areas"`
	Captured   int      `xml:"captured" json:"Captured" schema:"captured"` //Year, when picture was digitalized
	Place      string   `xml:"place" json:"Place" schema:"place"`          //Place where the picture was issued
	YearIssued string   `xml:"yearIssued" json:"YearIssued" schema:"yearIssued"`
	Persons    []Person `xml:"persons" json:"Persons" schema:"persons"`
	Links      []Link   `xml:"links" json:"Links" schema:"links"`
}

//Identify implements the Identifier interface for loading and saving
func (p *Picture) Identify() string {
	return p.ID
}

type Area struct {
	ID      string   `xml:"id,attr" json:"areaID" schema:"areaID"`
	Shape   string   `xml:"shape" json:"Shape" schema:"shape"`
	Coords  string   `xml:"coords" json:"Coords" schema:"coords"`
	Persons []Person `xml:"persons" json:"Persons" schema:"persons"`
	Text    string   `xml:"text" json:"Text" schema:"text"`
	Links   []Link   `xml:"links" json:"Links" schema:"links"`
}

type Person struct {
	ID         string `xml:"id,attr" json:"personID" schema:"personID"`
	Type       string `xml:"type,attr" json:"Type" schema:"type"`
	NameFamily string `xml:"name>family" json:"FamilyName"`
	NameGiven  string `xml:"name>given" json:"GivenName"`
	FullName   string `xml:"fullName" json:"FullName" schema:"fullName"`
	GND        int    `xml:"gnd" json:"GND" schema:"GND"`
	Links      []Link `xml:"links" json:"Links" schema:"links"`
}

type Link struct {
	URL   string `xml:"url" json:"Url" schema:"url"`
	Text  string `xml:"text" json:"Text" schema:"text"`
	Title string `xml:"title" json:"Title" schema:"title"`
}
