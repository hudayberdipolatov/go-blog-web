package userPostController

import (
	"github.com/hudayberdipolatov/go-blog-web/helpers"
	"github.com/hudayberdipolatov/go-blog-web/helpers/authHelpers"
	"html/template"
	"log"
	"net/http"
)

type UserPostController struct{}

// List user Posts Page

func (userPost UserPostController) ListUserPost(w http.ResponseWriter, r *http.Request) {
	session, _ := authHelpers.Store.Get(r, authHelpers.SESSION_ID)
	if len(session.Values) == 0 {
		view, _ := template.ParseFiles(helpers.Include("userData/userPosts/ListPost")...)
		_ = view.ExecuteTemplate(w, "ListUserPost", nil)
	} else {
		data := map[string]interface{}{
			"username": session.Values["Username"],
			"FullName": session.Values["FullName"],
			"Email":    session.Values["Email"],
			"loggedIn": session.Values["loggedIn"],
		}
		view, _ := template.ParseFiles(helpers.Include("userData/userPosts/ListPost")...)
		_ = view.ExecuteTemplate(w, "ListUserPost", data)
	}
}

// Create user Post page

func (userPost UserPostController) CreateUserPost(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles(helpers.Include("userData/userPosts/CreateUserPost")...)
	if err != nil {
		log.Println(err)
		return
	}
	_ = view.ExecuteTemplate(w, "CreateUserPost", nil)
}

// Store User Post

func (userPost UserPostController) StoreUserPost(w http.ResponseWriter, r *http.Request) {
	// store post
}

// Edit user Post page

func (userPost UserPostController) EditUserPost(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles(helpers.Include("userData/userPosts/EditUserPost")...)
	if err != nil {
		log.Println(err)
		return
	}
	_ = view.ExecuteTemplate(w, "EditUserPost", nil)
}

// update user post

func (userPost UserPostController) UpdateUserPost(w http.ResponseWriter, r *http.Request) {
	// update Post
}

// Delete user post

func (userPost UserPostController) DeleteUserPost(w http.ResponseWriter, r *http.Request) {
	// Delete Post
}
