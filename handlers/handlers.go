package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"

	"poker/database"
	"poker/models"
)

// This simply redirects users to the site root
func HomeRedirect(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
	})
}

func Home(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Populate the data needed for the page (these should nearly all be external functions)
		pagedata := getPageData(env, "sessionid", "Home")

		// Execute the template with our page data
		template := env.Templates["Home"]
		template.Execute(w, pagedata)
	})
}

func Login(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		template := env.Templates["Login"]

		// Populate the data needed for the page (these should nearly all be external functions)
		pagedata := models.PageData{
			Session: models.Session{
				LoggedIn:  false,
				PageLogin: true,
			},
		}

		if r.Method == "POST" {
			fmt.Printf("POST for Login\n")
			// this code gets username and password from the POST form
			r.ParseForm()
			username := r.PostFormValue("username")
			password := r.PostFormValue("password")
			userAccount := database.FindByUsername(env, username)
			name := userAccount.Name

			// Query the database to check if account even exists (via FindByUsername?)
			if username == "" || password == "" {
				fmt.Printf("One or more fields were left blank.\n")
				template.Execute(w, pagedata)
			} else if userAccount.Username != username {
				fmt.Printf("This user does not exist.\n")
				template.Execute(w, pagedata)
			} else if userAccount.Password != password {
				fmt.Printf("The password is incorrect.\n")
				template.Execute(w, pagedata)
			} else {
				fmt.Printf("Login successful.\n")

				// .. check credentials ..
				setSession(username, name, w)

				http.Redirect(w, r, env.SiteRoot+"/game/example/play", http.StatusTemporaryRedirect)
			}

		} else if r.Method == "GET" {
			fmt.Printf("GET login page\n")

			// Execute the template with our page data
			template.Execute(w, pagedata)
		}

	})
}

// func Logout(env *models.Env) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		clearSession(response)
// 		http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
// 	})
// }

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
			isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString

			fmt.Printf("User attempted to register.\n")
			r.ParseForm()
			username := r.PostFormValue("username")
			password := r.PostFormValue("password")
			name := r.PostFormValue("name")
			email := r.PostFormValue("email")
			password_repeat := r.PostFormValue("password-repeat")

			if len(username) < 5 || !isAlpha(username) {
				template.Execute(w, pagedata)
				fmt.Printf("Incorrect input for username.\n")
			} else if len(name) < 1 || !isAlpha(name) {
				template.Execute(w, pagedata)
				fmt.Printf("Incorrect input for name.\n")
			} else if len(email) < 1 {
				template.Execute(w, pagedata)
				fmt.Printf("Incorrect input for email.\n")
			} else if len(password) < 6 {
				template.Execute(w, pagedata)
				fmt.Printf("Incorrect input for password.\n")
			} else if password != password_repeat {
				template.Execute(w, pagedata)
				fmt.Printf("The password field should match the password repeat field.\n")
			} else if (database.FindByUsername(env, username)).Username == username {
				template.Execute(w, pagedata)
				fmt.Printf("This account name already exists.\n")
			} else {
				fmt.Printf("User has correctly registered!\n")
				err := database.UserRegister(env, username, name, email, password)
				if err != nil {
					panic("No database found")
				}

				// .. check credentials ..
				setSession(username, name, w)
				http.Redirect(w, r, env.SiteRoot+"/game/example/play", http.StatusTemporaryRedirect)
			}

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
			log.Print("User " + username + " not found.")

			// For now, just redirect them to the home page
			http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
			return
		}

		// Populate the data needed for the page
		pagedata := getPageData(env, "sessionid", "ViewUser")
		pagedata.UserPage = models.UserPage{
			MatchesSession: true,
			Username:       user.Username,
			Name:           user.Name,
			Email:          user.Email,
			PictureSlug:    user.PictureSlug,
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
		pagedata := getPageData(env, "sessionid", "EditUser")
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
		http.Redirect(w, r, env.SiteRoot+"/game/example/play", http.StatusTemporaryRedirect)
		// Else, send them to the lobby
		//http.Redirect(w, r, env.SiteRoot+"/lobby", http.StatusTemporaryRedirect)
	})
}

func PlayGame(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Populate the data needed for the page
		pagedata := getPageData(env, "sessionid", "PlayGame")

		// Execute the template with our page data
		template := env.Templates["PlayGame"]
		template.Execute(w, pagedata)
	})
}

func ViewLobby(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var lobby models.Lobby

		for _, listing := range env.Games {
			if listing.Status == "open" {
				lobby.Games = append(lobby.Games, listing)
			}
		}

		if len(lobby.Games) > 0 {
			// Sort the games here
			lobby.Empty = false
		} else {
			lobby.Empty = true
		}

		// Populate the data needed for the page
		pagedata := getPageData(env, "sessionid", "ViewLobby")
		pagedata.Lobby = lobby

		// Execute the template with our page data
		template := env.Templates["ViewLobby"]
		template.Execute(w, pagedata)
	})
}

func WatchGame(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Populate the data needed for the page
		pagedata := getPageData(env, "sessionid", "WatchGame")

		// Execute the template with our page data
		template := env.Templates["WatchGame"]
		template.Execute(w, pagedata)
	})
}

func Leaderboard(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		leaderboard, err := database.GetLeaderboard(env)
		if err != nil {
			// Big error maybe
			log.Fatal(err)
		}

		// Populate the data needed for the page
		pagedata := getPageData(env, "sessionid", "Leaderboard")
		pagedata.Leaderboard = *leaderboard

		// Execute the template with our page data
		template := env.Templates["Leaderboard"]
		template.Execute(w, pagedata)
	})
}
