package pdb

import (
	"testing"

	"github.com/kupferstich/datatool/data"
)

var persons = []data.Person{
	data.Person{
		ID:       1,
		FullName: "Person1",
	},
	data.Person{
		ID:       20,
		FullName: "Person20",
	},
	data.Person{
		ID:       3,
		FullName: "Person3",
	},
}

func TestAddPerson(t *testing.T) {
	db := New("rootPath")
	for _, p := range persons {
		db.AddPerson(&p)
	}
	if db.Persons[20].FullName != "Person20" {
		t.Fatalf("Expected: %v, Got: %v", "Person20", db.Persons[20].FullName)
	}
}

func TestFindPerson(t *testing.T) {
	db := New("rootPath")
	for _, p := range persons {
		db.AddPerson(&p)
	}
	p := persons[2]
	id, _ := db.FindPerson(&p)
	if id != p.ID {
		t.Fatalf("Expected: %v, Got: %v", p.ID, id)
	}
}
