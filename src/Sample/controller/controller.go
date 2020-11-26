package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController    home
	marriedController married
)

func Start(templates map[string]*template.Template) {

	homeController.homeTemplate = templates["home.html"]
	marriedController.marriedTemplate = templates["married.html"]
	homeController.registerRoutes()
	marriedController.registerRoutes()
	http.Handle("/images/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
