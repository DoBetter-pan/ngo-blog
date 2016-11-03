/**
* @file categoryController.go
* @brief category controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
	"net/http"
	"log"
	"html/template"
)

type CategoryController struct {
}

func NewCategoryController() *CategoryController {
	return &CategoryController{}
}


func (controller *CategoryController) IndexAction(w http.ResponseWriter, r *http.Request) {
     tmpl, err := template.ParseFiles("client/app/category/index.html")
    if err != nil {
        log.Fatal("CategoryController::IndexAction: ", err)
    }

    err = tmpl.Execute(w, controller)
}
