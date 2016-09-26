package stabi

import (
	"github.com/kupferstich/datatool/data"
	"github.com/kupferstich/datatool/mods"
)

// NewDataPicture is an adapter from mods structure to picture
func NewDataPicture(m *mods.Mets) *data.Picture {
	var pic data.Picture
	pic.ID = m.Mods.RecordIdentifier
	pic.Title = getTitleInfo(m.Mods.TitleInfos, "")
	pic.Topic = getTitleInfo(m.Mods.TitleInfos, "alternative")
	pic.YearIssued = m.Mods.OriginInfo.DateIssued
	for _, name := range m.Mods.Names {
		var p data.Person
		p.Type = name.Type
		p.FullName = name.DisplayForm
		pic.Persons = append(pic.Persons, p)
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
