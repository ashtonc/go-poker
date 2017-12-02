package handlers

import (
	"log"
	"net/http"

	"poker/database"
	"poker/models"
	"poker/sessions"

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
		// Populate the data needed for the page (these should nearly all be external functions)
		var session = sessions.GetSession()
		session.PageHome = true

		pagedata := models.PageData{
			Session: session,
		}

		// Execute the template with our page data
		template := env.Templates["Home"]
		template.Execute(w, pagedata)
	})
}

func Login(env *models.Env) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		// func Login(response http.ResponseWriter, request *http.Request) {

		if request.Method == "POST" {
			userName := request.FormValue("username")
			name := request.FormValue("username")
			pass := request.FormValue("password")
			redirectTarget := "/poker/login/"
			if name != "" && pass != "" {

				// .. check credentials ..
				setSession(userName, name, response)
				redirectTarget = "/poker/game/"
			}
			//redirect to "404 page not found if user "
			http.Redirect(response, request, redirectTarget, 302)
		} else {
			// Populate the data needed for the page (these should nearly all be external functions)
			pagedata := models.PageData{
				Session: models.Session{
					LoggedIn:  false,
					PageLogin: true,
				},
			}

			// Execute the template with our page data
			template := env.Templates["Login"]
			template.Execute(response, pagedata)
		}

	})
}

func Logout(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//clearSession(response)
		http.Redirect(w, r, "/poker/", http.StatusTemporaryRedirect)
	})
}

func Register(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			// Populate the data needed for the page (these should nearly all be external functions)
			pagedata := models.PageData{
				Session: models.Session{
					LoggedIn: false,
					PageUser: true,
				},
			}

			// Execute the template with our page data
			template := env.Templates["Register"]
			template.Execute(w, pagedata)
		} else if r.Method == "POST" {

		}
	})
}

func ViewUser(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]

		// Get the user page matching that username from the database
		user, err := database.GetUserPage(env, username)
		if err != nil {
			log.Print("Player " + username + " not found.")

			// For now, just redirect them to the home page
			http.Redirect(w, r, "/poker/", http.StatusTemporaryRedirect)
			return
		}

		var session = sessions.GetSession()
		session.PageUser = true

		// Populate the data needed for the page (these should nearly all be external functions)
		pagedata := models.PageData{
			Session: session,
			UserPage: models.UserPage{
				MatchesSession: true,
				Username:       user.Username,
				Name:           user.Name,
				Email:          user.Email,
				PictureUrl:     user.PictureUrl,
			},
		}

		// Execute the template with our page data
		template := env.Templates["ViewUser"]
		template.Execute(w, pagedata)

		log.Print("Displaying player " + username + ".")
	})
}

func EditUser(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]

		var session = sessions.GetSession()
		session.PageUser = true

		// Populate the data needed for the page (these should nearly all be external functions)
		pagedata := models.PageData{
			Session: session,
			UserPage: models.UserPage{
				MatchesSession: true,
				Username:       username,
				Name:           "User Name",
				Email:          "user@email.ca",
			},
		}

		// Execute the template with our page data
		template := env.Templates["EditUser"]
		template.Execute(w, pagedata)
	})
}

func RedirectGame(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If someone is stitting at a table, send them to that table
		http.Redirect(w, r, "/poker/game/play", http.StatusTemporaryRedirect)
		// Else, send them to the lobby
		//http.Redirect(w, r, "/poker/game/lobby", http.StatusTemporaryRedirect)
	})
}

func PlayGame(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var session = sessions.GetSession()
		session.PageGame = true

		// Populate the data needed for the page (these should nearly all be external functions)
		pagedata := models.PageData{
			Session: session,
		}

		// Execute the template with our page data
		template := env.Templates["PlayGame"]
		template.Execute(w, pagedata)
	})
}

func ViewLobby(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		lobby, err := database.GetLobby(env)
		if err != nil {
			// No lobby exists or worse error
			// return
			log.Fatal(err)
		}

		var session = sessions.GetSession()
		session.PageGame = true

		// Populate the data needed for the page (these should nearly all be external functions)
		pagedata := models.PageData{
			Session: session,
			Lobby:   *lobby,
		}

		// Execute the template with our page data
		template := env.Templates["ViewLobby"]
		template.Execute(w, pagedata)
	})
}

func WatchGame(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var session = sessions.GetSession()
		session.PageGame = true

		// Populate the data needed for the page (these should nearly all be external functions)
		pagedata := models.PageData{
			Session: session,
		}

		// Execute the template with our page data
		template := env.Templates["WatchGame"]
		template.Execute(w, pagedata)
	})
}

func Leaderboard(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		leaderboard, err := database.GetLeaderboard(env)
		if err != nil {
			// Big error
		}

		var session = sessions.GetSession()
		session.PageLeaderboard = true

		pagedata := models.PageData{
			Session:     session,
			Leaderboard: *leaderboard,
		}

		// Execute the template with our page data
		template := env.Templates["Leaderboard"]
		template.Execute(w, pagedata)
	})
}
