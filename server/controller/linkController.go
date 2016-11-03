/**
* @file linkController.go
* @brief link controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
	"net/http"
	"log"
	"html/template"
)

type LinkController struct {
}

func NewLinkController() *LinkController {
	return &LinkController{}
}


func (controller *LinkController) IndexAction(w http.ResponseWriter, r *http.Request) {
     tmpl, err := template.ParseFiles("client/app/link/index.html")
    if err != nil {
        log.Fatal("LinkController::IndexAction: ", err)
    }

    err = tmpl.Execute(w, controller)
}
