/**
* @file blogSrvController.go
* @brief blog service controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
    _ "fmt"
    "strconv"
	"net/http"
	"io/ioutil"
	model "go-angular/server/model"
)

type BlogSrvController struct {
}

func NewBlogSrvController() *BlogSrvController {
	return &BlogSrvController{}
}

func (controller *BlogSrvController) CheckRules() map[string] []string {
    rules := make(map[string] []string, 5)
    rules["New"] = []string {
        "allow admin,writer",
        "deny *" }
    rules["Update"] = []string {
        "allow admin,writer",
        "deny *" }
    rules["Delete"] = []string {
        "allow admin,writer",
        "deny *" }                
    return rules
}

func (controller *BlogSrvController) Query(w http.ResponseWriter, r *http.Request) {
    res := "[]"
    err := r.ParseForm()
    if err == nil {
        isFirstPage := true
        k := ""
        v := int64(0)
        p := int64(0)

        va, ok := r.Form["s"]
        if ok {
            k = "s"
            v, err = strconv.ParseInt(va[0], 10, 64)
            isFirstPage = false
        }
        va, ok = r.Form["c"]
        if ok {
            k = "c"
            v, err = strconv.ParseInt(va[0], 10, 64)
            isFirstPage = false
        }
        va, ok = r.Form["p"]
        if ok && err == nil {
            p, err = strconv.ParseInt(va[0], 10, 64)
        }
        if err == nil {
            blog := &model.BlogSrvModel{}
            if isFirstPage {
                res, err = blog.FindAllByCount(3)
            } else {
                res, err = blog.FindAllByKeyValue(k, v, p)
            }
            if err != nil {
                res = "[]"
            }
        }
    }

    SendBack(w, res)
}

func (controller *BlogSrvController) Get(w http.ResponseWriter, r *http.Request) {
    id := GetId(r)

    blog := &model.BlogSrvModel{}
    res, err := blog.Find(id)
    if err != nil {
        res = "{}"
    }

    SendBack(w, res)
}

func (controller *BlogSrvController) New(w http.ResponseWriter, r *http.Request) {
    res := "{}"

    //r.ParseForm()
    id, _, ok := r.BasicAuth()
    if !ok {
        id = "0"
    }
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        blog := &model.BlogSrvModel{}
        res, err = blog.Insert(id, string(data))
        if err != nil {
            res = "{}"
        }
    }

    SendBack(w, res)
}

func (controller *BlogSrvController) Update(w http.ResponseWriter, r *http.Request) {
    res := "{}"
    id := GetId(r)

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        blog := &model.BlogSrvModel{}
        res, err = blog.Update(id, string(data))
        if err != nil {
            res, err = blog.Find(id)
            if err != nil {
                res = "{}"
            }
        }
    }

    SendBack(w, res)
}

func (controller *BlogSrvController) Delete(w http.ResponseWriter, r *http.Request) {
    id := GetId(r)

    blog := &model.BlogSrvModel{}
    err := blog.Delete(id)
    res := GetError(err)

    SendBack(w, res)
}

