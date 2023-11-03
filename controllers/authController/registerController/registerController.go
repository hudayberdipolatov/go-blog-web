package registerController

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/hudayberdipolatov/go-blog-web/helpers"
	"github.com/hudayberdipolatov/go-blog-web/models"
	"github.com/hudayberdipolatov/go-blog-web/validate/authValidate/registerValidate"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"reflect"
)

type RegisterController struct{}

func (register RegisterController) RegisterPage(w http.ResponseWriter, r *http.Request) {
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
	registerInput := registerValidate.RegisterValidate{
		Username:        r.PostForm.Get("username"),
		FullName:        r.PostForm.Get("full_name"),
		Email:           r.PostForm.Get("email"),
		Password:        r.PostForm.Get("password"),
		ConfirmPassword: r.PostForm.Get("confirm_password"),
	}

	translator := en.New()
	uni := ut.New(translator, translator)
	trans, _ := uni.GetTranslator("en")
	validate := validator.New()

	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		labelName := field.Tag.Get("label")
		return labelName
	})

	_ = validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} meýdany hökmany", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
	validateErrors := validate.Struct(registerInput)
	errorMessages := make(map[string]interface{})
	// user- barlamak ucin
	getUserEmail := userModel.GetUserEmail(registerInput.Email)
	getUserUsername := userModel.GetUserUsername(registerInput.Username)

	if validateErrors != nil {
		for _, e := range validateErrors.(validator.ValidationErrors) {
			errorMessages[e.StructField()] = e.Translate(trans)
		}
		//fmt.Println(errorMessages)
		data := map[string]interface{}{
			"validation": errorMessages,
			"user":       registerInput,
		}
		view, _ := template.ParseFiles(helpers.Include("auth/register")...)
		_ = view.ExecuteTemplate(w, "RegisterPage", data)
		return
	}

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
