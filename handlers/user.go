package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"poker/database"
	"poker/models"
)

func User(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]
		action := vars["action"]

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
			pagedata = getPageData(env, r, []byte("sessionid"), "EditUser")

			if pagedata.Identity.Username != username {
				http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
				return
			}

			template = env.Templates["EditUser"]

			if r.Method == "POST" {
				r.ParseForm()

				user, err := database.GetUser(env, username)
				if err != nil {
					log.Print(err)
				} else {
					err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(r.PostFormValue("password")))
					if err != nil {
						// Wrong password
						log.Print(err)
					} else {
						user.Name = r.PostFormValue("name")
						user.Email = r.PostFormValue("email")
						user.Description = r.PostFormValue("description")

						if r.PostFormValue("newpassword") == r.PostFormValue("newpassword-repeat") && r.PostFormValue("newpassword") != "" {
							newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.PostFormValue("newpassword")), 8)
							if err != nil {
								log.Print(err)
							} else {
								user.HashedPassword = newHashedPassword
							}
						}

						err = database.UpdateUser(env, user)
						if err != nil {
							log.Print(err)
						}
					}
				}

				userPage, _ = database.GetUserPage(env, username)
			}
		} else {
			pagedata = getPageData(env, r, []byte("sessionid"), "ViewUser")
			template = env.Templates["ViewUser"]
		}

		pagedata.UserPage = userPage

		// Execute the template with our page data
		template.Execute(w, pagedata)
	})
}
