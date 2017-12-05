package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"poker/models"
)

func RedirectGame(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If someone is stitting at a table, send them to that table
		// http.Redirect(w, r, env.SiteRoot+"/game/example/play", http.StatusTemporaryRedirect)

		// Else, send them to the lobby
		http.Redirect(w, r, env.SiteRoot+"/lobby/", http.StatusTemporaryRedirect)
	})
}

func Game(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		gameslug := vars["gameslug"]
		action := vars["action"]

		pagedata := getPageData(env, "sessionid", "Game")
		template := env.Templates["WatchGame"] // hack

		gameListing := env.Games[gameslug]
		if gameListing == nil {
			// Game doesn't exist
			http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
			return
		}

		pagedata.GameListing = gameListing

		// Choose our template based on the action
		if action == "play" {
			template = env.Templates["PlayGame"]
		}

		// Execute the template with our page data
		template.Execute(w, pagedata)
	})
}

func WebsocketConnection(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := env.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("Couldn't upgrade.")
			log.Print(err)
			return
		}

		msg := "Connected!"

		err = conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Print("Couldn't write message.")
			log.Print(err)
		} else {
			_, _, err := conn.ReadMessage()
			if err != nil {
				log.Print("Couldn't read message.")
				log.Print(err)
			} else {
				log.Print("Reply recieved.")
			}
		}

		conn.Close()
	})
}
