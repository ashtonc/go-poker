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
		pagedata := getPageData(env, r, []byte("sessionid"), "PageLogin")
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
	})
}

func Logout(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := getSessionToken(env, r)
		if err != nil {
			log.Print(err)
		}

		clearSession(w)

		err = database.RevokeSession(env, token)
		if err != nil {
			log.Print(err)
		}

		http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
	})
}

func Register(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Populate the data needed for the page
		pagedata := getPageData(env, r, []byte("sessionid"), "PageLogin")
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
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.PostFormValue("password")), 16)
				if err != nil {
					// problem
				}

				user.HashedPassword = hashedPassword
			}

			err := database.AddUser(env, &user)
			if err != nil {
				// Couldn't be added(?)
				log.Print(err)
			}

			token, err := database.CreateSession(env, user.Username, time.Now().AddDate(0, 1, 0))
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
					isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString // only accounts with letters are permitted

					// validate user input
					if len(username) < 5 || !isAlpha(username) {
						template.Execute(w, pagedata)
						fmt.Printf("Incorrect input for username.\n")
					}

					if len(name) < 1 || !isAlpha(name) {
						template.Execute(w, pagedata)
						fmt.Printf("Incorrect input for name.\n")
					}

					if len(email) < 1 {
						template.Execute(w, pagedata)
						fmt.Printf("Incorrect input for email.\n")
					}

					if len(password) < 6 {
						template.Execute(w, pagedata)
						fmt.Printf("Incorrect input for password.\n")
					}

					if password != password_repeat {
						template.Execute(w, pagedata)
						fmt.Printf("The password field should match the password repeat field.\n")
					}

					if (database.FindByUsername(env, username)).Username == username {
						template.Execute(w, pagedata)
						fmt.Printf("This account name already exists.\n")
					}


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
