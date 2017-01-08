/**
* @file postlikeSrvController.go
* @brief postlike service controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
	"net/http"
	"fmt"
	_ "errors"
    "strings"
    "strconv"
    "time"
	_ "io/ioutil"
    "encoding/json"
    "math/rand"
	model "ngo-blog/server/model"
	//session "ngo-blog/server/session"
)

type Postlike struct {
    Message string `json:"message"`
    LikeCount int64 `json:"likecount"`
    UnlikeCount int64 `json:"unlikecount"`     
}

type PostlikeSrvController struct {
}

func NewPostlikeSrvController() *PostlikeSrvController {
	return &PostlikeSrvController{}
}

func (controller *PostlikeSrvController) Query(w http.ResponseWriter, r *http.Request) {
    /*
    postlike := &model.PostlikeSrvModel{}
    res, err := postlike.FindAll()
    if err != nil {
        res = "[]"
    }
    */

    res := "[]"
    SendBack(w, res)
}

func (controller *PostlikeSrvController) Get(w http.ResponseWriter, r *http.Request) {
    /*
    id := GetId(r)

    postlike := &model.PostlikeSrvModel{}
    res, err := postlike.Find(id)
    if err != nil {
        res = "{}"
    }
    */

    res := "{}"
    SendBack(w, res)
}

func (controller *PostlikeSrvController) New(w http.ResponseWriter, r *http.Request) {
    res := "{}"

    /*
    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        postlike := &model.PostlikeSrvModel{}
        res, err = postlike.Insert(string(data))
        if err != nil {
            res = "{}"
        }
    }
    */

    SendBack(w, res)
}

func (controller *PostlikeSrvController) Update(w http.ResponseWriter, r *http.Request) {
    res := "{}"
    /*
    id := GetId(r)

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        postlike := &model.PostlikeSrvModel{}
        res, err = postlike.Update(id, string(data))
        if err != nil {
            res, err = postlike.Find(id)
            if err != nil {
                res = "{}"
            }
        }
    }
    */

    SendBack(w, res)
}

func (controller *PostlikeSrvController) Delete(w http.ResponseWriter, r *http.Request) {
    /*
    id := GetId(r)

    postlike := &model.PostlikeSrvModel{}
    err := postlike.Delete(id)
    res := GetError(err)
    */

    res := GetError(nil)
    SendBack(w, res)
}

func (controller *PostlikeSrvController) Postlike(w http.ResponseWriter, r *http.Request) {
    method := r.Method
    if method == "POST" {
        res := `{"message":"Wrong status", "likecount":0, "unlikecount":0 }`
        var postlike Postlike
        cookieStr := ""
        err := r.ParseForm()
        if err == nil {
            va, ok := r.Form["pid"]
            if ok {
                postId, err := strconv.ParseInt(va[0], 10, 64)
                if err == nil {
                    var userId int64 = 0
                    strId, _, ok := r.BasicAuth()
                    if ok {
                        userId, _ = strconv.ParseInt(strId, 10, 64)
                    }
                    addr := strings.Split(r.RemoteAddr, ":") 
                    ip := addr[0]
                    cookiePostlike, err := r.Cookie("ngo-postlike")
                    if err != nil {
                        rand.Seed(time.Now().UnixNano())
                        r1 := rand.Intn(10000)
                        r2 := rand.Intn(10000)
                        cookieStr = fmt.Sprintf("%s_%d_%d", ip, r1, r2)
                    } else {
                        cookieStr = cookiePostlike.Value
                    }                   
                    postlikeModel := &model.PostlikeSrvModel{} 
                    count := postlikeModel.GetLikeUnlikeCount(postId, userId, cookieStr)
                    if count > 0 {
                        postlike.Message = "Already Voted!"
                    } else {
                        postlike.Message = "Thank You!"
                        var value int64 = 1
                        _, ok = r.Form["unlike"]
                        if ok {
                            value = -1
                        }
                        postlikeModel.Insert(postId, value, ip, cookieStr, userId)
                    }          
                    postlike.LikeCount = postlikeModel.GetLikeCount(postId)
                    postlike.UnlikeCount = postlikeModel.GetUnlikeCount(postId)
                }    
            }
        }

        data, err := json.Marshal(postlike)
        if err != nil {
            res = `{"message":"", "likecount":0, "unlikecount":0 }`
        } else {
            res = string(data)
        }

        cookie := http.Cookie{Name: "ngo-postlike", Value: cookieStr, MaxAge: 3600*24*31}
        http.SetCookie(w, &cookie)        
        SendBack(w, res)             
    } else {
        res := `{"message":"", "likecount":0, "unlikecount":0 }`
        var postlike Postlike
        err := r.ParseForm()
        if err == nil {
            va, ok := r.Form["pid"]
            if ok {
                postId, err := strconv.ParseInt(va[0], 10, 64)
                if err == nil {
                    postlikeModel := &model.PostlikeSrvModel{}
                    postlike.LikeCount = postlikeModel.GetLikeCount(postId)
                    postlike.UnlikeCount = postlikeModel.GetUnlikeCount(postId)     
                }
            }                
        }

        data, err := json.Marshal(postlike)
        if err != nil {
            res = `{"message":"", "likecount":0, "unlikecount":0 }`
        } else {
            res = string(data)
        }

        SendBack(w, res)        
    }
}


