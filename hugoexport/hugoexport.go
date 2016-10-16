// Package export collects all the data and copies them into the
// different folders. For the export the hugo structure is expected.
// The export subfolders are:
// - content
// - data
package hugoexport

import (
	"path/filepath"

	"github.com/disintegration/imaging"
)

// Size describes the size of an image
type Size struct {
	Width  int
	Height int
}

// JSONArtistSubfolder is where the person json files will be copied
var JSONArtistSubfolder = filepath.Join("data", "artist")

// JSONArtworkSubfolder is where the pictures json files will be copied
var JSONArtworkSubfolder = filepath.Join("data", "artwork")

// ImgArtworkSubfolder is where the images will be exported
var ImgArtworkSubfolder = filepath.Join("static", "img", "artwork")

// ImgArtworkSrcFilename is the filename of the picture inside the
// src (pictures) folder
var ImgArtworkSrcFilename = "00000001.jpg"

// ResizeSizes is for the sizes which will be created
var ResizeSizes = map[string]Size{
	"big":    {700, 700},
	"medium": {400, 400},
	"small":  {250, 250},
	"thumb":  {125, 125},
}

// ResizeFilter defines thee resize algorithm
var ResizeFilter = imaging.Lanczos
