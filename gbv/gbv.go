// Package gbv ist eine Schnittstelle zu http://unapi.gbv.de
// Anhand der PPN werden die aktuellen Daten von der Seite
// geladen.
package gbv

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/kupferstich/datatool/mods"
)

func GetModByPPN(ppn string) (*mods.Mods, error) {
	if strings.HasPrefix(ppn, "PPN") {
		ppn = string(ppn[3:])
	}
	var m mods.Mods
	url := fmt.Sprintf("http://unapi.gbv.de/?id=gvk:ppn:%s&format=mods", ppn)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal([]byte(data), &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
