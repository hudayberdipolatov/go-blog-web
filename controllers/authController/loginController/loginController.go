package loginController

import (
	"github.com/hudayberdipolatov/go-blog-web/helpers"
	"html/template"
	"log"
	"net/http"
)

type LoginController struct{}

//type LoginControllerInterface interface {
//	LoginPage(w http.ResponseWriter, r *http.Request)
//	Login(w http.ResponseWriter, r *http.Request)
//}

func (login LoginController) LoginPage(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles(helpers.Include("auth/login")...)
	if err != nil {
		log.Println(err)
		return
	}
	view.ExecuteTemplate(w, "LoginPage", nil)
}

func (login LoginController) Login(w http.ResponseWriter, r *http.Request) {

}
