// Package gbv ist eine Schnittstelle zu http://unapi.gbv.de
// Anhand der PPN werden die aktuellen Daten von der Seite
// geladen.
package gbv

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
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
	log.Println("Loading ", url)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("API did not find an entry to that PPN")
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
