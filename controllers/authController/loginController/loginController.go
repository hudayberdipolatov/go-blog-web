package loginController

import (
	"errors"
	"github.com/hudayberdipolatov/go-blog-web/helpers"
	"github.com/hudayberdipolatov/go-blog-web/helpers/authHelpers"
	"github.com/hudayberdipolatov/go-blog-web/models"
	"github.com/hudayberdipolatov/go-blog-web/validate"
	"github.com/hudayberdipolatov/go-blog-web/validate/authValidate/loginValidate"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

type LoginController struct{}

func (login LoginController) LoginPage(w http.ResponseWriter, r *http.Request) {
	session, _ := authHelpers.Store.Get(r, authHelpers.SESSION_ID)
	if len(session.Values) != 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	view, _ := template.ParseFiles(helpers.Include("auth/login")...)
	_ = view.ExecuteTemplate(w, "LoginPage", nil)
}

func (login LoginController) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var userModel models.Users
	var validation validate.Validation
	loginInput := loginValidate.LoginValidate{
		Username: r.PostForm.Get("username"),
		Password: r.PostForm.Get("password"),
	}
	errorMessages := validation.Struct(loginInput)
	if errorMessages != nil {
		data := map[string]interface{}{
			"validation": errorMessages,
			"user":       loginInput,
		}
		view, _ := template.ParseFiles(helpers.Include("auth/login")...)
		_ = view.ExecuteTemplate(w, "LoginPage", data)
		return
	} else {
		getUser := userModel.GetUser(loginInput.Username)
		var errorMessage error
		if getUser.ID == 0 {
			errorMessage = errors.New("Username ýa-da password yalnyş!!!")
		} else {
			errPassword := bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(loginInput.Password))
			if errPassword != nil {
				errorMessage = errors.New("Username ýa-da password yalnyş!!!")
			}
		}
		if errorMessage != nil {
			data := map[string]interface{}{
				"error": errorMessage,
			}
			view, _ := template.ParseFiles(helpers.Include("auth/login")...)
			_ = view.ExecuteTemplate(w, "LoginPage", data)
		} else {
			authHelpers.SessionUserData(w, r, loginInput.Username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func (login LoginController) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := authHelpers.Store.Get(r, authHelpers.SESSION_ID)
	session.Options.MaxAge = -1
	_ = session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
