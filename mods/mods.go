// Package mods provides the MODS/METS structure for the XML files of the
// Staats- und Universitaets Bibliothek Hamburg
// The package provides not all tags and attributes, which are inside the
// files.
package mods

import "encoding/xml"

// Mets struct represents the header of the XML file.
type Mets struct {
	XMLName xml.Name `xml:"mets"`
	MetsHDR string   `xml:"metsHdr"`
	Mods    Mods     `xml:"dmdSec>mdWrap>xmlData>mods"`
}

// Mods provides all the relevant data of the XML file
type Mods struct {
	XMLName              xml.Name              `xml:"mods"`
	Names                []Name                `xml:"name"`
	RecordIdentifier     string                `xml:"recordInfo>recordIidentifier"`
	Identifiers          []Identifier          `xml:"identifier"`
	PhysicalDescriptions []PhysicalDescription `xml:"physicalDescription"`
	PhysicalLocaltion    string                `xml:"location>physicalLocation"`
	ShelfLocator         string                `xml:"location>shelfLocator"`
	OriginInfo           OriginInfo            `xml:"originInfo"`
	Classification       []string              `xml:"classification"`
	TitleInfos           []TitleInfo           `xml:"titleInfo"`
}

// Name structure for the methadata
type Name struct {
	Type        string     `xml:"type,attr"`
	RoleTerm    string     `xml:"role>roleTerm"`
	NameParts   []NamePart `xml:"namePart"`
	DisplayForm string     `xml:"displayForm"`
}

// NamePart provides the information about one element of a name.
// NameType can be "family" or "given"
type NamePart struct {
	NameType string `xml:"type,attr"`
	NamePart string `xml:",innerxml"`
}

// Identifier represents different kinds of identifiers
type Identifier struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",innerxml"`
}

// PhysicalDescription has free text elements to describe the picture
type PhysicalDescription struct {
	Extend []string `xml:"extent"`
}

// OriginInfo provides date, place and publisher.
// If a value is inside [] the value is set from the Stabi
type OriginInfo struct {
	PlaceTerm  string `xml:"place>placeTerm"`
	DateIssued string `xml:"dateIssued"`
	Publisher  string `xml:"publisher"`
}

// TitleInfo provides different kind of titles. If Type is empty, it is
// the original title, which can be found on the picture.
// Alternative titles could be set by the Stabi.
type TitleInfo struct {
	Type  string `xml:"type,attr"`
	Title string `xml:"title"`
}
