package authHelpers

import (
	"github.com/gorilla/sessions"
	"github.com/hudayberdipolatov/go-blog-web/models"
	"net/http"
)

var SESSION_ID = "USER-AUTH"

var Store = sessions.NewCookieStore([]byte("hf#sk_jhf_jk23h4jkh234jhg235!!!3g563534534_@dfhsdfsdfjh"))

func SessionUserData(w http.ResponseWriter, r *http.Request, username string) {
	user := models.Users{}.GetUser(username)
	session, _ := Store.Get(r, SESSION_ID)
	session.Values["loggedIn"] = true
	session.Values["Username"] = user.Username
	session.Values["FullName"] = user.FullName
	session.Values["Email"] = user.Email
	session.Values["user_id"] = user.ID
	_ = session.Save(r, w)
}

func SessionGetUserData(r *http.Request) map[string]interface{} {
	session, _ := Store.Get(r, SESSION_ID)
	userData := map[string]interface{}{
		"username": session.Values["Username"],
		"FullName": session.Values["FullName"],
		"Email":    session.Values["Email"],
		"loggedIn": session.Values["loggedIn"],
	}
	return userData
}
