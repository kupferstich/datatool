// Package data contains the structure of all the data and is used
// to encode/decode JSON or XML format
package data

type Picture struct {
	ID         string   `xml:"id,attr" yaml:"ID"`
	Title      string   `xml:"title" yaml:"Title"`
	Topic      string   `xml:"topic" yaml:"Topic"`
	Text       string   `xml:"text" yaml:"Text"`
	Areas      []Area   `xml:"areas" yaml:"Areas"`
	Captured   int      `xml:"captured" yaml:"Captured"` //Year, when picture was digitalized
	Place      string   `xml:"place" yaml:"Place"`       //Place where the picture was issued
	YearIssued string   `xml:"yearIssued" yaml:"YearIssued"`
	Persons    []Person `xml:"persons" yaml:"Persons"`
}

type Area struct {
	ID      string   `xml:"id,attr" yaml:"areaID"`
	Shape   string   `xml:"shape" yaml:"Shape"`
	Coords  string   `xml:"coords" yaml:"Coords"`
	Persons []Person `xml:"persons" yaml:"Persons"`
	Text    string   `xml:"text" yaml:"Text"`
}

type Person struct {
	ID   string `xml:"id,attr" yaml:"personID"`
	Type string `xml:"type,attr" yaml:"Type"`
	//NameFamily  string `xml:"name>family" yaml:"FamilyName"`
	//NameGiven   string `xml:"name>given" yaml:"GivenName"`
	FullName string `xml:"fullName" yaml:"FullName"`
	GND      int    `xml:"gnd" yaml:"GND"`
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
