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
		imgConf := getPicConfig(picPath)
		scaleX := float32(imgConf.Width) / float32(p.CanvasWidth)
		scaleY := float32(imgConf.Height) / float32(p.CanvasHeight)
		scale := scaleY
		if scaleX > scaleY {
			scale = scaleX
		}
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
			log.Println("Writing to ", dstPath)
			imgArea := ImgExtractArea(picPath, area, scale)
			os.MkdirAll(filepath.Dir(dstPath), 0777)
			imaging.Save(imgArea, dstPath)
		}
	}
}

func ImgExtractArea(picPath string, area data.Area, scale float32) *image.NRGBA {
	//aid := strings.Replace(area.ID, " ", "_", -1)
	img := openPic(picPath)
	imgArea := imaging.Crop(img, area.ImageRect(scale))
	return imgArea
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