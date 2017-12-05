package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	// "github.com/gorilla/sessions"

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
			Session: &models.Session{
				LoggedIn:  false,
				PageLogin: true,
			},
		}

		fmt.Printf("User has visited the Login page.\n")
		// the user should only be able to login if they are not logged in
		if _, err := r.Cookie("session"); err == nil { 
			fmt.Printf("The user is already logged in, so this page should not be available. Redirecting to play page.\n")
			http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
		} else {

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
				} else if CheckPasswordHash(password, userAccount.Password) != true {
					fmt.Printf("The password is incorrect.\n")
					template.Execute(w, pagedata)
				} else {
					fmt.Printf("User input was correct.\n")

					// checks the user's cookie
					// if userAccount.Password == (database.FindByUsername(env, username)).Password {
					// 	session := database.CreateSession(env, "5")
					// 	cookie := http.Cookie {
					// 		Name: name,
					// 		Value: session.Uuid,
					// 		HttpOnly: true,
					// 	}
					// 	http.SetCookie(w, &cookie)
					// } else {
					// http.Redirect(w, r, "/login", 302)
					// }

					// .. check credentials ..
					setSession(username, name, w)

					http.Redirect(w, r, env.SiteRoot+"/game/example/play", http.StatusTemporaryRedirect)
				}

			} else if r.Method == "GET" {
				fmt.Printf("GET login page\n")

				// Execute the template with our page data
				template.Execute(w, pagedata)
			}
		}
	})
}

func Logout(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("User attempted to log out.\n")
		// if a valid cookie exists
		// if _, err := r.Cookie("session"); err == nil { 
		// 	cookieValue := make(map[string]string)
		// 	if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
		// 		loginUsername = cookieValue["username"]
		// 		loginName  = cookieValue["name"]
		// 	}
		// } else {
		// 	fmt.Printf("???")
		// }
		// fmt.Printf("session values: ", usera, userb, "\n")
		clearSession(w)
		http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
	})
}

func Register(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("User visited Register page.\n")
		// the user should only be able to register if they are not logged in
		if _, err := r.Cookie("session"); err == nil { 
			fmt.Printf("The user is already logged in, so this page should not be available. Redirecting to play page.\n")
			http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
		} else {

			if r.Method == "GET" {
				// Populate the data needed for the page (these should nearly all be external functions)
				pagedata := models.PageData{
					Session: &models.Session{
						LoggedIn: false,
						PageUser: true,
					},
				}

				// Execute the template with our page data
				template := env.Templates["Register"]
				template.Execute(w, pagedata)
			} else if r.Method == "POST" {
				// Populate the data needed for the page (these should nearly all be external functions)
				pagedata := models.PageData{
					Session: &models.Session{
						LoggedIn: false,
						PageUser: true,
					},
				}

				template := env.Templates["Register"]
				isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString // only accounts with letters are permitted

				fmt.Printf("User attempted to register.\n")
				// extract form data
				r.ParseForm()
				username := r.PostFormValue("username")
				password := r.PostFormValue("password")
				name := r.PostFormValue("name")
				email := r.PostFormValue("email")
				password_repeat := r.PostFormValue("password-repeat")

				// validate user input
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
					fmt.Printf("User has typed in correct input!\n")
					password_hash, err := HashPassword(password)
					if err != nil {
						panic("Something went wrong during the password hashing process.")
					}

					err = database.UserRegister(env, username, name, email, password_hash)
					if err != nil {
						panic("No database found")
					}

					// .. check credentials ..
					setSession(username, name, w)
					http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
				}

			}
		}
	})
}

func User(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]
		action := vars["action"]

		// the user should only be able to edit pages if they are logged in
		if _, err := r.Cookie("session"); err != nil { 
			fmt.Printf("The user is NOT logged in, so this page should not be available. Redirecting to login page.\n")
			http.Redirect(w, r, env.SiteRoot+"/login/", http.StatusTemporaryRedirect)
		} else {

			// Get the user page matching that username from the database
			userPage, err := database.GetUserPage(env, username)
			if err != nil {
				log.Print("User " + username + " not found.")

				// For now, just redirect them to the home page
				http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
				return
			}

			var pagedata models.PageData
			var template *template.Template

			// Create our pagedata model
			if action == "edit" {
				pagedata = getPageData(env, "sessionid", "EditUser")
				template = env.Templates["EditUser"]

				log.Print("Editing player " + username + ".")
			} else {
				pagedata = getPageData(env, "sessionid", "ViewUser")
				template = env.Templates["ViewUser"]

				log.Print("Displaying player " + username + ".")
			}
			pagedata.UserPage = userPage

			// Execute the template with our page data
			template.Execute(w, pagedata)
		}
	})
}

func RedirectGame(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If someone is stitting at a table, send them to that table
		// http.Redirect(w, r, env.SiteRoot+"/game/example/play", http.StatusTemporaryRedirect)

		// Else, send them to the lobby
		http.Redirect(w, r, env.SiteRoot+"/lobby/", http.StatusTemporaryRedirect)
	})
}

func Game(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := r.Cookie("session"); err != nil { 
			fmt.Printf("The user is NOT logged in, so this page should not be available. Redirecting to login page.\n")
			http.Redirect(w, r, env.SiteRoot+"/login/", http.StatusTemporaryRedirect)
		} else {

			vars := mux.Vars(r)
			gameslug := vars["gameslug"]
			action := vars["action"]

			pagedata := getPageData(env, "sessionid", "Game")
			template := env.Templates["WatchGame"] // hack

			gameListing := env.Games[gameslug]
			if gameListing == nil {
				// Game doesn't exist
				http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
				return
			}

			pagedata.GameListing = gameListing

			// Choose our template based on the action
			if action == "play" {
				template = env.Templates["PlayGame"]
			}

			// Execute the template with our page data
			template.Execute(w, pagedata)
		}
	})
}

func WebsocketConnection(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := env.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("Couldn't upgrade.")
			log.Print(err)
			return
		}

		msg := "Connected!"

		err = conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Print("Couldn't write message.")
			log.Print(err)
		} else {
			_, _, err := conn.ReadMessage()
			if err != nil {
				log.Print("Couldn't read message.")
				log.Print(err)
			} else {
				log.Print("Reply recieved.")
			}
		}

		conn.Close()
	})
}

func ViewLobby(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var lobby models.Lobby

		// the user should only be able to view lobbies if they are logged in
		if _, err := r.Cookie("session"); err != nil { 
			fmt.Printf("The user is NOT logged in, so this page should not be available. Redirecting to login page.\n")
			http.Redirect(w, r, env.SiteRoot+"/login/", http.StatusTemporaryRedirect)
		} else {

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
			pagedata.Lobby = &lobby

			// Execute the template with our page data
			template := env.Templates["ViewLobby"]
			template.Execute(w, pagedata)
		}
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
		pagedata.Leaderboard = leaderboard

		// Execute the template with our page data
		template := env.Templates["Leaderboard"]
		template.Execute(w, pagedata)
	})
}
