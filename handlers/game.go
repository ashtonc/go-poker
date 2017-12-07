package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"poker/connection"
	"poker/models"
	//for test
	"fmt"
	"log"
	//"math/rand"

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

		pagedata := getPageData(env, r, []byte("sessionid"), "Game")

		gameListing := env.Games[gameslug]
		if gameListing == nil {
			// Game doesn't exist
			http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
			return
		}

		pagedata.GamePage = gameListing

		/* ****************** */
		/* Put game code here */
		/* ****************** */

		game := gameListing.Game
		game.Stakes.MaxBet = 20
		game.Stakes.MinBet = 0
		game.Stakes.Ante = 1
		p1err := game.Join("Ashton", "Ashton", " ", 100, 0)
		if p1err != nil {
			fmt.Printf("Player 1 was not added to the game! \n")
		}
		p2err := game.Join("Adam", "Adam",  " ", 100, 1)
		if p2err != nil {
			fmt.Printf("Player 2 was not added to the game! \n")
			log.Print(p2err)
		}
		p3err := game.Join("Matthew", "Matthew", " ", 100, 2)
		if p3err != nil {
			fmt.Printf("Player 3 was not added to the game! \n")
			log.Print(p3err)
		//Start a new round

		/* ******************* */
		/* Stop game code here */
		/* ******************* */

		// Choose our template based on the action
		template := env.Templates["WatchGame"]
		if action == "play" {
			template = env.Templates["PlayGame"]
		}
}
		// Execute the template with our page data
		template.Execute(w, pagedata)
	})
}

func GameAction(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		gameslug := vars["gameslug"]
		action := vars["action"]

		// Get the user session and determine whether or not they are a player in the game
		pagedata := getPageData(env, r, []byte("sessionid"), "Game")
		username := pagedata.Identity.Username

		gameListing := env.Games[gameslug]
		if gameListing == nil {
			// Game doesn't exist
			http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
			return
		}

		game := gameListing.Game
		if game == nil {
			// Wasn't instantiated properly
			http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
			return
		}

		if action == "sit" {
			// Tell the game the player joined, and what seat they are trying to sit in
			// If the seat is occupied, tell them to get out of here

			game.Join(pagedata.Identity.Name, username, "img", 200, 2)
		}

		if action == "leave" {
			// Tell the game the player left
			// Send them to the game lobby
			game.Leave(username)
		}

		if action == "check" {
			// Tell the game they checked
			game.Check(username)
		}

		if action == "bet" {
			// Get their bet amount
			// Tell the game they bet n amount

			//game.Bet(username, betamount)

		}

		if action == "call" {
			// Tell the game they called
			game.Call(username)
		}

		if action == "fold" {
			// Tell the game they folded
			game.Fold(username)
		}

		if action == "discard" {
			// Get the indices of the cards that they discarded
			// Tell the game they discarded n cards
			game.Discard(username, 1, 3, 4)
		}

		if action == "start_round"{
			fmt.Printf("Start round \n")
			dealterToken := 0
			newErr := game.NewRound(dealterToken)
			if newErr != nil {
				log.Print(newErr)
				}
		}

		// Have the default here (back to game)
		http.Redirect(w, r, env.SiteRoot+"/game/"+gameslug+"/play", http.StatusTemporaryRedirect)
	})
}

func WebsocketConnection(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//vars := mux.Vars(r)
		//gameslug := vars["gameslug"]

		// Choose the correct hub based on the session of the user
		hub := connection.NewHub()

		// Get the user id from their session
		hub.HandleWebSocket(env, w, r)
	})
}
