/**
* @file categorySrvController.go
* @brief category service controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
	"net/http"
	"io/ioutil"
	model "go-angular/server/model"
)

type CategorySrvController struct {
}

func NewCategorySrvController() *CategorySrvController {
	return &CategorySrvController{}
}

func (controller *CategorySrvController) Query(w http.ResponseWriter, r *http.Request) {
    category := &model.CategorySrvModel{}
    res, err := category.FindAll()
    if err != nil {
        res = "[]"
    }

    SendBack(w, res)
}

func (controller *CategorySrvController) Get(w http.ResponseWriter, r *http.Request) {
    id := GetId(r)

    category := &model.CategorySrvModel{}
    res, err := category.Find(id)
    if err != nil {
        res = "{}"
    }

    SendBack(w, res)
}

func (controller *CategorySrvController) New(w http.ResponseWriter, r *http.Request) {
    res := "{}"

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        category := &model.CategorySrvModel{}
        res, err = category.Insert(string(data))
        if err != nil {
            res = "{}"
        }
    }

    SendBack(w, res)
}

func (controller *CategorySrvController) Update(w http.ResponseWriter, r *http.Request) {
    res := "{}"
    id := GetId(r)

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        category := &model.CategorySrvModel{}
        res, err = category.Update(id, string(data))
        if err != nil {
            res, err = category.Find(id)
            if err != nil {
                res = "{}"
            }
        }
    }

    SendBack(w, res)
}

func (controller *CategorySrvController) Delete(w http.ResponseWriter, r *http.Request) {
    id := GetId(r)

    category := &model.CategorySrvModel{}
    err := category.Delete(id)
    res := GetError(err)

    SendBack(w, res)
}

