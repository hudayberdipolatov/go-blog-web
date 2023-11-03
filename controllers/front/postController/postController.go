package postController

import (
	"github.com/hudayberdipolatov/go-blog-web/helpers"
	"html/template"
	"net/http"
)

type FrontPostController struct{}

func (frontPost FrontPostController) ListPost(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles(helpers.Include("posts")...)
	_ = view.ExecuteTemplate(w, "FrontPosts", nil)

}
