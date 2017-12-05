package handlers

import (
	"fmt"
	// "log"   --> use this instead of fmt
	"net/http"
	"regexp"

	// "github.com/gorilla/sessions"

	"poker/database"
	"poker/models"
)

func Login(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Populate the data needed for the page
		pagedata := getPageData(env, "sessionid", "PageLogin")
		pagedata.Session.LoggedIn = false

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

					// var store = sessions.NewCookieStore([]byte(username))
					// session, err := store.Get(r, "session")
					// if err != nil {
					//     http.Error(w, err.Error(), http.StatusInternalServerError)
					//     return
					// }
					// session.Options = &sessions.Options{
					//     Path:     "/",
					//     MaxAge:   86400, // log out
					//     HttpOnly: true,
					// }
					// session.Values["username"] = username
					// session.Values["name"] = name
					// session.Save(r, w)

					// fmt.Printf("session values: ", session.Values["username"], session.Values["name"], "\n")

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

		// Populate the data needed for the page
		pagedata := getPageData(env, "sessionid", "PageUser")
		pagedata.Session.LoggedIn = false

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
					fmt.Printf("User has correctly registered!\n")
					err := database.UserRegister(env, username, name, email, password)
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
