package hugoexport

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/kupferstich/datatool/data"
)

// ImgArtwork exports all the images. PicRootFolder describes the source
// folder, where all the pictrues are edited. exportRootPath is the
// destination root folder.
func ImgArtwork(picRootFolder, exportRootPath string) {
	pics := data.LoadPictures(picRootFolder)
	for _, p := range pics {
		picPath := filepath.Join(
			picRootFolder,
			p.ID,
			ImgArtworkSrcFilename,
		)
		log.Println("Exporting ", p.ID)
		exportImage(picPath, p, exportRootPath)
		exportAreas(picPath, p, exportRootPath)
	}
}

func exportImage(picPath string, p data.Picture, exportRootPath string) {
	for key, size := range ResizeSizes {
		dstPath := filepath.Join(
			exportRootPath,
			ImgArtworkSubfolder,
			p.ID,
			fmt.Sprintf("%s_%s.jpg", p.ID, key),
		)
		img := openPic(picPath)
		resizePic(img, size, dstPath)
	}

}

func resizePic(img image.Image, size Size, dstPath string) {
	thumb := imaging.Fit(img, size.Width, size.Height, ResizeFilter)
	os.MkdirAll(filepath.Dir(dstPath), 0777)
	err := imaging.Save(thumb, dstPath)
	if err != nil {
		log.Println(err)
	}
}

func exportAreas(picPath string, p data.Picture, exportRootPath string) {
	scale := getScale(picPath, p)
	for i, area := range p.Areas {
		aid := fmt.Sprintf(
			"%02d_%s",
			i+1,
			strings.Replace(area.ID, " ", "_", -1),
		)
		dstPath := filepath.Join(
			exportRootPath,
			ImgArtworkSubfolder,
			p.ID,
			fmt.Sprintf("%s.jpg", aid),
		)
		imgArea := extractArea(picPath, area, scale)
		os.MkdirAll(filepath.Dir(dstPath), 0777)
		imaging.Save(imgArea, dstPath)
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
