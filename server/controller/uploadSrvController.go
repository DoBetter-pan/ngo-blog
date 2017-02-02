/**
* @file uploadController.go
* @brief upload controller
* @author yingx
* @date 2017-02-02
 */

package controller

import (
	"net/http"
	"fmt"
	_ "errors"
    _ "strings"
    "strconv"
    _ "time"
	"os"
	"io"    
	_ "io/ioutil"
    _ "encoding/json"
    _ "math/rand"
	//model "ngo-blog/server/model"
	//session "ngo-blog/server/session"
)

func UploadError(w http.ResponseWriter, error string) {
	fmt.Fprintf(w, "{\"error\":\"%s\"}", error)
}

type UploadSrvController struct {
} 

func NewUploadController() *UploadSrvController {
	return &UploadSrvController{}
}

func (controller *UploadSrvController) Query(w http.ResponseWriter, r *http.Request) {
    res := "[]"

    SendBack(w, res)
}

func (controller *UploadSrvController) Get(w http.ResponseWriter, r *http.Request) {
	//UploadError(w, "Wrong method!!!")
    res := "{}"

    SendBack(w, res)
}

func (controller *UploadSrvController) New(w http.ResponseWriter, r *http.Request) {
    res := "{}"

    strId, _, ok := r.BasicAuth()
    if ok {
    	var userId int64 = 0
        userId, _ = strconv.ParseInt(strId, 10, 64)
        fmt.Println("=============>", userId)
		//parse the multipart form in the request
		err := r.ParseMultipartForm(100000)
		if err != nil {
			UploadError(w, "Failed to upload the file!!!")
			return
		}

		m := r.MultipartForm
		fmt.Println("=============>", m)
		path, _ := os.Getwd()
		files := m.File["xlsx-upload-file"]
		for i, _ := range files {
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				UploadError(w, "Can not open the file!!!")
				return
			}
			filename := path + "/public/upload/" + files[i].Filename
			fmt.Println("=============>", filename)
			dst, err := os.Create(path + "/public/upload/" + files[i].Filename)
			if err != nil {
				fmt.Println("=============>", filename, err)
				UploadError(w, "Can not create the file!!!")
				return
			}
			defer dst.Close()		
			//copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
				UploadError(w, "Can not copy the file!!!")
				return
			}
		}
		fmt.Fprint(w, "{}")        
    }

    SendBack(w, res)
}

func (controller *UploadSrvController) Update(w http.ResponseWriter, r *http.Request) {
    res := "{}"

    SendBack(w, res)
}

func (controller *UploadSrvController) Delete(w http.ResponseWriter, r *http.Request) {
    res := GetError(nil)

    SendBack(w, res)
}
