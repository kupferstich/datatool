package stabi

import (
	"log"

	"github.com/kupferstich/datatool/data"
	"github.com/kupferstich/datatool/mods"
)

// NewDataPicture is an adapter from mods structure to picture
func NewDataPicture(m *mods.Mets, pdb data.PersonDBer) *data.Picture {
	var pic data.Picture
	pic.ID = m.Mods.RecordIdentifier
	pic.Title = getTitleInfo(m.Mods.TitleInfos, "")
	pic.Topic = getTitleInfo(m.Mods.TitleInfos, "alternative")
	pic.YearIssued = m.Mods.OriginInfo.DateIssued
	for _, name := range m.Mods.Names {
		var p data.Person
		p.Type = name.Type
		p.FullName = name.DisplayForm
		p.NameFamily = getNamePart(name.NameParts, "family")
		p.NameGiven = getNamePart(name.NameParts, "given")

		err := pdb.SavePerson(&p)
		if err != nil {
			log.Println(err)
		}
		// SavePerson saves the data and adds the ID
		pic.Persons = append(pic.Persons, p.ID)
	}
	return &pic
}

func getTitleInfo(ti []mods.TitleInfo, s string) string {
	for _, t := range ti {
		if t.Type == s {
			return t.Title
		}
	}
	return ""
}

func getNamePart(nameParts []mods.NamePart, nameType string) string {
	for _, np := range nameParts {
		if np.NameType == nameType {
			return np.NamePart
		}
	}
	return ""
}
