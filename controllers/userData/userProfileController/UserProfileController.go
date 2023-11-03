package userProfileController

import (
	"github.com/hudayberdipolatov/go-blog-web/helpers"
	"html/template"
	"log"
	"net/http"
)

type UserProfileController struct{}

func (userProfile UserProfileController) UserProfile(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles(helpers.Include("userData/userProfile")...)
	if err != nil {
		log.Println(err)
		return
	}
	view.ExecuteTemplate(w, "UserProfile", nil)
}
