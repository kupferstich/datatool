// Package data contains the structure of all the data and is used
// to encode/decode JSON or XML format
package data

type Picture struct {
	ID         string `xml:"id,attr"`
	Title      string `xml:"title"`
	Text       string `xml:"text"`
	Areas      []Area `xml:"areas"`
	Captured   int    `xml:"captured"` //Year, when picture was digitalized
	Place      string `xml:"place"`    //Place where the picture was issued
	YearIssued int    `xml:"yearIssued"`
}

type Area struct {
	ID      string   `xml:"id,attr"`
	Shape   string   `xml:"shape"`
	Coords  string   `xml:"coords"`
	Persons []Person `xml:"persons"`
	Text    string   `xml:"text"`
}

type Person struct {
	ID         string `xml:"id,attr"`
	NameFamily string `xml:"name>family"`
	NameGiven  string `xml:"name>given"`
	GND        int    `xml:"gnd"`
}
