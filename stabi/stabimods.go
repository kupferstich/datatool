package stabi

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kupferstich/datatool/data"
	"github.com/kupferstich/datatool/gbv"
	"github.com/kupferstich/datatool/mods"
)

// NewDataPicture is an adapter from mods structure to picture
func NewDataPicture(m *mods.Mets, pdb data.PersonDBer) *data.Picture {
	var pic data.Picture
	pic.ID = m.Mods.RecordIdentifier
	pic.Title = getTitleInfo(m.Mods.TitleInfos, "")
	pic.Topic = getTitleInfo(m.Mods.TitleInfos, "alternative")
	pic.YearIssued = m.Mods.OriginInfo.DateIssued
	// Ask the gbv api to get the current infos about the persons
	gbvMod, err := gbv.GetModByPPN(pic.ID)
	var modsNames []mods.Name
	var gnd int
	if err != nil || len(gbvMod.Names) == 0 {
		modsNames = m.Mods.Names
		gnd = 0
	} else {
		modsNames = gbvMod.Names
		gnd = -1
	}
	for _, name := range modsNames {
		var p data.Person
		p.Type = name.Type
		//p.FullName = name.DisplayForm
		p.NameFamily = getNamePart(name.NameParts, "family")
		p.NameGiven = getNamePart(name.NameParts, "given")
		p.FullName = fmt.Sprintf("%s, %s", p.NameFamily, p.NameGiven)
		if gnd == -1 {
			parts := strings.Split(name.ValueURI, "/")
			gnd, _ = strconv.Atoi(parts[len(parts)-1])
		}
		p.GND = gnd
		p.Pictures = append(p.Pictures, pic.ID)
		// Check if person is in db. If there is an entry to the ID
		// the data is not going to be saved here.
		_, ok := pdb.GetPerson(p.GetID())
		if !ok {
			err := pdb.SavePerson(&p)
			if err != nil {
				log.Println(err)
			}
		}
		// SavePerson saves the data and adds the ID
		pic.Persons = append(pic.Persons, p.GetID())
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
