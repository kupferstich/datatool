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

//Lister interface is used to get a list of all pictures inside a folder
type Lister interface {
	List() (*[]Picture, error)
}

type Identifier interface {
	Identify() string
}

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

func MakePath(i Identifier, root string) string {
	ident := i.Identify()
	return filepath.Join(
		root,
		ident,
		fmt.Sprintf("%s.json", ident),
	)
}
