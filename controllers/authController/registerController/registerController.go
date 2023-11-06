package registerController

import (
	"github.com/hudayberdipolatov/go-blog-web/helpers"
	"github.com/hudayberdipolatov/go-blog-web/helpers/authHelpers"
	"github.com/hudayberdipolatov/go-blog-web/models"
	"github.com/hudayberdipolatov/go-blog-web/validate"
	"github.com/hudayberdipolatov/go-blog-web/validate/authValidate/registerValidate"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
)

type RegisterController struct{}

func (register RegisterController) RegisterPage(w http.ResponseWriter, r *http.Request) {
	session, _ := authHelpers.Store.Get(r, authHelpers.SESSION_ID)
	if len(session.Values) != 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	view, err := template.ParseFiles(helpers.Include("auth/register")...)
	if err != nil {
		log.Println(err)
		return
	}
	_ = view.ExecuteTemplate(w, "RegisterPage", nil)
}

func (register RegisterController) Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var userModel models.Users
	var validation validate.Validation
	registerInput := registerValidate.RegisterValidate{
		Username:        r.PostForm.Get("username"),
		FullName:        r.PostForm.Get("full_name"),
		Email:           r.PostForm.Get("email"),
		Password:        r.PostForm.Get("password"),
		ConfirmPassword: r.PostForm.Get("confirm_password"),
	}
	errorMessages := validation.Struct(registerInput)
	if errorMessages != nil {
		data := map[string]interface{}{
			"validation": errorMessages,
			"user":       registerInput,
		}
		view, _ := template.ParseFiles(helpers.Include("auth/register")...)
		_ = view.ExecuteTemplate(w, "RegisterPage", data)
		return
	}
	// user- barlamak ucin
	getUserEmail := userModel.GetUserEmail(registerInput.Email)
	getUserUsername := userModel.GetUserUsername(registerInput.Username)
	if getUserEmail || getUserUsername {
		ErrorUserMessage := make(map[string]interface{})
		if getUserEmail == true {
			ErrorUserMessage["has_email"] = "Email salgysy eýýäm ulanylýar!!!"
		}
		if getUserUsername == true {
			ErrorUserMessage["has_username"] = "Username ady eýýäm ulanylýar!!!"
		}
		data := map[string]interface{}{
			"user_exists": ErrorUserMessage,
			"user":        registerInput,
		}
		view, _ := template.ParseFiles(helpers.Include("auth/register")...)
		_ = view.ExecuteTemplate(w, "RegisterPage", data)
		return
	}
	// username we email address user table-de yok bolsa register bolmak ucin
	if !getUserEmail && !getUserUsername {
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(registerInput.Password), bcrypt.DefaultCost)
		models.Users{
			Username: registerInput.Username,
			FullName: registerInput.FullName,
			Email:    registerInput.Email,
			Password: string(hashPassword),
		}.CreateUser()
		data := map[string]interface{}{
			"success": "Siz üsdünlikli hasaba alyndyňyz. Indi ulgama girip bilersiňiz!!!",
		}
		view, _ := template.ParseFiles(helpers.Include("auth/login")...)
		_ = view.ExecuteTemplate(w, "LoginPage", data)
	}
	return
}
