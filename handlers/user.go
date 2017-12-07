package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"

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
			pagedata = getPageData(env, r, "sessionid", "EditUser")
			template = env.Templates["EditUser"]

			log.Print("Editing player " + username + ".")
		} else {
			pagedata = getPageData(env, r, "sessionid", "ViewUser")
			template = env.Templates["ViewUser"]

			log.Print("Displaying player " + username + ".")
		}
		pagedata.UserPage = userPage

		// Execute the template with our page data
		template.Execute(w, pagedata)

	})
}
