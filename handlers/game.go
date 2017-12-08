package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"poker/connection"
	"poker/models"
	"time"
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

		// game := gameListing.game
		// do some game tests...

		/* ******************* */
		/* Stop game code here */
		/* ******************* */

		// Choose our template based on the action
		template := env.Templates["WatchGame"]
		if action == "play" {
			template = env.Templates["PlayGame"]
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

		if r.Method == "POST" {
			r.ParseForm()

			if action == "sit" {
				// Tell the game the player joined, and what seat they are trying to sit in
				// If the seat is occupied, tell them to get out of here

				seat, _ := strconv.Atoi(r.PostFormValue("seat"))
				buyin, _ := strconv.Atoi(r.PostFormValue("buyin"))
				seat--

				log.Print("Game " + gameslug + ": " + username + " joined seat " + r.PostFormValue("seat") + " with a buyin of " + r.PostFormValue("buyin"))

				if seat <= 5 && seat >= 0 && buyin > 0 {
					game.Join(pagedata.Identity.Name, username, "img.png", buyin, seat)
				}
			}

			if action == "bet" {
				// Get their bet amount
				// Tell the game they bet n amount
				bet, _ := strconv.Atoi(r.PostFormValue("bet"))

				game.Bet(username, bet)
				if game.Phase == 5 {
					winner := game.Showdown()
					log.Print(winner)
					game.Seats[winner.Seat].Winner = true
					game.Dealer_Token +=1
					<-time.After(8 * time.Second)
						log.Print("New round...")
    					go game.NewRound(game.Dealer_Token)

				}

			}

			if action == "discard" {
				// Get the indices of the cards that they discarded
				// Tell the game they discarded n cards

				var discarded []int

				if r.PostFormValue("card1discard") != "" {
					discarded = append(discarded, 0)
				}

				if r.PostFormValue("card2discard") != "" {
					discarded = append(discarded, 1)
				}

				if r.PostFormValue("card3discard") != "" {
					discarded = append(discarded, 2)
				}

				if r.PostFormValue("card4discard") != "" {
					discarded = append(discarded, 3)
				}

				if r.PostFormValue("card5discard") != "" {
					discarded = append(discarded, 4)
				}

				game.Discard(username, discarded...)
			}
		}

		if action == "call" {
			// Tell the game they called
			game.Call(username)
			if game.Phase == 5{
					winner := game.Showdown()
					log.Print(winner)
					game.Seats[winner.Seat].Winner = true
					game.Dealer_Token +=1
					<-time.After(8 * time.Second)
						log.Print("New round...")
    					go game.NewRound(game.Dealer_Token)
					////// game.EndRound()
				}
		}

		if action == "leave" {
			// Tell the game the player left
			// Send them to the game lobby
			game.Leave(username)
		}

		if action == "check" {
			// Tell the game they checked
			game.Check(username)
			if game.Phase == 5{
					winner := game.Showdown()
					log.Print(winner)
					game.Seats[winner.Seat].Winner = true
					game.Dealer_Token +=1
					<-time.After(8 * time.Second)
						log.Print("New round...")
    					go game.NewRound(game.Dealer_Token)
					////// game.EndRound()
				}
		}

		if action == "fold" {
			// Tell the game they folded
			game.Fold(username)
			_, winner := game.Winner_check()
			if winner != nil {
				log.Print(winner)
				game.Seats[winner.Seat].Winner = true
				game.Dealer_Token +=1
				<-time.After(8 * time.Second)
					go game.NewRound(game.Dealer_Token)


				///////game.EndRound()
			}
		}

		if action == "start" {
			// Start the game I suppose
			log.Print("Start round \n")
			_ = game.NewRound(game.Dealer_Token)
		}

		// Have the default here (back to game)
		http.Redirect(w, r, env.SiteRoot+"/game/"+gameslug+"/play", http.StatusTemporaryRedirect)
	})
}

func WebsocketConnection(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Choose the correct hub based on the session of the user, not just a new one..
		hub := connection.NewHub()

		// Get the user id from their session
		hub.HandleWebSocket(env, w, r)
	})
}
