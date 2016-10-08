package gbv

import (
	"fmt"
	"testing"
)

func TestGetModByPPN(t *testing.T) {
	m, err := GetModByPPN("PPN782360696")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", m)
	_, err = GetModByPPN("PPN7333382360696")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%#v", m)
}
