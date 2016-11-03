/**
* @file basicSrvController.go
* @brief basic service controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
	"net/http"
	"fmt"
    "strings"
    "strconv"
	"io/ioutil"
	model "go-angular/server/model"
)

type BasicSrvController struct {
}

func NewBasicSrvController() *BasicSrvController {
	return &BasicSrvController{}
}

func (controller *BasicSrvController) Query(w http.ResponseWriter, r *http.Request) {
    basic := &model.BasicSrvModel{}
    res, err := basic.FindAll()
    if err != nil {
        res = "[]"
    }

    SendBack(w, res)
}

func (controller *BasicSrvController) Get(w http.ResponseWriter, r *http.Request) {
    id := GetId(r)

    basic := &model.BasicSrvModel{}
    res, err := basic.Find(id)
    if err != nil {
        res = "{}"
    }

    SendBack(w, res)
}

func (controller *BasicSrvController) New(w http.ResponseWriter, r *http.Request) {
    res := "{}"

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        basic := &model.BasicSrvModel{}
        res, err = basic.Insert(string(data))
        if err != nil {
            res = "{}"
        }
    }

    SendBack(w, res)
}

func (controller *BasicSrvController) Update(w http.ResponseWriter, r *http.Request) {
    res := "{}"
    id := GetId(r)

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        basic := &model.BasicSrvModel{}
        res, err = basic.Update(id, string(data))
        if err != nil {
            res, err = basic.Find(id)
            if err != nil {
                res = "{}"
            }
        }
    }

    SendBack(w, res)
}

func (controller *BasicSrvController) Delete(w http.ResponseWriter, r *http.Request) {
    id := GetId(r)

    basic := &model.BasicSrvModel{}
    err := basic.Delete(id)
    res := GetError(err)

    SendBack(w, res)
}

