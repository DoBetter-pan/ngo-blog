/**
* @file basicController.go
* @brief basic controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
	"net/http"
	"log"
	"html/template"
)

type BasicController struct {
}

func NewBasicController() *BasicController {
	return &BasicController{}
}


func (controller *BasicController) IndexAction(w http.ResponseWriter, r *http.Request) {
     tmpl, err := template.ParseFiles("client/app/basic/index.html")
    if err != nil {
        log.Fatal("BasicController::IndexAction: ", err)
    }

    err = tmpl.Execute(w, controller)
}
