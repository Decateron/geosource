package fields

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/gographics/imagick/imagick"
	"github.com/pborman/uuid"
)

const (
	MAX_WIDTH     = 2000
	MAX_HEIGHT    = 2000
	IMAGE_QUALITY = 70
	IMAGE_TYPE    = ".jpg"

	APP_DIR   = "../../app/"
	IMAGE_DIR = "images/"
)

type ImagesForm struct{}

func (imagesForm *ImagesForm) Validate() error {
	return nil
}

func (imagesForm *ImagesForm) ValidateValue(value Value) error {
	imagesValue, ok := value.(*ImagesValue)
	if !ok {
		return errors.New("Type mismatch.")
	}
	if imagesValue == nil {
		return nil
	}
	for i, base64str := range *imagesValue {
		filename, err := SaveImage(base64str)
		if err != nil {
			// TODO: Delete other images
			return err
		}
		(*imagesValue)[i] = filename
	}
	return nil
}

func (imagesForm *ImagesForm) UnmarshalValue(blob []byte) (Value, error) {
	if len(blob) <= 0 {
		return nil, nil
	}
	var value ImagesValue
	err := json.Unmarshal(blob, &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

type ImagesValue []string

func (imagesValue *ImagesValue) IsComplete() bool {
	return imagesValue != nil && len(*imagesValue) > 0
}

// Converts the base64 string into an image, saves it to the file system, and
// returns its filename. If it is unsuccessful, an error is returned.
func SaveImage(base64str string) (string, error) {
	// removes header information
	strs := strings.Split(base64str, ",")
	str := strs[len(strs)-1]

	blob, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	filename := uuid.NewUUID().String() + IMAGE_TYPE

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

	file, err := os.Create(APP_DIR + IMAGE_DIR + filename)
	if err != nil {
		return "", err
	}

	err = mw.WriteImageFile(file)
	if err != nil {
		return "", err
	}

	return filename, nil
}
