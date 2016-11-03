/**
* @file mainController.go
* @brief main controller
* @author yingx
* @date 2015-01-10
 */

package controller

import (
	"net/http"
	"log"
	_ "fmt"
	"bytes"
	"html/template"
)

type MainController struct {
	Title string
	Stylesheets []string
	Javscripts []string
	Banner template.HTML
	Content template.HTML
	Footer template.HTML
	Startup template.HTML
}

func LoadInfoFromTemplate(filename string) string {
	var b bytes.Buffer
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		log.Fatal("MainController::LoadInfoFromTemplate: ", err)
	}

	err = tmpl.Execute(&b, nil)
	return b.String()
}

func LoadHeaderFromTemplate() template.HTML {
	return template.HTML(LoadInfoFromTemplate("server/views/main/header.html"))
}

func LoadContentFromTemplate() template.HTML {
	return template.HTML(LoadInfoFromTemplate("server/views/main/content.html"))
}

func LoadFooterFromTemplate() template.HTML {
	return template.HTML(LoadInfoFromTemplate("server/views/main/footer.html"))
}

func NewMainController() *MainController {
	controller := &MainController{
		Title: "go-angular",
		Stylesheets: []string {
			"../extensions/bootstrap-3.3.5/dist/css/bootstrap.min.css",
            "../assets/css/ie10-viewport-bug-workaround.css",
            "../extensions/angular-1.5.0/togglemenu_boot/togglemenu.css",
            "../extensions/angular-1.5.0/angular-aside/css/angular-aside.min.css",
            "../extensions/angular-1.5.0/angular-busy/angular-busy.css",
            "../extensions/angular-1.5.0/toaster/toaster.css",
            "../css/MaterialIcons/material-icons.css",
            "../css/fontawesome/css/font-awesome.min.css",
			"../css/normalize.css",
			"../css/main-style.css" },
		Javscripts: []string {
			"../js/jquery-1.11.3/jquery-1.11.3.min.js",
            "../extensions/bootstrap-3.3.5/dist/js/bootstrap.min.js",
            "../extensions/angular-1.5.0/angular.min.js",
            "../extensions/angular-1.5.0/ui-bootstrap-tpls-1.2.0.min.js",
            "../extensions/angular-1.5.0/angular-route.min.js",
            "../extensions/angular-1.5.0/togglemenu_boot/togglemenu.js",
            "../extensions/angular-1.5.0/angular-aside/js/angular-aside.min.js",
            "../extensions/angular-1.5.0/angular-sanitize.min.js",
            "../extensions/angular-1.5.0/angular-animate/angular-animate.min.js",
            "../extensions/angular-1.5.0/i18n/angular-locale_zh-cn.js",
            "../extensions/angular-1.5.0/angular-busy/angular-busy.js",
            "../extensions/angular-1.5.0/toaster/toaster.js",
            "../js/app.js"},
		Startup : "" }
	controller.Banner = LoadHeaderFromTemplate()
	controller.Content = LoadContentFromTemplate()
	controller.Footer = LoadFooterFromTemplate()
	return controller
}

func (controller *MainController) RenderMainFrame(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("server/views/main/main.html")
	if err != nil {
		log.Fatal("MainController::RenderMainFrame: ", err)
	}

	err = tmpl.Execute(w, controller)
}
