/**
* @file loginSrvController.go
* @brief login service controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
	"net/http"
	"fmt"
	"errors"
    _ "strings"
    _ "strconv"
	"io/ioutil"
    "encoding/json"
	model "go-angular/server/model"
	session "go-angular/server/session"
)

type Login struct {
     Id int64 `json:"id"`
     Name  string `json:"name"`
     Password string `json:"password"`
     Stored int64 `json:"stored"`
}

type LoginSrvController struct {
}

func NewLoginSrvController() *LoginSrvController {
	return &LoginSrvController{}
}

func (controller *LoginSrvController) Query(w http.ResponseWriter, r *http.Request) {
    /*
    login := &model.LoginSrvModel{}
    res, err := login.FindAll()
    if err != nil {
        res = "[]"
    }
    */

    fmt.Println("==========>Query!!!")
    res := "[]"
    SendBack(w, res)
}

func (controller *LoginSrvController) Get(w http.ResponseWriter, r *http.Request) {
    /*
    id := GetId(r)

    login := &model.LoginSrvModel{}
    res, err := login.Find(id)
    if err != nil {
        res = "{}"
    }
    */

    fmt.Println("==========>Get!!!")
    res := "{}"
    SendBack(w, res)
}

func (controller *LoginSrvController) New(w http.ResponseWriter, r *http.Request) {
    res := "{}"

    /*
    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        login := &model.LoginSrvModel{}
        res, err = login.Insert(string(data))
        if err != nil {
            res = "{}"
        }
    }
    */

    fmt.Println("==========>New!!!")
    SendBack(w, res)
}

func (controller *LoginSrvController) Update(w http.ResponseWriter, r *http.Request) {
    res := "{}"
    /*
    id := GetId(r)

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        login := &model.LoginSrvModel{}
        res, err = login.Update(id, string(data))
        if err != nil {
            res, err = login.Find(id)
            if err != nil {
                res = "{}"
            }
        }
    }
    */

    fmt.Println("==========>Update!!!")
    SendBack(w, res)
}

func (controller *LoginSrvController) Delete(w http.ResponseWriter, r *http.Request) {
    /*
    id := GetId(r)

    login := &model.LoginSrvModel{}
    err := login.Delete(id)
    res := GetError(err)
    */

    fmt.Println("==========>Delete!!!")
    res := GetError(nil)
    SendBack(w, res)
}

func (controller *LoginSrvController) Getuser(w http.ResponseWriter, r *http.Request) {
    var login Login

    login.Stored = 0
    validated, id, name, role, nonce := session.ValidateSessionByCookie(r)
    if validated {
        loginModel := &model.LoginSrvModel{}
        dataLogin, err := loginModel.FindObject(id)
        if err == nil && name == dataLogin.Name && role == dataLogin.Role && nonce == dataLogin.Nonce {
            login.Id = id
            login.Name = name
            login.Password = dataLogin.Password
            login.Stored = 1
        }
    }

    if login.Stored == 0 {
        login.Id = -1
        login.Name = ""
        login.Password = ""
    }

    res := "{}"
    data, err := json.Marshal(login)
    if err != nil {
        res = "{}"
    } else {
        res = string(data)
    }

    SendBack(w, res)
}

func (controller *LoginSrvController) Checkuser(w http.ResponseWriter, r *http.Request) {
    var login Login
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        err = json.Unmarshal([]byte(data), &login)
        if err == nil {
            loginModel := &model.LoginSrvModel{}
            dataLogin, err2 := loginModel.FindObjectByName(login.Name)
            err = err2
            if err == nil {
                if login.Name != dataLogin.Name || login.Password != dataLogin.Password {
                    err = errors.New("wrong name or password!!!")
                } else {
                    //store cookie
                    if login.Stored == 1 {
                        session.WriteBackSessionCookie(w, dataLogin.Id, login.Name, dataLogin.Role, dataLogin.Nonce, "/", 7200)
                    } else {
                        session.WriteBackSessionCookie(w, dataLogin.Id, login.Name, dataLogin.Role, dataLogin.Nonce, "/", 0)
                    }
                }
            }
        }
    }

    res := GetError(err)
    SendBack(w, res)
}

