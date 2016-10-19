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

// ImgArtistSubfolder is where the images will be exported
var ImgArtistSubfolder = filepath.Join("static", "img", "artist")

// ContentArtworkSubfolder the folder for the md content files.
var ContentArtworkSubfolder = filepath.Join("content", "artwork")

// ContentArtistSubfolder the folder for the md content files.
var ContentArtistSubfolder = filepath.Join("content", "artist")

// ImgArtworkSrcFilename is the filename of the picture inside the
// src (pictures) folder
var ImgArtworkSrcFilename = "00000001.jpg"

// ResizeSizes is for the sizes which will be created
// thumb will be resized with the thumbnail method. So it will be
// croped to the given size.
var ResizeSizes = map[string]Size{
	"full":   {1600, 900},
	"big":    {700, 600},
	"medium": {500, 400},
	"small":  {300, 300},
	"thumb":  {100, 100},
	"square": {350, 290},
}

// AreaMaxSize is the size where an area should fit into.
var AreaMaxSize = Size{500, 200}

// ResizeFilter defines thee resize algorithm
var ResizeFilter = imaging.Lanczos

// ResizeType defines the method for resizing the image
type ResizeType int

// Types for resizing
const (
	ResizeFit ResizeType = iota
	ResizeThumbnail
)
