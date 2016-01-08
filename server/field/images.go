package field

import (
	"encoding/base64"
	"github.com/gographics/imagick/imagick"
	"github.com/nu7hatch/gouuid"
	"log"
	"os"
	"strings"
)

const (
	MAX_WIDTH     = 2000
	MAX_HEIGHT    = 2000
	IMAGE_QUALITY = 70
	IMAGE_TYPE    = ".jpg"
	IMAGE_DIR     = "../app/images/"
)

type Images []Image

func (images *Images) Validate() error {
	if images != nil {
		for i, image := range *images {
			filename, err := image.Save()
			if err != nil {
				log.Println(err)
				return err
			}
			[]Image(*images)[i] = Image(filename)
		}
	}
	return nil
}

func (images *Images) IsEmpty() bool {
	return images == nil || len(*images) == 0
}

type Image string

// Saves the base64 string to an image, and returns its filename.
func (image Image) Save() (string, error) {
	// removes header information
	str := strings.Split(string(image), ",")[1]
	blob, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	u4, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	err = mw.ReadImageBlob(blob)
	if err != nil {
		return "", err
	}

	width := mw.GetImageWidth()
	newWidth := width
	height := mw.GetImageHeight()
	newHeight := height

	if width > MAX_WIDTH {
		newWidth = MAX_WIDTH
		newHeight = uint(float64(height) / float64(width/MAX_WIDTH))
	}
	if newHeight > MAX_HEIGHT {
		newHeight = MAX_HEIGHT
		newWidth = uint(float64(width) / float64(height/MAX_HEIGHT))
	}

	// Resize the image using the Lanczos filter
	// The blur factor is a float, where > 1 is blurry, < 1 is sharp
	err = mw.ResizeImage(newWidth, newHeight, imagick.FILTER_LANCZOS, 1)
	if err != nil {
		return "", err
	}

	err = mw.SetImageCompressionQuality(IMAGE_QUALITY)
	if err != nil {
		return "", err
	}

	file, err := os.Create(IMAGE_DIR + u4.String() + IMAGE_TYPE)
	if err != nil {
		return "", err
	}

	err = mw.WriteImageFile(file)
	if err != nil {
		return "", err
	}

	return u4.String(), nil
}
