/**
* @file menuSrvController.go
* @brief menu service controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
	"net/http"
	_ "fmt"
    _ "strings"
    _ "strconv"
	"io/ioutil"
	model "go-angular/server/model"
)

type MenuSrvController struct {
}

func NewMenuSrvController() *MenuSrvController {
	return &MenuSrvController{}
}

func (controller *MenuSrvController) Query(w http.ResponseWriter, r *http.Request) {
    menu := &model.MenuSrvModel{}
    res, err := menu.FindAll()
    if err != nil {
        res = "[]"
    }

    SendBack(w, res)
}

func (controller *MenuSrvController) Get(w http.ResponseWriter, r *http.Request) {
    id := GetId(r)

    menu := &model.MenuSrvModel{}
    res, err := menu.Find(id)
    if err != nil {
        res = "{}"
    }

    SendBack(w, res)
}

func (controller *MenuSrvController) New(w http.ResponseWriter, r *http.Request) {
    res := "{}"

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        menu := &model.MenuSrvModel{}
        res, err = menu.Insert(string(data))
        if err != nil {
            res = "{}"
        }
    }

    SendBack(w, res)
}

func (controller *MenuSrvController) Update(w http.ResponseWriter, r *http.Request) {
    res := "{}"
    id := GetId(r)

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        menu := &model.MenuSrvModel{}
        res, err = menu.Update(id, string(data))
        if err != nil {
            res, err = menu.Find(id)
            if err != nil {
                res = "{}"
            }
        }
    }

    SendBack(w, res)
}

func (controller *MenuSrvController) Delete(w http.ResponseWriter, r *http.Request) {
    id := GetId(r)

    menu := &model.MenuSrvModel{}
    err := menu.Delete(id)
    res := GetError(err)

    SendBack(w, res)
}

