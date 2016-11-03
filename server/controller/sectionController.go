/**
* @file sectionController.go
* @brief section controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
	"net/http"
	"log"
	"html/template"
)

type SectionController struct {
}

func NewSectionController() *SectionController {
	return &SectionController{}
}


func (controller *SectionController) IndexAction(w http.ResponseWriter, r *http.Request) {
     tmpl, err := template.ParseFiles("client/app/section/index.html")
    if err != nil {
        log.Fatal("SectionController::IndexAction: ", err)
    }

    err = tmpl.Execute(w, controller)
}
