/**
* @file loginController.go
* @brief login controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
	"net/http"
	"log"
	"html/template"
	model "go-angular/server/model"
)

type LoginController struct {
}

func NewLoginController() *LoginController {
	return &LoginController{}
}

func (controller *LoginController) IndexAction(w http.ResponseWriter, r *http.Request) {
    mainParams := &BlogMainParams{
        Stylesheets: []string {
            "../extensions/bootstrap-3.3.5/dist/css/bootstrap.min.css",
            "../app/login/styles/app.css" },
        Javscripts: []string {
            "../extensions/angular-1.5.0/angular.js",
            "../extensions/angular-1.5.0/angular-route.js",
            "../extensions/angular-1.5.0/angular-resource.js",
            "../app/login/scripts/directives/directives.js",
            "../app/login/scripts/services/services.js",
            "../app/login/scripts/app.js",
            "../app/login/scripts/controllers/controllers.js" },
        Startup : "" }

    linkModel := &model.LinkSrvModel{}
    linkList, _ := linkModel.FindAllLinks()
    mainParams.Shortcuts = make([]BlogLink, 0, 6)
    for _, link := range(linkList) {
        mainParams.Shortcuts = append(mainParams.Shortcuts, BlogLink{Name:link.Name, Url:link.Url})
    }

    menuModel := &model.MenuSrvModel{}
    menuList, _ := menuModel.FindAllMenus()
    mainParams.Menus = make([]BlogMenu, 0, 12)
    for _, menu := range(menuList) {
        var blogMenu BlogMenu
        blogMenu.MainMenu.Name = menu.MainMenu.Name
        blogMenu.MainMenu.Url = menu.MainMenu.Url
        blogMenu.SubMenu = make([]BlogLink, 0, 12)
        for _, subMenu := range(menu.SubMenu){
            blogMenu.SubMenu = append(blogMenu.SubMenu, BlogLink{Name:subMenu.Name, Url:subMenu.Url})
        }
        blogMenu.HasSubMenu = (len(blogMenu.SubMenu) > 0)
        mainParams.Menus = append(mainParams.Menus, blogMenu)
    }
    tmpl, err := template.ParseFiles("client/app/login/index.html")
    if err != nil {
        log.Fatal("BlogController::IndexAction: ", err)
    }

    err = tmpl.Execute(w, mainParams)
}
