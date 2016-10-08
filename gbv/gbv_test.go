package gbv

import (
	"fmt"
	"testing"
)

func TestGetModByPPN(t *testing.T) {
	m, err := GetModByPPN("782360696")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%#v", m.Names)
	for _, n := range m.Names {
		fmt.Println(n.NameParts, n.ValueURI)
	}
	_, err = GetModByPPN("PPN7333382360696")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%#v", m)
}
