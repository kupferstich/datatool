package mods

import "encoding/xml"

type Mets struct {
	XMLName xml.Name `xml:"mets"`
	MetsHDR string   `xml:"metsHdr"`
	Modss   []Mods   `xml:"dmdSec>mdWrap>xmlData>mods"`
}
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

type Name struct {
	Type        string     `xml:"type,attr"`
	RoleTerm    string     `xml:"role>roleTerm"`
	NameParts   []NamePart `xml:"namePart"`
	DisplayForm string     `xml:"displayForm"`
}
type NamePart struct {
	NameType string `xml:"type,attr"`
	NamePart string `xml:",innerxml"`
}

type Identifier struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",innerxml"`
}

type PhysicalDescription struct {
	Extend []string `xml:"extent"`
}

type OriginInfo struct {
	PlaceTerm  string `xml:"place>placeTerm"`
	DateIssued string `xml:"dateIssued"`
	Publisher  string `xml:"publisher"`
}

type TitleInfo struct {
	Type  string `xml:"type,attr"`
	Title string `xml:"title"`
}

func Unmarshal(b []byte, v interface{}) error {
	err := xml.Unmarshal(b, v)
	return err

}
