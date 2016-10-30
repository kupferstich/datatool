package hugoexport

import (
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/disintegration/imaging"
	"github.com/kupferstich/datatool/data"
)

// ImgArtwork exports all the images. PicRootFolder describes the source
// folder, where all the pictrues are edited. exportRootPath is the
// destination root folder.
func ImgArtwork(picRootFolder, postsRootFolder, exportRootPath string) {
	pics := data.LoadPictures(picRootFolder)
	posts := data.NewPosts(postsRootFolder)
	posts.Load()
	sort.Sort(data.ByPostDate(posts.Posts))
	for _, p := range pics {
		if p.Status != "fertig" && ImgExportDraft == false {
			continue
		}
		picPath := filepath.Join(
			picRootFolder,
			p.ID,
			ImgArtworkSrcFilename,
		)
		log.Println("Exporting ", p.ID)
		ExportImage(picPath, p, exportRootPath)
		ExportImageData(p, exportRootPath)
		ExportImageContent(&p, posts, exportRootPath)
		ExportAreas(picPath, p, exportRootPath)
	}
}

// ExportImage resizes the image into the ResizeSizes and saves them
// into the exportRootPath.
func ExportImage(picPath string, p data.Picture, exportRootPath string) {

	for key, size := range ResizeSizes {
		dstPath := filepath.Join(
			exportRootPath,
			ImgArtworkSubfolder,
			p.ID,
			fmt.Sprintf("%s_%s.jpg", p.ID, key),
		)
		if skipImageCreation(dstPath) {
			// Aboart when dst File already exists
			continue
		}
		// Just openPic when needed
		img := openPic(picPath)
		rType := ResizeFit
		/*if key == "thumb" || key == "square" {
			rType = ResizeThumbnail
		}*/
		resizePic(img, size, dstPath, rType)
	}

}

// ExportImageData exports the pic data as JSON into the JSONArtworkSubfolder
func ExportImageData(p data.Picture, exportRootPath string) {
	// Export the data
	dataPath := filepath.Join(
		exportRootPath,
		JSONArtworkSubfolder,
		fmt.Sprintf("%s.json", p.ID),
	)
	os.MkdirAll(filepath.Dir(dataPath), 0777)
	b, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile(dataPath, b, 0777)
	if err != nil {
		log.Println(err)
	}
}

// ExportImageContent exports the picture data into the content folder
func ExportImageContent(p *data.Picture, posts *data.Posts, exportRootPath string) {
	dstPath := filepath.Join(
		exportRootPath,
		ContentArtworkSubfolder,
		fmt.Sprintf("%s.md", p.ID),
	)
	os.MkdirAll(filepath.Dir(dstPath), 0777)
	f, err := os.Create(dstPath)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	ContentFromPicture(p, posts, f)
}

// ExportAreas the defined areas are cut out of the original image.
func ExportAreas(picPath string, p data.Picture, exportRootPath string) {
	scale := getScale(picPath, p)
	for i, area := range p.Areas {
		aid := fmt.Sprintf(
			"area_%d",
			i,
		)
		dstPath := filepath.Join(
			exportRootPath,
			ImgArtworkSubfolder,
			p.ID,
			fmt.Sprintf("%s.jpg", aid),
		)
		if skipImageCreation(dstPath) {
			continue
		}
		imgArea := extractArea(picPath, area, scale)
		resizePic(imgArea, AreaMaxSize, dstPath, ResizeFit)
	}
}

func getScale(picPath string, p data.Picture) float32 {
	imgConf := getPicConfig(picPath)
	scaleX := float32(imgConf.Width) / float32(p.CanvasWidth)
	scaleY := float32(imgConf.Height) / float32(p.CanvasHeight)
	scale := scaleY
	if scaleX > scaleY {
		scale = scaleX
	}
	return scale
}

func getPicConfig(picturePath string) image.Config {
	file, err := os.Open(filepath.Join(picturePath))
	defer file.Close()
	if err != nil {
		log.Println(err)
	}
	imgc, err := jpeg.DecodeConfig(file)
	if err != nil {
		log.Println(err)
	}
	return imgc
}

func extractArea(picPath string, area data.Area, scale float32) *image.NRGBA {
	//aid := strings.Replace(area.ID, " ", "_", -1)
	img := openPic(picPath)
	imgArea := imaging.Crop(img, area.ImageRect(scale))
	return imgArea
}

// resizePic is for resizing an Image.
// Just as const defined ResizeTypes are allowed to use for resizing:
// - ResizeFit
// - ResizeThumbnail
func resizePic(img image.Image, size Size, dstPath string, resizeType ResizeType) {
	var resizedImg *image.NRGBA
	switch resizeType {
	case ResizeThumbnail:
		resizedImg = imaging.Thumbnail(img, size.Width, size.Height, ResizeFilter)
	case ResizeFit:
		resizedImg = imaging.Fit(img, size.Width, size.Height, ResizeFilter)
	default:
		log.Println("resizePic got not supporte resizeType")
	}

	os.MkdirAll(filepath.Dir(dstPath), 0777)
	err := imaging.Save(resizedImg, dstPath)
	if err != nil {
		log.Println(err)
	}
}

// skipImageCreation returns true, if the image should be created or not.
// If the var ImgAlwaysResize is true that function always returns false.
// No file is skipped then.
// If not it is checked, if the file exists. An existing file will be skipped.
func skipImageCreation(fpath string) bool {
	if ImgAlwaysResize {
		return false
	}
	if _, err := os.Stat(fpath); err == nil {
		return true
	}
	return false
}

func openPic(picturePath string) image.Image {
	file, err := os.Open(filepath.Join(picturePath))
	defer file.Close()
	if err != nil {
		log.Println(err)
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Println(err)
	}
	return img
}
