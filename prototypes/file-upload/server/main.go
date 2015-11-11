package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gographics/imagick/imagick"
	"github.com/gorilla/mux"
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
	"strings"
)

const (
	MAX_WIDTH  = 400
	MAX_HEIGHT = 200
)

type Post struct {
	Channelname *string  `json:"channelname"`
	Title       *string  `json:"title"`
	Fields      []*Field `json:"fields"`
}

type Field struct {
	Label     *string     `json:"label"`
	Type      *string     `json:"type"`
	Value     interface{} `json:"value"`
	Subfields []*Field    `json:"subfields"`
}

func saveImage(blob []byte) error {
	u4, err := uuid.NewV4()
	if err != nil {
		log.Print(err)
		return err
	}

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	err = mw.ReadImageBlob(blob)
	if err != nil {
		log.Print(err)
		return err
	}

	width := mw.GetImageWidth()
	height := mw.GetImageHeight()

	var newWidth, newHeight uint
	if width > MAX_WIDTH {
		newWidth = MAX_WIDTH
		newHeight = uint(float32(height) / float32(width/MAX_WIDTH))
	}
	if newHeight > MAX_HEIGHT {
		newHeight = MAX_HEIGHT
		newWidth = uint(float32(width) / float32(height/MAX_HEIGHT))
	}

	// Resize the image using the Lanczos filter
	// The blur factor is a float, where > 1 is blurry, < 1 is sharp
	err = mw.ResizeImage(newWidth, newHeight, imagick.FILTER_LANCZOS, 1)
	if err != nil {
		log.Print(err)
		return err
	}

	err = mw.SetImageCompressionQuality(70)
	if err != nil {
		log.Print(err)
		return err
	}

	err = mw.WriteImage("app/images/" + u4.String() + ".jpg")
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

// upload logic
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var post Post
	err := decoder.Decode(&post)
	if err != nil {
		log.Print(err)
		return
	}

	if post.Fields != nil {
		for _, field := range post.Fields {
			fmt.Println(*field.Type)
			if *field.Type == "Images" {
				images := field.Value.([]interface{})
				for _, image := range images {
					str := image.(string)
					str = strings.Split(str, ",")[1]
					blob, err := base64.StdEncoding.DecodeString(str)
					if err != nil {
						log.Println(err)
						return
					}
					err = saveImage(blob)
					if err != nil {
						log.Println(err)
						return
					}
				}
			}
		}
	}

	fmt.Println(post)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	imagick.Initialize()
	defer imagick.Terminate()

	r := mux.NewRouter()
	r.HandleFunc("/upload", uploadHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("app/")))

	http.Handle("/", r)
	log.Printf("serving http on :7183")
	err := http.ListenAndServe(":7183", nil)
	if err != nil {
		log.Fatal(err)
	}
}
