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
    "time"
	"os"
	"io"
	_ "io/ioutil"
    "encoding/json"
    "math/rand"
	//model "ngo-blog/server/model"
	//session "ngo-blog/server/session"
)

type UploadStatus struct {
    Success bool `json:"success"`
    Msg string `json:"msg"`
    FilePath string `json:"file_path"`
}

func SendBackUploadStatus(w http.ResponseWriter, success bool, msg, filepath string) {
    status := &UploadStatus{success, msg, filepath}
    data, err := json.Marshal(status)
    if err != nil {
        data = []byte(`{"success": false, "msg": "unknown error", "file_path":""}`)
    }
	fmt.Fprintf(w, string(data))
}

type UploadSrvController struct {
}

func NewUploadSrvController() *UploadSrvController {
	return &UploadSrvController{}
}

func (controller *UploadSrvController) Query(w http.ResponseWriter, r *http.Request) {
    res := "[]"

    SendBack(w, res)
}

func (controller *UploadSrvController) Get(w http.ResponseWriter, r *http.Request) {
    res := "{}"

    SendBack(w, res)
}

func (controller *UploadSrvController) New(w http.ResponseWriter, r *http.Request) {
    success := false
    msg := "unknow error"
    filepath := ""

    strId, _, ok := r.BasicAuth()
    if ok {
	    var userId int64 = 0
        userId, _ = strconv.ParseInt(strId, 10, 64)
        t := time.Now()
		basepath, _ := os.Getwd()
        timepath := t.Format("2006/01/02/")
        path := fmt.Sprintf("%s/public/upload/%s", basepath, timepath)
        err := os.MkdirAll(path, os.ModePerm)
        if err == nil {
            //parse the multipart form in the request
            err = r.ParseMultipartForm(100000)
            if err == nil {
                m := r.MultipartForm
                success = true
                files := m.File["upload_file"]
                for i, _ := range files {
                    file, err := files[i].Open()
                    if err != nil {
                        success = false
                        msg = "Can't open the file!"
                        break
                    }
                    defer file.Close()
                    rand.Seed(t.UnixNano())
                    r1 := rand.Intn(10000)
                    r2 := rand.Intn(10000)
                    f := t.Format("150405")
                    filename := fmt.Sprintf("%s%s%d%d%d", path, f, userId, r1, r2)
                    filepath = fmt.Sprintf("/upload/%s%s%d%d%d", timepath, f, userId, r1, r2)
                    dst, err := os.Create(filename)
                    if err != nil {
                        success = false
                        msg = "Can't create the file!"
                        break
                    }
                    defer dst.Close()
                    //copy the uploaded file to the destination file
                    if _, err := io.Copy(dst, file); err != nil {
                        success = false
                        msg = "Can't copy the file!"
                        break
                    }
                }
                if success {
                    msg = "OK"
                }
            } else {
                msg = "Failed to upload the file!"
            }
        } else {
            msg = "Not Permitted!"
        }
    } else {
        msg = "Not Authorized!"
    }

    SendBackUploadStatus(w, success, msg, filepath)
}

func (controller *UploadSrvController) Update(w http.ResponseWriter, r *http.Request) {
    res := "{}"

    SendBack(w, res)
}

func (controller *UploadSrvController) Delete(w http.ResponseWriter, r *http.Request) {
    res := GetError(nil)

    SendBack(w, res)
}
