package routes

import (
	"github.com/gorilla/mux"
	"github.com/hudayberdipolatov/go-blog-web/controllers/authController/loginController"
	"github.com/hudayberdipolatov/go-blog-web/controllers/authController/registerController"
	"github.com/hudayberdipolatov/go-blog-web/controllers/front/postController"
	"github.com/hudayberdipolatov/go-blog-web/controllers/userData/userPostController"
	"github.com/hudayberdipolatov/go-blog-web/controllers/userData/userProfileController"
	"net/http"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	var login loginController.LoginController
	var register registerController.RegisterController
	var frontPosts postController.FrontPostController
	var userPost userPostController.UserPostController
	var userProfile userProfileController.UserProfileController
	// auth routes begin
	// register routes
	router.HandleFunc("/register", register.RegisterPage).Methods("GET")
	router.HandleFunc("/register", register.Register).Methods("POST")

	// login routes
	router.HandleFunc("/login", login.LoginPage).Methods("GET")
	//auth routes end

	// web for routes
	router.HandleFunc("/", frontPosts.ListPost)
	// user data
	// user profile data
	router.HandleFunc("/user", userProfile.UserProfile).Methods("GET")

	// user posts
	router.HandleFunc("/user/posts", userPost.ListUserPost).Methods("GET")
	router.HandleFunc("/user/posts/create", userPost.CreateUserPost).Methods("GET")
	router.HandleFunc("/user/posts/store", userPost.StoreUserPost).Methods("POST")
	router.HandleFunc("/user/posts{post_id}/edit", userPost.EditUserPost).Methods("GET")
	router.HandleFunc("/user/post/{post_id}/update", userPost.UpdateUserPost).Methods("PUT")
	router.HandleFunc("/user/posts/{post_id}/delete", userPost.DeleteUserPost).Methods("DELETE")

	// File server
	fs := http.FileServer(http.Dir("./public/assets"))
	router.PathPrefix("/public/assets/").Handler(http.StripPrefix("/public/assets", fs))
	return router
}
