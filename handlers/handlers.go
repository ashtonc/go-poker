package handlers

import (
	"log"
	"net/http"
	"fmt"

	"poker/database"
	"poker/models"

	"github.com/gorilla/mux"
)

// This simply redirects users to /poker/
func HomeRedirect(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
	})
}

func Home(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("User visited Home page.\n")
		err := database.CreateLeaderboardEntries(env) // to be removed later, testing inserting lobbies into database
		if err != nil {
			panic("No database found")
		}

		/*		// Populate the data needed for the page (these should nearly all be external functions)
				vars := mux.Vars(r)
				username := vars["username"]*/

		/*		// Get the user page matching that username from the database
				user, err := database.UserRegister(env, username)
				if err != nil {
					// TODO
				}*/

		// Populate the data needed for the page (these should nearly all be external functions)
		pagedata := getPageData("sessionid", "Home")

		// Execute the template with our page data
		template := env.Templates["Home"]
		template.Execute(w, pagedata)
	})
}

func Login(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			userName := r.FormValue("username")
			name := r.FormValue("username")
			pass := r.FormValue("password")
			redirectTarget := env.SiteRoot + "/login/"
			if name != "" && pass != "" {

				// .. check credentials ..
				setSession(userName, name, w)
				redirectTarget = env.SiteRoot + "/game/"
			}
			//redirect to "404 page not found if user "
			http.Redirect(w, r, redirectTarget, 302)
		} else if r.Method == "GET" {
			// Populate the data needed for the page (these should nearly all be external functions)
			pagedata := models.PageData{
				Session: models.Session{
					LoggedIn:  false,
					PageLogin: true,
				},
			}

			// Execute the template with our page data
			template := env.Templates["Login"]
			template.Execute(w, pagedata)
		}

	})
}

func Logout(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//clearSession(response)
		http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
	})
}

func Register(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("User visited Register page.\n")

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

			pagedata := models.PageData{
				Session: models.Session{
					LoggedIn: false,
					PageUser: true,
				},
			}

			template := env.Templates["Register"]


			fmt.Printf("User attempted to register.\n")
			r.ParseForm()
			username := r.PostFormValue("username")
			password := r.PostFormValue("password")
			name := r.PostFormValue("name")
			email := r.PostFormValue("email")
			password_repeat := r.PostFormValue("password-repeat")

			if len(username) < 5 {
				template.Execute(w, pagedata)
			} else if len(name) < 1 {
				template.Execute(w, pagedata)
			} else if len(password) < 6 {
				template.Execute(w, pagedata)
			} else if password != password_repeat {
				template.Execute(w, pagedata)
			// } else if database.UserCount(env, username) == nil {
			// 	panic("HI!")
			}

			err := database.UserRegister(env, username, name, email, password)
			if err != nil {
				panic("No database found")
			}

			fmt.Printf(username, password, name, email, password_repeat)
			http.Redirect(w, r, env.SiteRoot+"/game/play", http.StatusTemporaryRedirect)

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
			http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
			return
		}

		// Populate the data needed for the page
		pagedata := getPageData("sessionid", "ViewUser")
		pagedata.UserPage = models.UserPage{
			MatchesSession: true,
			Username:       user.Username,
			Name:           user.Name,
			Email:          user.Email,
			PictureUrl:     user.PictureUrl,
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

		// Populate the data needed for the page
		pagedata := getPageData("sessionid", "EditUser")
		pagedata.UserPage = models.UserPage{
			MatchesSession: true,
			Username:       username,
			Name:           "User Name",
			Email:          "user@email.ca",
		}

		// Execute the template with our page data
		template := env.Templates["EditUser"]
		template.Execute(w, pagedata)
	})
}

func RedirectGame(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If someone is stitting at a table, send them to that table
		http.Redirect(w, r, env.SiteRoot+"/game/play", http.StatusTemporaryRedirect)
		// Else, send them to the lobby
		//http.Redirect(w, r, env.SiteRoot /game/lobby", http.StatusTemporaryRedirect)
	})
}

func PlayGame(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Populate the data needed for the page
		pagedata := getPageData("sessionid", "PlayGame")

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

		// Populate the data needed for the page
		pagedata := getPageData("sessionid", "ViewLobby")
		pagedata.Lobby = *lobby

		// Execute the template with our page data
		template := env.Templates["ViewLobby"]
		template.Execute(w, pagedata)
	})
}

func WatchGame(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Populate the data needed for the page
		pagedata := getPageData("sessionid", "WatchGame")

		// Execute the template with our page data
		template := env.Templates["WatchGame"]
		template.Execute(w, pagedata)
	})
}

func Leaderboard(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("User visited Leaderboard page.\n")

		leaderboard, err := database.GetLeaderboard(env)
		if err != nil {
			// Big error
		}

		// Populate the data needed for the page
		pagedata := getPageData("sessionid", "ViewLobby")
		pagedata.Leaderboard = *leaderboard

		// Execute the template with our page data
		template := env.Templates["Leaderboard"]
		template.Execute(w, pagedata)

		// fmt.Printf(leaderboard)

	})
}
