// Package data contains the structure of all the data and is used
// to encode/decode JSON or XML format
package data

type Picture struct {
	ID         string   `xml:"id,attr" json:"ID"`
	Title      string   `xml:"title" json:"Title"`
	Topic      string   `xml:"topic" json:"Topic"`
	Text       string   `xml:"text" json:"Text"`
	Areas      []Area   `xml:"areas" json:"Areas"`
	Captured   int      `xml:"captured" json:"Captured"` //Year, when picture was digitalized
	Place      string   `xml:"place" json:"Place"`       //Place where the picture was issued
	YearIssued string   `xml:"yearIssued" json:"YearIssued"`
	Persons    []Person `xml:"persons" json:"Persons"`
}

type Area struct {
	ID      string   `xml:"id,attr" json:"areaID"`
	Shape   string   `xml:"shape" json:"Shape"`
	Coords  string   `xml:"coords" json:"Coords"`
	Persons []Person `xml:"persons" json:"Persons"`
	Text    string   `xml:"text" json:"Text"`
}

type Person struct {
	ID   string `xml:"id,attr" json:"personID"`
	Type string `xml:"type,attr" json:"Type"`
	//NameFamily  string `xml:"name>family" json:"FamilyName"`
	//NameGiven   string `xml:"name>given" json:"GivenName"`
	FullName string `xml:"fullName" json:"FullName"`
	GND      int    `xml:"gnd" json:"GND"`
}

//Lister interface is used to get a list of all pictures inside a folder
type Lister interface {
	List() (*[]Picture, error)
}

//Loader loads the data of a picture the input can be a path or an id.
type Loader interface {
	Load(string) *Picture
}

//Saver saves the data of a picture
type Saver interface {
	Save(string) bool
}

//Sourcer is the abstraction for the different source implementations
type Sourcer interface {
	Lister
	Loader
	Saver
}
