package fields

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"github.com/pborman/uuid"
	"gopkg.in/gographics/imagick.v1/imagick"
)

// These variables are not set to constant so that they could be modified
// programatically at some point if needed

var ImageMaxWidth uint = 2000
var ImageMaxHeight uint = 2000
var ImageQuality uint = 70
var ImageType string = ".jpg"

var ThumbnailMaxWidth uint = 100
var ThumbnailMaxHeight uint = 100
var ThumbnailQuality uint = 70
var ThumbnailType string = ".jpg"

var AppDir string = "app/"
var MediaDir string = "media/"
var ImageDir string = "images/"
var ThumbnailDir string = "thumbnails/"

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

// This function assumes that the images have already been saved to files.
func (imagesValue *ImagesValue) GenerateThumbnail() (string, error) {
	if !imagesValue.IsComplete() {
		return "", errors.New("No images to generate thumbnail for.")
	}
	filename := MediaDir + ThumbnailDir + uuid.NewUUID().String() + ThumbnailType
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	err := mw.ReadImage(AppDir + (*imagesValue)[0])
	if err != nil {
		return "", err
	}

	err = mw.ThumbnailImage(ThumbnailMaxWidth, ThumbnailMaxHeight)
	if err != nil {
		return "", err
	}

	err = mw.SetImageCompressionQuality(ThumbnailQuality)
	if err != nil {
		return "", err
	}

	err = mw.WriteImage(AppDir + filename)
	if err != nil {
		return "", err
	}

	return filename, nil
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
	filename := MediaDir + ImageDir + uuid.NewUUID().String() + ImageType

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

	if width > ImageMaxWidth {
		newWidth = ImageMaxWidth
		newHeight = uint(float64(height) / float64(width/ImageMaxWidth))
	}
	if newHeight > ImageMaxHeight {
		newHeight = ImageMaxHeight
		newWidth = uint(float64(width) / float64(height/ImageMaxHeight))
	}

	// Resize the image using the Lanczos filter
	// The blur factor is a float, where > 1 is blurry, < 1 is sharp
	err = mw.ResizeImage(newWidth, newHeight, imagick.FILTER_LANCZOS, 1)
	if err != nil {
		return "", err
	}

	err = mw.SetImageCompressionQuality(ImageQuality)
	if err != nil {
		return "", err
	}

	err = mw.WriteImage(AppDir + filename)
	if err != nil {
		return "", err
	}

	return filename, nil
}
