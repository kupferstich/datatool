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

// JSONPostSubfolder is where the post json files will be copied
var JSONPostSubfolder = filepath.Join("data", "post")

// ImgArtworkSubfolder is where the images will be exported
var ImgArtworkSubfolder = filepath.Join("static", "img", "artwork")

// ImgArtistSubfolder is where the images will be exported
var ImgArtistSubfolder = filepath.Join("static", "img", "artist")

// ImgArtistSubfolder is where the images will be exported
var ImgPostSubfolder = filepath.Join("static", "img", "post")

// ContentArtworkSubfolder the folder for the md content files.
var ContentArtworkSubfolder = filepath.Join("content", "artwork")

// ContentArtistSubfolder the folder for the md content files.
var ContentArtistSubfolder = filepath.Join("content", "artist")

// ContentPostSubfolder the folder for the md content files.
var ContentPostSubfolder = filepath.Join("content", "post")

// ImgArtworkSrcFilename is the filename of the picture inside the
// src (pictures) folder
var ImgArtworkSrcFilename = "00000001.jpg"

// ImgAlwaysResize is an option that defines if an existing file is changed
// or not. If the value is true all pcitures will be resized. That makes the
// export much longer.
var ImgAlwaysResize = true

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

// PostPicSize defines the size for the pictures embeded inside a
// blog post.
var PostPicSize = Size{350, 350}

// PostPicResizeType defines the ResizeType for all the blog pics.
var PostPicResizeType = ResizeThumbnail
