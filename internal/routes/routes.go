package routes

import (
	"net/http"

	"bitbucket.org/janpavtel/site/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func CreateRoutes(view *handlers.View) http.Handler {
	mux := chi.NewMux()

	mux.Get("/", view.Home)
	mux.Get("/about", view.About)
	mux.Get("/signup", view.SignUp)


	mux.Get("/users", view.ShowUsers)
	mux.Get("/users/{id}", view.ShowUser)
	mux.Get("/users/{id}/edit", view.EditUser)
	mux.Post("/users/{id}", view.UpdateUser)
	mux.Post("/users", view.NewUserCreation)
	


	mux.Get("/users/login", view.UserLoginShow)
	mux.Get("/users/login/captcha.jpg", view.GenerateCaptcha)
	mux.Post("/users/login", view.UserLoginPost)

	mux.Route("/news", func(mux chi.Router){
		mux.Get("/all", view.NewsAll)
	})


	mux.Route("/posts", func(mux chi.Router) {
		mux.Use(AuthenticationValidation, UpdateJWT)

		mux.Get("/create", view.PostsCreateShow)
		mux.Post("/create", view.PostsCreatePosts)

		mux.Get("/myposts", view.PostsMyPosts)
		mux.Get("/allposts", view.PostsAllPosts)


		mux.Route("/{id}", func(mux chi.Router) {
			mux.Use(AdminOrEditorRoleValidation)
	
			mux.Get("/edit", view.PostsEditShow)
			mux.Post("/", view.PostsUpdate)
		})
	})



	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(AuthenticationValidation, UpdateJWT, AdminRoleValidation)

		mux.Get("/users", view.ManagementGetUsers)
		mux.Get("/users/{id}", view.ManagementEditUser)
		mux.Post("/users/{id}", view.ManagementUpdateUser)
	})

	mux.Route("/users/profile", func(mux chi.Router) {
		mux.Use(AuthenticationValidation, UpdateJWT)

		mux.Get("/edit", view.UserProfileEditShow)
		mux.Post("/edit", view.UserProfileEditUpdate)

		mux.Post("/image/upload", view.UserProfileUploadImage)
		mux.Get("/image", view.UserProfileImage)
	})

	fileServer := http.FileServer(http.Dir("./www/static/"))
	
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}