package handlers

import (
	"fmt"
	// "log"   --> use this instead of fmt
	"net/http"
	"regexp"

	"poker/database"
	"poker/models"
)

func Login(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Populate the data needed for the page
		pagedata := getPageData(env, r, "sessionid", "PageLogin")
		pagedata.Identity.LoggedIn = false

		template := env.Templates["Login"]

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
					/*
						setSession(username, name, w)
						sessions.CreateSession(env, username)
					*/

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

		clearSession(w)

		http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
	})
}

func Register(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("User visited Register page.\n")

		// Populate the data needed for the page
		pagedata := getPageData(env, r, "sessionid", "PageUser")
		pagedata.Identity.LoggedIn = false

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
					/*
						setSession(username, name, w)
						sessions.CreateSession(env, username)
					*/

					http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
				}

			}
		}
	})
}
