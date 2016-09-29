package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Conf stores the configuration
var Conf Options

// Options is the structure of the config file
type Options struct {
	ServerPort         string `yaml:"ServerPort"`
	FilesFolder        string `yaml:"FilesFolder"`
	SourceFolder       string `yaml:"SourceFolder"`
	DataFolderPictures string `yaml:"DataFolderPictures"`
	HugoFolder         string `yaml:"HugoFolder"`
}

type dependencies struct{}

func loadConf() {
	b, err := ioutil.ReadFile(*ConfFile)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(b, &Conf)
	if err != nil {
		log.Fatal(err)
	}
}
