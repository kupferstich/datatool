// Package data contains the structure of all the data and is used
// to encode/decode JSON or XML format
package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ErrFileNotFound is used when try to load a file and that file does not exist
var ErrFileNotFound = errors.New("File not found.")

// Lister interface is used to get a list of all pictures inside a folder
type Lister interface {
	List() (*[]Picture, error)
}

// Identifier is used to get the ID of a type. That ID is used to generate the
// filepath for saving the data.
type Identifier interface {
	Identify() string
	TypeName() string
}

// PersonDBer is the interface for a simple PersonDB
type PersonDBer interface {
	AddPerson(*Person)
	GetPerson(string) (*Person, bool)
	GetAll() map[string]Person
	SavePerson(p *Person) error
}

// LoadType loads the data from the data folder for a given type
// That type needs to implement the identifier interface.
func LoadType(i Identifier, root string) error {
	fpath := MakePath(i, root)
	if _, err := os.Stat(fpath); os.IsNotExist(err) {
		return ErrFileNotFound
	}
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, i)
	if err != nil {
		return err
	}
	return nil
}

// SaveType stores the data from a type into the data folder. The type needs
// to implement the Identifier interface.
func SaveType(i Identifier, root string) error {
	b, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return err
	}
	fpath := MakePath(i, root)
	os.MkdirAll(filepath.Dir(fpath), 0777)
	err = ioutil.WriteFile(fpath, b, 0777)
	if err != nil {
		return err
	}
	return nil
}

// MakePath generates the path to a identifier, when a root folder is given.
func MakePath(i Identifier, root string) string {
	return filepath.Join(
		root,
		i.Identify(),
		fmt.Sprintf("%s.json", i.TypeName()),
	)
}
