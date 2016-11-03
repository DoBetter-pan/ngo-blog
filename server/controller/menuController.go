/**
* @file menuController.go
* @brief menu controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
	"net/http"
	"log"
	"html/template"
)

type MenuController struct {
}

func NewMenuController() *MenuController {
	return &MenuController{}
}


func (controller *MenuController) IndexAction(w http.ResponseWriter, r *http.Request) {
     tmpl, err := template.ParseFiles("client/app/menu/index.html")
    if err != nil {
        log.Fatal("MenuController::IndexAction: ", err)
    }

    err = tmpl.Execute(w, controller)
}
