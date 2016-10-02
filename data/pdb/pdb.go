// pdb.go is the file where everything about storing the person
// data is found.
// Every person gets a intern numeric id. Because it should be
// possible to strore two different persons with the same name.
// The pdb stores the data of each person inside a single folder.
// The foldername contains the name and the id.

package pdb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/kupferstich/datatool/data"
)

// ErrGotNilPointer is, when a function gets a nil pointer
var ErrGotNilPointer = errors.New("Got a nil pointer as input.")

// PersonDB is the struct for storing all the persons.
type PersonDB struct {
	Root        string              `json:"rootFolder"`
	PictureRoot string              `json:"pictureRootFolder"`
	NextID      int                 `json:"nextID"`
	Persons     map[int]data.Person `json:"Persons"`
}

func New(root string) *PersonDB {
	pdb := PersonDB{Root: root, NextID: 1}
	pdb.Persons = make(map[int]data.Person)
	return &pdb
}

func Load(root string) (*PersonDB, error) {
	pdb := New(root)
	err := filepath.Walk(root, func(fpath string, info os.FileInfo, ierr error) error {
		if info.IsDir() {
			return nil
		}
		if !strings.EqualFold(filepath.Base(fpath), "personData.json") {
			return nil
		}
		b, err := ioutil.ReadFile(fpath)
		if err != nil {
			log.Println(err)
			return err
		}
		p := &data.Person{}
		err = json.Unmarshal(b, p)
		if err != nil {
			log.Println("pdb:load:", err, p, string(b))
			return err
		}
		if p.ID >= pdb.NextID {
			pdb.NextID = p.ID + 1
		}
		pdb.Persons[p.ID] = *p
		return nil
	})
	return pdb, err
}

// AddPerson adda data.Person to the pdb instance. If the person has no ID a new ID
// is created and set to the person.
func (pdb *PersonDB) AddPerson(p *data.Person) {
	// Try to find the person by name
	id, ok := pdb.FindPerson(p)
	if ok && p.ID == 0 {
		// If allready in db set the id to the person.
		p.ID = id
	}
	if p.ID == 0 {
		// New ID if until here no id is found.
		p.ID = pdb.NextID
		pdb.NextID++
	}
	if p.ID > pdb.NextID {
		// Set the nextID if there are IDs
		pdb.NextID = p.ID + 1
	}
	pp, ok := pdb.Persons[p.ID]
	if ok {
		// If there is that person in the db set the Pictures.
		p.Pictures = append(p.Pictures, pp.Pictures...)
		p.Pictures = removeDuplicates(p.Pictures)
	}
	pdb.Persons[p.ID] = *p
}

// GetPerson by Person.ID. Second parameter returns true if an entry is found.
// If there is no such ID inside the pdb the function return nil, false
func (pdb *PersonDB) GetPerson(id int) (*data.Person, bool) {
	p, ok := pdb.Persons[id]
	// If a MasterID is set, there is a master entry. This should be returned.
	if p.MasterID > 0 {
		m, mok := pdb.Persons[p.MasterID]
		if mok {
			return &m, mok
		}
	}
	return &p, ok
}

// SavePerson takes a Person and saves the data inside the struct. That method should
// always be used to save Person data to ensure that the person always get an ID.
// If the input person has no ID the NextID is set to the person and counter +1.
func (pdb *PersonDB) SavePerson(p *data.Person) error {
	if p == nil {
		return ErrGotNilPointer
	}
	pdb.AddPerson(p)
	return data.SaveType(p, pdb.Root)
}

// FindPerson takes a Person and searches with the name if there is an entry in
// the db. If there is a person with that name it returns the ID and true.
// If no person is found it return 0 and false
func (pdb *PersonDB) FindPerson(p *data.Person) (int, bool) {
	if p == nil {
		return 0, false
	}
	for _, pp := range pdb.Persons {
		// Not only the name is important. Because of a possible master entry the
		// id has to have no master ID set.
		if pp.FullName == p.FullName && pp.MasterID == 0 {
			return pp.ID, true
		}
	}
	return 0, false
}

func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}
	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}
