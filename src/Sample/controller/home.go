package controller

import (
	"SampleGo/src/Sample/model"
	"SampleGo/src/Sample/viewmodel"
	"fmt"
	"html/template"
	"log"
	"os"

	"net/http"
)

type home struct {
	homeTemplate    *template.Template
	loginTemplate   *template.Template
	accountTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/login", h.handleLogin)
	http.HandleFunc("/account", h.handleAccount)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewHome2()
	h.homeTemplate.Execute(w, vm)

}

func (h home) handleLogin(w http.ResponseWriter, r *http.Request) {
	log := log.New(os.Stdout)
	vm := viewmodel.NewLogin()
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Info(fmt.Errorf("Error logging in: %v", err))
		}
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		if user, err := model.Login(email, password); err == nil {
			log.Infof("User has logged in: %v\n", user)
			vm := viewmodel.NewHome(*user)
			h.homeTemplate.Execute(w, vm)
			return
		} else {
			log.Errorf("Failed to log user in with email: %v, error was: %v\n", email, err)
			vm.Email = email
			vm.Password = password
		}
	}
	w.Header().Add("Content-Type", "text/html")
	h.loginTemplate.Execute(w, vm)
}

func (h home) handleAccount(w http.ResponseWriter, r *http.Request) {
	log := log.New(os.Stdout)
	vm := viewmodel.NewAccount()
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Infof(fmt.Errorf("Error logging in: %v", err))
		}
		email := r.Form.Get("email")
		password := r.Form.Get("psw")
		firstName := r.Form.Get("firstName")
		lastName := r.Form.Get("lastName")
		log.Infof("Data of users: %v, %v, %v\n", firstName, lastName, email)

		if user, err := model.AddNewUser(email, firstName, lastName, password); err == nil {
			log.Infof("User has logged in: %v\n", user)
			vm := viewmodel.NewHome(*user)
			h.homeTemplate.Execute(w, vm)
			return
		} else {
			log.Infof("Failed to log user in with email: %v, error was: %v\n", email, err)
			vm.Email = email
			vm.Password = password
		}

	}
	h.accountTemplate.Execute(w, vm)

}
