package controller

import (
	"html/template"
	"net/http"
)

var homeController home

func Start(templates map[string]*template.Template) {
	homeController.homeTemplate = templates["home.html"]
	homeController.registerRoutes()
	http.Handle("/images/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
