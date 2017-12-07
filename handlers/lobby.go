package handlers

import (
	"net/http"

	"poker/models"
)

func ViewLobby(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var lobby models.Lobby

		for _, listing := range env.Games {
			if listing.Status == "open" {
				lobby.Games = append(lobby.Games, listing)
			}
		}

		if len(lobby.Games) > 0 {
			sort.Slice(lobby.Games, func(i, j int) bool {
				return lobby.Games[i].NameAxis < lobby.Games[j].Name
			})

			lobby.Empty = false
		} else {
			lobby.Empty = true
		}

		// Populate the data needed for the page
		pagedata := getPageData(env, r, "sessionid", "Lobby")
		pagedata.Lobby = &lobby

		// Execute the template with our page data
		template := env.Templates["Lobby"]
		template.Execute(w, pagedata)
	})
}
