package loginController

import (
	"errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/hudayberdipolatov/go-blog-web/helpers"
	"github.com/hudayberdipolatov/go-blog-web/helpers/authHelpers"
	"github.com/hudayberdipolatov/go-blog-web/models"
	"github.com/hudayberdipolatov/go-blog-web/validate/authValidate/loginValidate"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"reflect"
)

type LoginController struct{}

func (login LoginController) LoginPage(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles(helpers.Include("auth/login")...)
	if err != nil {
		log.Println(err)
		return
	}
	_ = view.ExecuteTemplate(w, "LoginPage", nil)
}

func (login LoginController) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var userModel models.Users
	loginInput := loginValidate.LoginValidate{
		Username: r.PostForm.Get("username"),
		Password: r.PostForm.Get("password"),
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
	validateErrors := validate.Struct(loginInput)
	errorMessages := make(map[string]interface{})

	if validateErrors != nil {
		for _, e := range validateErrors.(validator.ValidationErrors) {
			errorMessages[e.StructField()] = e.Translate(trans)
		}
		//fmt.Println(errorMessages)
		data := map[string]interface{}{
			"validation": errorMessages,
			"user":       loginInput,
		}
		//log.Println(data)
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
			user := userModel.GetUser(loginInput.Username)
			session, _ := authHelpers.Store.Get(r, authHelpers.SESSION_ID)
			session.Values["loggedIn"] = true
			session.Values["Username"] = user.Username
			session.Values["FullName"] = user.FullName
			session.Values["Email"] = user.Email
			session.Values["user_id"] = user.ID
			_ = session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}
