package handlers

import (
	"fmt"
	"net/http"
	"html/template"

	"poker/models"

	"github.com/gorilla/mux"
)

// This simply redirects users to /poker
func HomeRedirect(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/poker/", http.StatusTemporaryRedirect)
	})
}

func Home(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Home")
	})
}

func Login(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Login")
	})
}

func Register(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Register")
	})
}

func User(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]

		// Populate the data needed for the page (these should nearly all be external functions)
		pagedata := models.PageData{
			Session: models.Session{
				LoggedIn: true,
				Username: "current-user",
				Name: "Current User",
				PageUser: true,
			},
			UserPage: models.UserPage{
				MatchesSession: true,
				Username: username,
				Name: "User Name",
				Email: "user@email.ca",
			},
		}

		// Build our template using the required files (need base, head, navigation, and content)
		// This should be moved to a caching function
		t, _ := template.ParseFiles("./templates/base.tmpl", "./templates/head_base.tmpl", "./templates/navigation.tmpl", "./templates/user_view.tmpl")

		// Execute the template with our page data
		t.Execute(w, pagedata)
	})
}

func UserEdit(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Edit a user")
	})
}

func Game(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Game")
	})
}

func Leaderboard(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Leaderboard")
	})
}
