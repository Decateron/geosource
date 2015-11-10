package main

import (
	"crypto/md5"
	"fmt"
	"github.com/gographics/imagick/imagick"
	"github.com/gorilla/mux"
	"github.com/nu7hatch/gouuid"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
)

const (
	MAX_WIDTH  = 400
	MAX_HEIGHT = 200
)

func saveImage(file multipart.File) error {
	u4, err := uuid.NewV4()
	if err != nil {
		log.Print(log.Llongfile, err)
		return err
	}

	blob, err := ioutil.ReadAll(file)
	if err != nil {
		log.Print(log.Llongfile, err)
		return err
	}

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	err = mw.ReadImageBlob(blob)
	if err != nil {
		log.Print(log.Llongfile, err)
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
		log.Print(log.Llongfile, err)
		return err
	}

	err = mw.SetImageCompressionQuality(70)
	if err != nil {
		log.Print(log.Llongfile, err)
		return err
	}

	err = mw.WriteImage("app/images/" + u4.String() + ".jpg")
	if err != nil {
		log.Print(log.Llongfile, err)
		return err
	}

	return nil
}

// upload logic
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("uploadHandler.gtpl")
		t.Execute(w, token)
	} else {
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			log.Print(log.Llongfile, err)
			return
		}

		fmt.Println(r.FormValue("field-1"))

		formdata := r.MultipartForm // ok, no problem so far, read the Form data

		//get the *fileheaders
		files := formdata.File["field-2"] // grab the filenames
		for i, _ := range files {         // loop through the files one by one
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				log.Print(log.Llongfile, err)
				return
			}

			err = saveImage(file)
			if err != nil {
				log.Print(log.Llongfile, err)
				return
			}
		}

		checkboxes := formdata.Value["field-3"]
		for i, _ := range checkboxes {
			checkbox := checkboxes[i]
			fmt.Println(checkbox)
		}

		fmt.Println(r.FormValue("field-4"))
	}
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
