package handlers

import (
	"fmt"
	// "log" -> Use this instead of fmt
	"net/http"
	"regexp"

	"poker/database"
	"poker/models"
)

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

func Logout(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("User attempted to log out.\n")
		clearSession(w)
		http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
	})
}

func Register(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("User visited Register page.\n")

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

			pagedata := models.PageData{
				Session: &models.Session{
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
