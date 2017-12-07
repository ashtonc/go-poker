package handlers

import (
	_ "log"
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

		pagedata := getPageData(env, r, "sessionid", "Game")
		template := env.Templates["WatchGame"]

		gameListing := env.Games[gameslug]
		if gameListing == nil {
			// Game doesn't exist
			http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
			return
		}

		pagedata.GamePage = gameListing

		// Choose our template based on the action
		if action == "play" {
			template = env.Templates["PlayGame"]
		}
		//create new game
		game := gameListing.Game
		game.Stakes.MaxBet = 20
		game.Stakes.MinBet = 1
		game.Stakes.Ante = 1
		//create players for the game
		p1err := game.Join("Ashton", 100, 0)
		if p1err != nil{
			fmt.Printf("Player 1 was not added to the game! \n")
		}
		p2err := game.Join("Adam", 100, 1)
		if p2err != nil {
			fmt.Printf("Player 2 was not added to the game! \n")
			log.Fatal(p2err)
		}
		p3err := game.Join("Matthew", 100, 2)
		if p3err != nil {
			fmt.Printf("Player 3 was not added to the game! \n")
			log.Fatal(p3err)
		}
		//Start a new round
		dealterToken := 0
		newErr := game.NewRound(dealterToken)
		if newErr != nil{
			log.Fatal(newErr)
		}
		fmt.Printf("Current Player is %s \n", game.Get_current_player_name())
		//main round loop
		for{
			err, winner := game.Winner_check()
			if err != nil{
				log.Fatal(err)
			}
			if winner != nil{
				fmt.Printf("A winner is %s \n", winner.Name)
				break
			}
			if (game.Phase == 0 || game.Phase == 2 || game.Phase == 4){
				player := game.Get_current_player_name()
				decision := rand.Float32()
				if decision < 0.25{
					fmt.Printf("Bet \n")
					raise := rand.Intn(5)
					err := game.Bet(player, raise)
					if err != nil{
						log.Fatal(err)
					}
				}else if decision < 0.88{
					pindex, err := game.GetPlayerIndex(player)
					if err != nil{
						log.Fatal(err)
					}
					if game.Players[pindex].Bet == game.Current_Bet{
						fmt.Printf("Check \n")
						err = game.Check(player)
						if err != nil{
							log.Fatal(err)
						}
					}else{
						fmt.Printf("Call \n")
						err := game.Call(player)
						if err != nil{
							log.Fatal(err)
						}
					}
				}else{
					fmt.Printf("Fold \n")
					err := game.Fold(player)
					if err != nil{
						log.Fatal(err)
					}	
				}
			}else if (game.Phase == 1 || game.Phase == 3){
				player := game.Get_current_player_name()
				num_discard := rand.Intn(4)
				fmt.Printf("%s will discard %d cards \n", player, num_discard)
				
				discard := make([]int, 0)    
				for i := 0; i <= num_discard; i++{
					discard = append(discard, i)
				}
				err := game.Discard(player, discard...)
				if err != nil{
					log.Fatal(err)
				}
			}else{
				//game.Phase == 5
				winner := game.Showdown()
				fmt.Printf("A winner is %s", winner.Name)
				break
				}
			}
		
		// Execute the template with our page data
		template.Execute(w, pagedata)})
	}





func GameAction(env *models.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		gameslug := vars["gameslug"]
		action := vars["action"]

		// Get the user session and determine whether or not they are a player in the game
		// getSession

		gameListing := env.Games[gameslug]

		if gameListing == nil {
			// Game doesn't exist
			http.Redirect(w, r, env.SiteRoot+"/", http.StatusTemporaryRedirect)
			return
		}

		game := gameListing.Game
		if game == nil {
			// Wasn't instantiated properly
		}

		// Get the seat number from their session info... abstract into another function here
		username := "username"

		if action == "sit" {
			// Tell the game the player joined, and what seat they are trying to sit in
			// If the seat is occupied, tell them to get out of here

			//game.Join(accountinfo, buyin, seatnumber)
		}

		if action == "leave" {
			// Tell the game the player left (we can figure out the seat from their session)
			// Send them to the game lobby
			game.Leave(username)
		}

		if action == "check" {
			// Tell the game they checked (we can figure out the seat from their session)
			game.Check(username)
		}

		if action == "bet" {
			// Get their bet amount
			// Tell the game they bet n amount (we can figure out the seat from their session)

			//game.Bet(username, betamount)

		}

		if action == "call" {
			// Tell the game they called (we can figure out the seat from their session)
			game.Call(username)
		}

		if action == "fold" {
			// Tell the game they folded (we can figure out the seat from their session)
			game.Fold(username)
		}

		if action == "discard" {
			// Get the indices of the cards that they discarded
			// Tell the game they discard n cards (we can figure out the seat from their session)
			game.Discard(username, 1, 3, 4)
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
