package postController

import (
	"github.com/hudayberdipolatov/go-blog-web/helpers"
	"github.com/hudayberdipolatov/go-blog-web/helpers/authHelpers"
	"html/template"
	"net/http"
)

type FrontPostController struct{}

func (frontPost FrontPostController) ListPost(w http.ResponseWriter, r *http.Request) {
	session, _ := authHelpers.Store.Get(r, authHelpers.SESSION_ID)
	if len(session.Values) == 0 {
		view, _ := template.ParseFiles(helpers.Include("posts")...)
		_ = view.ExecuteTemplate(w, "FrontPosts", nil)
	} else {
		data := map[string]interface{}{
			"username": session.Values["Username"],
			"FullName": session.Values["FullName"],
			"Email":    session.Values["Email"],
			"loggedIn": session.Values["loggedIn"],
		}
		view, _ := template.ParseFiles(helpers.Include("posts")...)
		_ = view.ExecuteTemplate(w, "FrontPosts", data)
	}
}
