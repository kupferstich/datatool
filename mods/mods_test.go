package mods

import (
	"encoding/xml"
	"fmt"
	"testing"
)

var data = `
<mets:mets xsi:schemaLocation="  http://www.loc.gov/standards/mix/mix.xsd">
 <mets:metsHdr CREATEDATE="2014-06-03T13:50:30">12</mets:metsHdr>
 <mets:dmdSec ID="DMDLOG_0000">
  <mets:mdWrap MDTYPE="MODS">
    <mets:xmlData>
     <mods:mods><mods:name type="personal"></mods:name><mods:name type="personal"><mods:role><mods:roleTerm authority="marcrelator" type="code">art</mods:roleTerm></mods:role><mods:namePart type="family">Galle</mods:namePart><mods:namePart type="given">Philippe</mods:namePart><mods:displayForm>Galle, Philippe</mods:displayForm></mods:name><mods:name type="personal"><mods:role><mods:roleTerm authority="marcrelator" type="code">art</mods:roleTerm></mods:role><mods:namePart type="family">Peeters</mods:namePart><mods:namePart type="given">Marten</mods:namePart><mods:displayForm>Peeters, Marten</mods:displayForm></mods:name><mods:recordInfo><mods:recordIdentifier source="gbv-ppn">PPN78289027X</mods:recordIdentifier></mods:recordInfo><mods:identifier type="purl">http://resolver.sub.uni-hamburg.de/goobi/PPN78289027X</mods:identifier><mods:identifier type="urn">urn:nbn:de:gbv:18-5-PPN78289027X5</mods:identifier><mods:identifier type="PPNanalog">PPN757658636</mods:identifier><mods:physicalDescription><mods:extent>Blattgr. 21 x 25 cm</mods:extent><mods:extent>Kupferstich</mods:extent><mods:extent>1 Bl.</mods:extent></mods:physicalDescription><mods:location><mods:physicalLocation>Staats- und Universitätsbibliothek Hamburg</mods:physicalLocation><mods:shelfLocator>Kupfer 3: 1</mods:shelfLocator></mods:location><mods:originInfo><mods:place><mods:placeTerm type="text">[Antwerpen]</mods:placeTerm></mods:place><mods:dateIssued encoding="iso8601" keyDate="yes">1565 [ca. 1565]</mods:dateIssued><mods:publisher>Martinus Petri</mods:publisher></mods:originInfo><mods:classification authority="subhh">Kupferstichsammlung</mods:classification><mods:classification authority="subhh">Graphik</mods:classification><mods:classification authority="subhh">700 - Kunst</mods:classification><mods:titleInfo><mods:title>Maria Mater Dei</mods:title></mods:titleInfo><mods:titleInfo type="alternative"><mods:title>[Acht berühmte Frauen aus dem Alten und Neuen Testament]</mods:title></mods:titleInfo></mods:mods></mets:xmlData></mets:mdWrap></mets:dmdSec><mets:dmdSec ID="DMDPHYS_0000"></mets:dmdSec><mets:amdSec ID="AMD"></mets:amdSec><mets:fileSec></mets:fileSec><mets:structMap TYPE="LOGICAL"></mets:structMap><mets:structMap TYPE="PHYSICAL"></mets:structMap><mets:structLink></mets:structLink></mets:mets>
`

var dataMods = `
<mods:mods><mods:name type="personal"><mods:role>
<mods:roleTerm authority="marcrelator" type="code">art</mods:roleTerm></mods:role>
<mods:namePart type="family">Heemskerk</mods:namePart>
<mods:namePart type="given">Maarteen van</mods:namePart>
<mods:displayForm>Heemskerk, Maarteen van</mods:displayForm></mods:name>
<mods:name type="personal"><mods:role>
<mods:roleTerm authority="marcrelator" type="code">art</mods:roleTerm></mods:role>
<mods:namePart type="family">Galle</mods:namePart>
<mods:namePart type="given">Philippe</mods:namePart>
<mods:displayForm>Galle, Philippe</mods:displayForm></mods:name>
<mods:name type="personal"><mods:role>
<mods:roleTerm authority="marcrelator" type="code">art</mods:roleTerm></mods:role>
<mods:namePart type="family">Peeters</mods:namePart>
<mods:namePart type="given">Marten</mods:namePart>
<mods:displayForm>Peeters, Marten</mods:displayForm></mods:name>
<mods:recordInfo>
<mods:recordIdentifier source="gbv-ppn">PPN78289027X</mods:recordIdentifier>
</mods:recordInfo>
<mods:identifier type="purl">http://resolver.sub.uni-hamburg.de/goobi/PPN78289027X</mods:identifier>
<mods:identifier type="urn">urn:nbn:de:gbv:18-5-PPN78289027X5</mods:identifier>
<mods:identifier type="PPNanalog">PPN757658636</mods:identifier>
<mods:physicalDescription><mods:extent>Blattgr. 21 x 25 cm</mods:extent>
<mods:extent>Kupferstich</mods:extent><mods:extent>1 Bl.</mods:extent>
</mods:physicalDescription>
<mods:location>
  <mods:physicalLocation>Staats- und Universitätsbibliothek Hamburg</mods:physicalLocation>
  <mods:shelfLocator>Kupfer 3: 1</mods:shelfLocator></mods:location>
<mods:originInfo>
 <mods:place><mods:placeTerm type="text">[Antwerpen]</mods:placeTerm></mods:place>
 <mods:dateIssued encoding="iso8601" keyDate="yes">1565 [ca. 1565]</mods:dateIssued>
 <mods:publisher>Martinus Petri</mods:publisher>
</mods:originInfo>
<mods:classification authority="subhh">Kupferstichsammlung</mods:classification>
<mods:classification authority="subhh">Graphik</mods:classification>
<mods:classification authority="subhh">700 - Kunst</mods:classification>
<mods:titleInfo><mods:title>Maria Mater Dei</mods:title></mods:titleInfo>
<mods:titleInfo type="alternative"><mods:title>[Acht berühmte Frauen aus dem Alten und Neuen Testament]</mods:title>
</mods:titleInfo></mods:mods>
`

func TestUnmarshal(t *testing.T) {
	var v Mets
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v", v)
	/*np := NamePart{"Vorname", "Mueller"}
	a := Mets{
		Modss: []Mods{
			Mods{
				Names: []Name{
					Name{
						NameParts: []NamePart{
							np,
						},
					},
					Name{
						NameParts: []NamePart{
							np,
						},
					},
				},
			},
		},
	}
	fmt.Println(a)*/
	//output, err := xml.MarshalIndent(a, "  ", "    ")
	//os.Stdout.Write(output)
}
