/**
* @file linkSrvController.go
* @brief link service controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
	"net/http"
	"io/ioutil"
	model "go-angular/server/model"
)

type LinkSrvController struct {
}

func NewLinkSrvController() *LinkSrvController {
	return &LinkSrvController{}
}

func (controller *LinkSrvController) Query(w http.ResponseWriter, r *http.Request) {
    link := &model.LinkSrvModel{}
    res, err := link.FindAll()
    if err != nil {
        res = "[]"
    }

    SendBack(w, res)
}

func (controller *LinkSrvController) Get(w http.ResponseWriter, r *http.Request) {
    id := GetId(r)

    link := &model.LinkSrvModel{}
    res, err := link.Find(id)
    if err != nil {
        res = "{}"
    }

    SendBack(w, res)
}

func (controller *LinkSrvController) New(w http.ResponseWriter, r *http.Request) {
    res := "{}"

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        link := &model.LinkSrvModel{}
        res, err = link.Insert(string(data))
        if err != nil {
            res = "{}"
        }
    }

    SendBack(w, res)
}

func (controller *LinkSrvController) Update(w http.ResponseWriter, r *http.Request) {
    res := "{}"
    id := GetId(r)

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        link := &model.LinkSrvModel{}
        res, err = link.Update(id, string(data))
        if err != nil {
            res, err = link.Find(id)
            if err != nil {
                res = "{}"
            }
        }
    }

    SendBack(w, res)
}

func (controller *LinkSrvController) Delete(w http.ResponseWriter, r *http.Request) {
    id := GetId(r)

    link := &model.LinkSrvModel{}
    err := link.Delete(id)
    res := GetError(err)

    SendBack(w, res)
}

