package handlers

import (
	"log"
	"net/http"

	"poker/database"
	"poker/models"
)

func Leaderboard(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		leaderboard, err := database.GetLeaderboard(env)
		if err != nil {
			// Big error maybe
			log.Fatal(err)
		}

		// Populate the data needed for the page
		pagedata := getPageData(env, r, []byte("sessionid"), "Leaderboard")
		pagedata.Leaderboard = leaderboard

		// Execute the template with our page data
		template := env.Templates["Leaderboard"]
		template.Execute(w, pagedata)
	})
}
