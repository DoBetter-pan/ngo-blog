/**
* @file counterSrvController.go
* @brief counter service controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
	"net/http"
	"fmt"
	_ "errors"
    "strings"
    _ "strconv"
	_ "io/ioutil"
    "encoding/json"
	model "go-angular/server/model"
	//session "go-angular/server/session"
)

type Counter struct {
     VisitorCounter int64 `json:"visitorCounter"`
}

type CounterSrvController struct {
}

func NewCounterSrvController() *CounterSrvController {
	return &CounterSrvController{}
}

func (controller *CounterSrvController) Query(w http.ResponseWriter, r *http.Request) {
    /*
    counter := &model.CounterSrvModel{}
    res, err := counter.FindAll()
    if err != nil {
        res = "[]"
    }
    */

    fmt.Println("==========>Query!!!")
    res := "[]"
    SendBack(w, res)
}

func (controller *CounterSrvController) Get(w http.ResponseWriter, r *http.Request) {
    /*
    id := GetId(r)

    counter := &model.CounterSrvModel{}
    res, err := counter.Find(id)
    if err != nil {
        res = "{}"
    }
    */

    fmt.Println("==========>Get!!!")
    res := "{}"
    SendBack(w, res)
}

func (controller *CounterSrvController) New(w http.ResponseWriter, r *http.Request) {
    res := "{}"

    /*
    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        counter := &model.CounterSrvModel{}
        res, err = counter.Insert(string(data))
        if err != nil {
            res = "{}"
        }
    }
    */

    fmt.Println("==========>New!!!")
    SendBack(w, res)
}

func (controller *CounterSrvController) Update(w http.ResponseWriter, r *http.Request) {
    res := "{}"
    /*
    id := GetId(r)

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        counter := &model.CounterSrvModel{}
        res, err = counter.Update(id, string(data))
        if err != nil {
            res, err = counter.Find(id)
            if err != nil {
                res = "{}"
            }
        }
    }
    */

    fmt.Println("==========>Update!!!")
    SendBack(w, res)
}

func (controller *CounterSrvController) Delete(w http.ResponseWriter, r *http.Request) {
    /*
    id := GetId(r)

    counter := &model.CounterSrvModel{}
    err := counter.Delete(id)
    res := GetError(err)
    */

    fmt.Println("==========>Delete!!!")
    res := GetError(nil)
    SendBack(w, res)
}

func (controller *CounterSrvController) Visitor(w http.ResponseWriter, r *http.Request) {    
    res := "{}"
    var counter Counter
    _, name, ok := r.BasicAuth()
    if !ok {
        name = "guest"
    }
    //r.Header.Get("X-Forwarded-For")
    addr := strings.Split(r.RemoteAddr, ":") 
    ip := addr[0]
    url := ""
    counterModel := &model.CounterSrvModel{}
    counterModel.Insert(name, ip, url)
    counter.VisitorCounter = counterModel.FindCounter(url)

    data, err := json.Marshal(counter)
    if err != nil {
        res = "{}"
    } else {
        res = string(data)
    }

    SendBack(w, res)
}


