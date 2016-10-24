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
	"strconv"
	"strings"

	"github.com/kupferstich/datatool/data"
)

// ErrGotNilPointer is, when a function gets a nil pointer
var ErrGotNilPointer = errors.New("Got a nil pointer as input.")

// ErrPersonDeleted describes, when a person is not any more inside the db.
var ErrPersonDeleted = errors.New("Person instance inside the db was deleted.")

// PersonDB is the struct for storing all the persons.
type PersonDB struct {
	Root        string                 `json:"rootFolder"`
	PictureRoot string                 `json:"pictureRootFolder"`
	NextID      int                    `json:"nextID"`
	Persons     map[string]data.Person `json:"Persons"`
}

func New(root, pictureRoot string) *PersonDB {
	pdb := PersonDB{Root: root, PictureRoot: pictureRoot, NextID: 1}
	pdb.Persons = make(map[string]data.Person)
	return &pdb
}

func Load(root, pictureRoot string) (*PersonDB, error) {
	pdb := New(root, pictureRoot)
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
		pdb.Persons[p.GetID()] = *p
		return nil
	})
	return pdb, err
}

// SavePerson takes a Person and saves the data inside the struct. That method should
// always be used to save Person data to ensure that the person always get an ID.
// If the input person has no ID the NextID is set to the person and counter +1.
func (pdb *PersonDB) SavePerson(p *data.Person) error {
	if p == nil {
		return ErrGotNilPointer
	}
	pdb.AddPerson(p)
	err := pdb.EditPerson(p)
	if err != ErrPersonDeleted {
		return data.SaveType(p, pdb.Root)
	}
	return nil
}

// AddPerson adda data.Person to the pdb instance. If the person has no ID a new ID
// is created and set to the person.
func (pdb *PersonDB) AddPerson(p *data.Person) {
	// Try to find the person
	pp, ok := pdb.FindPerson(p)
	if ok && p.ID == 0 {
		// If allready in db set the id to the person. This is used for the initial
		// runs. When e.g. the data is read the first time out of the xml files.
		// Where just the name is given.
		// This means that the person is edited.
		p.ID = pp.ID
	}
	if p.ID > pdb.NextID {
		// Set the nextID if there are IDs
		pdb.NextID = p.ID + 1
	}
	if p.ID == 0 {
		// New ID if until here no id is found.
		p.ID = pdb.NextID
		pdb.NextID++
		pdb.Persons[p.GetID()] = *p
	}

}

// EditPerson handles if data inside the person has changed. It is compulsory
// that the person has a valid ID.
func (pdb *PersonDB) EditPerson(p *data.Person) error {
	dbPerson, ok := pdb.Persons[p.GetID()]
	if !ok {
		// Try to find the person over the name, because it is possible that
		// the ID changed. This happens, when e.g. a GND is added.
		// If there are not unique names that part can causes some strange
		// behaviour.
		fp, fok := pdb.FindPerson(p)
		if !fok {
			return nil
		}
		dbPerson = *fp
	}
	// The person is edited.
	// If there is that person in the db set the Pictures.
	p.Pictures = append(p.Pictures, dbPerson.Pictures...)
	p.Pictures = removeDuplicates(p.Pictures)

	// If the folderpath changed after edit
	if data.MakePath(&dbPerson, pdb.Root) != data.MakePath(p, pdb.Root) {
		// Rename the folder
		err := os.Rename(
			filepath.Dir(data.MakePath(&dbPerson, pdb.Root)),
			filepath.Dir(data.MakePath(p, pdb.Root)),
		)
		if err != nil {
			return err
		}
		// Save person at new location
		data.SaveType(p, pdb.Root)
		pdb.Persons[p.GetID()] = *p
		// If the ID changed the old ID must be removed from the db
		if dbPerson.GetID() != p.GetID() {
			delete(pdb.Persons, dbPerson.GetID())
		}
		return ErrPersonDeleted

	}

	// If the masterID is set or the ID changes the references at the pictures
	// had to be updated.
	if p.MasterID != 0 {
		// Save the person before updating references
		data.SaveType(p, pdb.Root)
		pdb.UpdatePictures(pdb.PictureRoot)
		// Remove the file of the person after the update
		err := os.RemoveAll(filepath.Dir(data.MakePath(p, pdb.Root)))
		if err != nil {
			log.Println(err)
		}
		delete(pdb.Persons, p.GetID())
		return ErrPersonDeleted
	}
	// Update the entry inside the pdb
	pdb.Persons[p.GetID()] = *p
	return nil
}

// GetPerson by Person.ID. Second parameter returns true if an entry is found.
// If there is no such ID inside the pdb the function return nil, false
func (pdb *PersonDB) GetPerson(id string) (*data.Person, bool) {
	var p data.Person
	if strings.HasPrefix(id, "GND") {
		p.GND = string(id[3:])
	} else {
		var err error
		p.ID, err = strconv.Atoi(id)
		if err != nil {
			log.Println(err)
		}
	}
	pp, ok := pdb.FindPerson(&p)
	p.ExtID = id
	if !ok {
		return nil, false
	}
	pdb.GetProfilePics(pp)
	return pp, ok
}

func (pdb *PersonDB) GetProfilePics(p *data.Person) {
	picExt := []string{".jpg", ".jpeg", ".png"}
	ProfilePics, _ := data.GetFiles(filepath.Dir(data.MakePath(p, pdb.Root)), picExt)
	// Store the db values in var, because the pic source could be renamed or
	// deleted, then the old value should be not loaded.
	dbProfilePics := p.ProfilePics
	// Delete the old values
	p.ProfilePics = make(map[string]data.Source)
	for _, profPic := range ProfilePics {
		dbSource, pok := dbProfilePics[profPic]
		// If file is already in the db set the old value
		if pok {
			p.ProfilePics[profPic] = dbSource
		} else {
			p.ProfilePics[profPic] = data.Source{Value: profPic}
		}
	}
}

// GetAll returns all the persons inside the DB as map
func (pdb *PersonDB) GetAll() map[string]data.Person {
	// If an entry has a master entry it is not listed
	all := make(map[string]data.Person)
	for k, p := range pdb.Persons {
		if p.MasterID == 0 {
			p.ExtID = p.GetID()
			all[k] = p
		}
	}
	return all
}

// FindPerson takes a Person and searches with the name if there is an entry in
// the db. If there is a person with that name it returns the ID and true.
// If no person is found it return 0 and false
func (pdb *PersonDB) FindPerson(p *data.Person) (*data.Person, bool) {
	if p == nil {
		return nil, false
	}
	// Check if there is allready an ID availiable and inside the db
	log.Println(p.GetID())
	dbp, ok := pdb.Persons[p.GetID()]
	if ok {
		return &dbp, true
	}
	for _, pp := range pdb.Persons {
		// Check for the intern ID and then for the name.
		// Not only the name is important. Because of a possible master entry the
		// id has to have no master ID set.
		if pp.ID == p.ID || (pp.FullName == p.FullName && pp.MasterID == 0) {
			return &pp, true
		}
	}
	return nil, false
}

// UpdatePictures regenerates the references to a person, which has a masterID.
func (pdb *PersonDB) UpdatePictures(root string) {
	pictures := data.LoadPictures(root)
	for _, pic := range pictures {
		for i, p := range pic.Persons {
			dbPerson, ok := pdb.GetPerson(p)
			if ok {
				pic.Persons[i] = dbPerson.GetID()
				pic.Persons = removeDuplicates(pic.Persons)
			}
		}

		data.SaveType(&pic, root)
	}
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
