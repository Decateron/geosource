package main

import (
	"net/http"
	"strconv"
	"html/template"
	"fmt"
	"time"
	"log"
	"crypto/md5"
	"io"
	"io/ioutil"
	"github.com/gorilla/mux"
	"github.com/nu7hatch/gouuid"
	"github.com/gographics/imagick/imagick"
)

const (
	MAX_WIDTH = 400
	MAX_HEIGHT = 200
)

//scales 
func scale(width, height uint) newWidth, newHeight uint {

}

// upload logic
func upload(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
        token := fmt.Sprintf("%x", h.Sum(nil))

        t, _ := template.ParseFiles("upload.gtpl")
        t.Execute(w, token)
    } else {
        r.ParseMultipartForm(32 << 20)
        file, handler, err := r.FormFile("uploadfile")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()

        fmt.Fprintf(w, "%v", handler.Header)

        u4, err := uuid.NewV4()
		if err != nil {
		    fmt.Println("error:", err)
		    return
		}

        // // f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
        // f, err := os.Create("./test/" + u4.String() + filepath.Ext(handler.Filename))
        // if err != nil {
        //     fmt.Println(err)
        //     return
        // }
        // defer f.Close()

        // io.Copy(f, file)

        blob, err := ioutil.ReadAll(file)
        if err != nil {
        	fmt.Println("error:", err)
        	return
        }

        mw := imagick.NewMagickWand()
		defer mw.Destroy()

		// err = mw.ReadImage("./test/" + u4.String() + filepath.Ext(handler.Filename))
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		err = mw.ReadImageBlob(blob)
		if err != nil {
			fmt.Println(err)
			return
		}

		width := mw.GetImageWidth()
		height := mw.GetImageHeight()

		var newWidth, newHeight uint
		if width > MAX_WIDTH {
			newWidth = MAX_WIDTH
			newHeight = uint(float32(height) / float32(width / MAX_WIDTH))
		}
		if newHeight > MAX_HEIGHT {
			newHeight = MAX_HEIGHT
			newWidth = uint(float32(width) / float32(height / MAX_HEIGHT))
		}

		// Resize the image using the Lanczos filter
		// The blur factor is a float, where > 1 is blurry, < 1 is sharp
		err = mw.ResizeImage(newWidth, newHeight, imagick.FILTER_LANCZOS, 1)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = mw.SetImageCompressionQuality(70)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = mw.WriteImage("app/images/" + u4.String() + ".jpg")
		if err != nil {
			fmt.Println(err)
			return
		}
    }
}

func main() {

	imagick.Initialize()
    defer imagick.Terminate()

    r := mux.NewRouter()
    r.HandleFunc("/upload", upload)
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("app/")))

    http.Handle("/", r)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
		log.Fatal(err)
	}
}