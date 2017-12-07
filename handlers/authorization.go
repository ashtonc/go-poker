package handlers

import (
	"log"
	"net/http"
	//"regexp"

	"golang.org/x/crypto/bcrypt"
	"time"

	"poker/database"
	"poker/models"
)

func Login(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Populate the data needed for the page
		pagedata := getPageData(env, r, "sessionid", "PageLogin")
		template := env.Templates["Login"]

		// If the user is already logged in, send them to the home page
		if pagedata.Identity.LoggedIn == true {
			http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
			return
		}

		if r.Method == "POST" {
			r.ParseForm()
			username := r.PostFormValue("username")

			user, err := database.GetUser(env, username)
			if err != nil {
				// User doesn't exist
				log.Print(err)
				template.Execute(w, pagedata)
				return
			}

			err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(r.PostFormValue("password")))
			if err != nil {
				// Wrong password
				log.Print(err)
				template.Execute(w, pagedata)
				return
			}

			token, err := database.CreateSession(env, username, time.Now().AddDate(0, 1, 0))
			if err != nil {
				// Couldn't create a session
				log.Print(err)
				template.Execute(w, pagedata)
				return
			}

			err = setSessionToken(env, w, r, token)
			if err != nil {
				// Couldn't set the session token
				log.Print(err)
				template.Execute(w, pagedata)
				return
			}

			http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
			return
		}

		template.Execute(w, pagedata)

		/*
			if r.Method == "POST" {
				fmt.Printf("POST for Login\n")
				// this code gets username and password from the POST form
				userAccount := database.FindByUsername(env, username)
				//name := userAccount.Name

				// Query the database to check if account even exists (via FindByUsername?)
				if username == "" || password == "" {
					fmt.Printf("One or more fields were left blank.\n")
					template.Execute(w, pagedata)
				} else if userAccount.Username != username {
					fmt.Printf("This user does not exist.\n")
					template.Execute(w, pagedata)
				} else if CheckPasswordHash(password, userAccount.HashedPassword) != true {
					fmt.Printf("The password is incorrect.\n")
					template.Execute(w, pagedata)
				} else {
					fmt.Printf("Login successful.\n")

					// .. check credentials ..
						setSession(username, name, w)
						sessions.CreateSession(env, username)

					http.Redirect(w, r, env.SiteRoot+"/game/example/play", http.StatusTemporaryRedirect)
				}
			}
		*/
	})
}

func Logout(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clearSession(w)
		http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
	})
}

func Register(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Populate the data needed for the page
		pagedata := getPageData(env, r, "sessionid", "PageLogin")
		template := env.Templates["Register"]

		// If the user is already logged in, send them to the home page
		if pagedata.Identity.LoggedIn == true {
			http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
			return
		}

		if r.Method == "POST" {
			var user models.User

			r.ParseForm()

			user.Username = r.PostFormValue("username")
			user.Name = r.PostFormValue("name")
			user.Email = r.PostFormValue("email")
			user.PictureSlug = "temp.png"
			user.Description = "temp"

			if r.PostFormValue("password") == r.PostFormValue("password-repeat") {
				user.HashedPassword = bcrypt.GenerateFromPassword([]byte(r.PostFormValue("password")), 16)
			}

			err := database.AddUser(env, user)
			if err != nil {
				log.Print(err)
				// Couldn't be added(?)
			}

			token, err := database.CreateSession(env, username, time.Now().AddDate(0, 1, 0))
			if err != nil {
				// Couldn't create a session
				log.Print(err)
				template.Execute(w, pagedata)
				return
			}

			err = setSessionToken(env, w, r, token)
			if err != nil {
				// Couldn't set the session token
				log.Print(err)
				template.Execute(w, pagedata)
				return
			}

			http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
			return
		}

		template.Execute(w, pagedata)

		/*
			// the user should only be able to register if they are not logged in
			if _, err := r.Cookie("session"); err == nil {
				fmt.Printf("The user is already logged in, so this page should not be available. Redirecting to play page.\n")
				http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
			} else {

				if r.Method == "GET" {

					// Execute the template with our page data
					template := env.Templates["Register"]
					template.Execute(w, pagedata)
				} else if r.Method == "POST" {
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
						fmt.Printf("User has correctly registered!\n")
						password_hash, _ := HashPassword(password)
						err := database.UserRegister(env, username, name, email, password_hash)
						if err != nil {
							panic("No database found")
						}
						// .. check credentials ..
							setSession(username, name, w)
							sessions.CreateSession(env, username)

						http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
					}

				}
			}
		*/
	})
}
