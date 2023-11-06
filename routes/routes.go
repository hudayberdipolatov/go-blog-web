package routes

import (
	"github.com/gorilla/mux"
	"github.com/hudayberdipolatov/go-blog-web/controllers/authController/loginController"
	"github.com/hudayberdipolatov/go-blog-web/controllers/authController/registerController"
	"github.com/hudayberdipolatov/go-blog-web/controllers/front/postController"
	"github.com/hudayberdipolatov/go-blog-web/controllers/userData/userPostController"
	"github.com/hudayberdipolatov/go-blog-web/controllers/userData/userProfileController"
	"github.com/hudayberdipolatov/go-blog-web/helpers/authHelpers"
	"net/http"
)

func Routes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	var login loginController.LoginController
	var register registerController.RegisterController
	var frontPosts postController.FrontPostController
	var userPost userPostController.UserPostController
	var userProfile userProfileController.UserProfileController
	// auth routes begin
	// register routes
	router.HandleFunc("/register", register.RegisterPage).Methods("GET")
	router.HandleFunc("/register", register.Register).Methods(http.MethodPost)
	// login routes
	router.HandleFunc("/login", login.LoginPage).Methods("GET")
	router.HandleFunc("/login", login.Login).Methods(http.MethodPost)
	router.HandleFunc("/logout", login.Logout).Methods(http.MethodPost)
	//auth routes end

	// web for routes
	router.HandleFunc("/", frontPosts.ListPost)
	// user data
	// user profile data
	router.HandleFunc("/user", AuthMiddleware(userProfile.UserProfile)).Methods("GET")

	// user posts
	router.HandleFunc("/user/posts", AuthMiddleware(userPost.ListUserPost)).Methods("GET")
	router.HandleFunc("/user/posts/create", AuthMiddleware(userPost.CreateUserPost)).Methods("GET")
	router.HandleFunc("/user/posts/store", AuthMiddleware(userPost.StoreUserPost)).Methods(http.MethodPost)
	router.HandleFunc("/user/posts{post_id}/edit", AuthMiddleware(userPost.EditUserPost)).Methods("GET")
	router.HandleFunc("/user/post/{post_id}/update", AuthMiddleware(userPost.UpdateUserPost)).Methods(http.MethodPut)
	router.HandleFunc("/user/posts/{post_id}/delete", AuthMiddleware(userPost.DeleteUserPost)).Methods(http.MethodDelete)

	// File server
	fs := http.FileServer(http.Dir("./public/assets"))
	router.PathPrefix("/public/assets/").Handler(http.StripPrefix("/public/assets", fs))
	return router
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := authHelpers.Store.Get(r, authHelpers.SESSION_ID)
		if len(session.Values) == 0 {
			//log.Println("user no login")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		} else {
			//log.Println("user login success")
			next.ServeHTTP(w, r)
			return
		}
	}
}
