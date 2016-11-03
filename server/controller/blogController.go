/**
* @file blogController.go
* @brief blog controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
    _ "fmt"
	"net/http"
	"log"
	"html/template"
	model "go-angular/server/model"
)

type BlogLink struct {
    Name string
    Url string
}

type BlogMenu struct {
    MainMenu BlogLink
    HasSubMenu bool
    SubMenu []BlogLink
}

type BlogMainParams struct {
    Stylesheets []string
    Javscripts []string
    Startup template.HTML
    Shortcuts []BlogLink
    Menus []BlogMenu
}

type BlogNewParams struct {
    Stylesheets []string
    Javscripts []string
    Startup template.HTML
}

type BlogController struct {
}

func NewBlogController() *BlogController {
    controller := &BlogController{}
    return controller
}

func (controller *BlogController) CheckRules() map[string] []string {
    rules := make(map[string] []string, 5)
    rules["AdminAction"] = []string {
        "allow admin",
        "deny *" }
    return rules
}

func (controller *BlogController) IndexAction(w http.ResponseWriter, r *http.Request) {
    mainParams := &BlogMainParams{
        Stylesheets: []string {
            "../extensions/bootstrap-3.3.5/dist/css/bootstrap.min.css",
            "../app/blog/styles/blog.css" },
        Javscripts: []string {
            "../extensions/angular-1.5.0/angular.js",
            "../extensions/angular-1.5.0/angular-route.js",
            "../extensions/angular-1.5.0/angular-resource.js",
            "../app/blog/scripts/directives/directives.js",
            "../app/blog/scripts/directives/visitorcounter.js",
            "../app/blog/scripts/services/services.js",
            "../app/blog/scripts/blog.js",
            "../app/blog/scripts/controllers/blog.js" },
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
    tmpl, err := template.ParseFiles("client/app/blog/blog.html")
    if err != nil {
        log.Fatal("BlogController::IndexAction: ", err)
    }

    err = tmpl.Execute(w, mainParams)
}

func (controller *BlogController) AdminAction(w http.ResponseWriter, r *http.Request) {
    /*
    startup := `
    <script type="text/javascript">
    $(function() {
        var ng_writer = new Simditor({ textarea: $('#ng-writer')});
    }); 
    </script>`
    */

    newParams := &BlogNewParams {
        Stylesheets: []string {
            "../extensions/bootstrap-3.3.5/dist/css/bootstrap.min.css",
            "../extensions/simditor/styles/simditor.css",
            "../app/blog/styles/admin.css" },
        Javscripts: []string {
            "../js/jquery-1.11.3/jquery-1.11.3.min.js",
            "../extensions/bootstrap-3.3.5/dist/js/bootstrap.min.js",
            "../extensions/simditor/scripts/module.js",
            "../extensions/simditor/scripts/hotkeys.js",
            "../extensions/simditor/scripts/uploader.js",
            "../extensions/simditor/scripts/simditor.js",
            "../extensions/angular-1.5.0/angular.js",
            "../extensions/angular-1.5.0/angular-route.js",
            "../extensions/angular-1.5.0/angular-resource.js",
            "../app/blog/scripts/directives/directives.js",
            "../app/blog/scripts/directives/ngsimditor.js",
            "../app/blog/scripts/services/services.js",
            "../app/section/scripts/services/services.js",
            "../app/category/scripts/services/services.js",
            "../app/blog/scripts/admin.js",
            "../app/blog/scripts/controllers/admin.js" },
        Startup : "" }

    //newParams.Startup = template.HTML(startup)

    tmpl, err := template.ParseFiles("client/app/blog/admin.html")
    if err != nil {
        log.Fatal("BlogController::NewAction: ", err)
    }

    err = tmpl.Execute(w, newParams)
}
