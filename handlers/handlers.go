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

func User(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]

		// Populate the data needed for the page
		pagedata := models.PageData{
			Session: models.Session{
				LoggedIn: true,
				Username: "current-user",
				Name: "Current User Name",
				PageUser: true,
			},
			UserPage: models.UserPage{
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

func Game(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Game")
	})
}
