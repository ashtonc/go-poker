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
		pagedata := models.PageData{
			Session: models.Session{
				LoggedIn: true,
				Username: "current-user",
				Name: "Current User",
				PageHome: true,
			},
		}

		// Build our template using the required files (need base, head, navigation, and content)
		// This should be moved to a caching function: https://elithrar.github.io/article/approximating-html-template-inheritance/
		t, _ := template.ParseFiles("./templates/base.tmpl", "./templates/head_base.tmpl", "./templates/navigation.tmpl", "./templates/index.tmpl")

		// Execute the template with our page data
		t.Execute(w, pagedata)
	})
}

func Login(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pagedata := models.PageData{
			Session: models.Session{
				LoggedIn: false,
				PageLogin: true,
			},
		}

		// Build our template using the required files (need base, head, navigation, and content)
		// This should be moved to a caching function: https://elithrar.github.io/article/approximating-html-template-inheritance/
		t, _ := template.ParseFiles("./templates/base.tmpl", "./templates/head_base.tmpl", "./templates/navigation.tmpl", "./templates/login.tmpl")

		// Execute the template with our page data
		t.Execute(w, pagedata)
	})
}

func Register(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pagedata := models.PageData{
			Session: models.Session{
				LoggedIn: false,
				PageRegister: true,
			},
		}

		// Build our template using the required files (need base, head, navigation, and content)
		// This should be moved to a caching function: https://elithrar.github.io/article/approximating-html-template-inheritance/
		t, _ := template.ParseFiles("./templates/base.tmpl", "./templates/head_base.tmpl", "./templates/navigation.tmpl", "./templates/register.tmpl")

		// Execute the template with our page data
		t.Execute(w, pagedata)
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
		// This should be moved to a caching function: https://elithrar.github.io/article/approximating-html-template-inheritance/
		t, _ := template.ParseFiles("./templates/base.tmpl", "./templates/head_base.tmpl", "./templates/navigation.tmpl", "./templates/user_view.tmpl")

		// Execute the template with our page data
		t.Execute(w, pagedata)
	})
}

func UserEdit(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]

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
		// This should be moved to a caching function: https://elithrar.github.io/article/approximating-html-template-inheritance/
		t, _ := template.ParseFiles("./templates/base.tmpl", "./templates/head_base.tmpl", "./templates/navigation.tmpl", "./templates/user_edit.tmpl")

		// Execute the template with our page data
		t.Execute(w, pagedata)
	})
}

func Game(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pagedata := models.PageData{
			Session: models.Session{
				LoggedIn: true,
				Username: "current-user",
				Name: "Current User",
				PageGame: true,
			},
		}

		// Build our template using the required files (need base, head, navigation, and content)
		// This should be moved to a caching function: https://elithrar.github.io/article/approximating-html-template-inheritance/
		t, _ := template.ParseFiles("./templates/base.tmpl", "./templates/head_base.tmpl", "./templates/navigation.tmpl", "./templates/game_lobby.tmpl")

		// Execute the template with our page data
		t.Execute(w, pagedata)
	})
}

func Leaderboard(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Leaderboard")
	})
}
