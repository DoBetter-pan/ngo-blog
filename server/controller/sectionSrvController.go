/**
* @file sectionSrvController.go
* @brief section service controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
	"net/http"
	"io/ioutil"
	model "go-angular/server/model"
)

type SectionSrvController struct {
}

func NewSectionSrvController() *SectionSrvController {
	return &SectionSrvController{}
}

func (controller *SectionSrvController) Query(w http.ResponseWriter, r *http.Request) {
    section := &model.SectionSrvModel{}
    res, err := section.FindAll()
    if err != nil {
        res = "[]"
    }

    SendBack(w, res)
}

func (controller *SectionSrvController) Get(w http.ResponseWriter, r *http.Request) {
    id := GetId(r)

    section := &model.SectionSrvModel{}
    res, err := section.Find(id)
    if err != nil {
        res = "{}"
    }

    SendBack(w, res)
}

func (controller *SectionSrvController) New(w http.ResponseWriter, r *http.Request) {
    res := "{}"

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        section := &model.SectionSrvModel{}
        res, err = section.Insert(string(data))
        if err != nil {
            res = "{}"
        }
    }

    SendBack(w, res)
}

func (controller *SectionSrvController) Update(w http.ResponseWriter, r *http.Request) {
    res := "{}"
    id := GetId(r)

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        section := &model.SectionSrvModel{}
        res, err = section.Update(id, string(data))
        if err != nil {
            res, err = section.Find(id)
            if err != nil {
                res = "{}"
            }
        }
    }

    SendBack(w, res)
}

func (controller *SectionSrvController) Delete(w http.ResponseWriter, r *http.Request) {
    id := GetId(r)

    section := &model.SectionSrvModel{}
    err := section.Delete(id)
    res := GetError(err)

    SendBack(w, res)
}

