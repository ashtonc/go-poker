package handlers

import (
	"net/http"

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
		pagedata := getPageData(env, r, []byte("sessionid"), "Home")

		// Execute the template with our page data
		template := env.Templates["Home"]
		template.Execute(w, pagedata)
	})
}
