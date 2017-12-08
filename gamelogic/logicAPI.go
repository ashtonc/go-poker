package gamelogic

import (
	_ "bufio"
	"errors"
	_ "os"
	_ "time"
	"time"
	"log"
	"fmt"
)

func (g *Game) Join(username string, name string, pictureslug string, buyin int, seatNumber int) error {
	//log.Print("A player has joined \n")
	log.Print(g.Stakes.Ante)
	if g.Seats[seatNumber].Occupied == true {
		log.Print("Seat Occupied")
		return errors.New("Seat is already Occupied")
	}
	if buyin > g.Stakes.Ante*100 {
		log.Print("Buyin too high")
		return errors.New("Buyin exceeds the limit for this table")
	}
	if buyin < g.Stakes.Ante*50 {
		log.Print("Buyin too low")
		return errors.New("Buyin is too low for this table")
	}
	log.Print("A player has joined \n")
	player := new(Player)
	player.Name = name
	log.Print(player.Name)
	log.Print(player.Username)
	player.Username = username
	player.PictureSlug = pictureslug
	player.Money = buyin
	player.Seat = seatNumber
	player.Discarded = false
	player.Folded = false
	player.Called = false
	g.Sitters = append(g.Sitters, player)
	g.Seats[seatNumber].Occupied = true
	g.Seats[seatNumber].Occupier = player
	g.Seats[seatNumber].Winner = false
	fmt.Printf("Currently %d players sitting \n", len(g.Sitters))
	return nil
}

func (g *Game) Leave(name string) error {
	index, error := g.GetSitterIndex(name)
	if error == nil {
		g.Sitters = append(g.Sitters[:index], g.Sitters[index+1:]...)
		if g.Seats[g.Sitters[index].Seat].Occupied == false {
			return errors.New("Seat already empty")
		}
		g.Seats[g.Sitters[index].Seat].Occupied = false
		if len(g.Players) > 0 {
			index, err := g.GetPlayerIndex(name)
			if err == nil {
				g.Players = append(g.Players[:index], g.Sitters[index+1])
			}
		}
	}
	return nil
}

func (g *Game) NewRound(dealterToken int) error {
	Rand_init()
	for _, p := range g.Sitters {
		g.Players = append(g.Players, p)
		fmt.Printf("%s added to g.Players \n", p.Name)
	}
	fmt.Printf(" Number of players: %d \n", len(g.Players))
	if len(g.Players) < 2 {
		log.Print("Not enough players")
		return errors.New("A round of Poker requires at least two players")
	}
	dealterToken += 1
	if dealterToken < 0{
		dealterToken = 0
	}
	g.Phase = 0
	g.Bet_Counter = -1
	g.Current_Bet = 0
	g.Dealer_Token = dealterToken
	log.Print(dealterToken)
	g.Current_Player = g.Players[dealterToken].Name
	fmt.Print("Current Player is: %s", g.Current_Player)
	cardTypes, suites := Init_card_cat()
	g.Pot = 0
	g.ResetBetCounter()
	// ant played by each player
	fmt.Printf("Number of players: %d \n", len(g.Players))
	fmt.Printf("Ante is: %d \n", g.Stakes.Ante)

	for i := 0; i < len(g.Players); i++ {
		g.Players[i].Money = g.Players[i].Money - g.Stakes.Ante
		g.Pot = g.Pot + g.Stakes.Ante
	}

	fmt.Printf("After ante, pot is: %d \n", g.Pot)
	//Create deck, shuffle cards, deal cards to players
	g.Deck = createDeck(cardTypes, suites)
	g.Deck = shuffle(g.Deck)
	d := 0
	for d < 5 {
		for i := 0; i < len(g.Players); i++ {
			card := draw(g.Deck)
			g.Deck = g.Deck[1:]
			g.Players[i].Hand = append(g.Players[i].Hand, card)
			//fmt.Printf(" Player %d is %s \n", i, g.Players[i].Name)
			fmt.Printf(" %s is delt a %s of %s \n ", g.Players[i].Name, card.Face, card.Suit)
		}
		d++
	}
	return nil
}

// Eachetting round lasts until each player has either: (a) folded (b) called
func (g *Game) Bet(p_name string, bet int) error {
	log.Print("Bet is being made")
	pindex, err := g.GetPlayerIndex(p_name)
	if err != nil{
		log.Print("Error!!!")
		log.Print(err)
		return errors.New("No player of that name!")
	}
	if g.Phase != 0 && g.Phase != 2 && g.Phase != 4 {
		log.Print("Game is not in betting phase")
		return errors.New("Game is not in a betting phase!")
	}
	fmt.Printf("P_name: %s \n", p_name)
	current, _ := g.GetPlayerIndex(g.Current_Player)

	fmt.Printf("Current player: %s \n", g.Current_Player)
	fmt.Printf("Sanity check: %s \n", g.Players[current].Name)
	if p_name != g.Current_Player {
		log.Print("Not current Player!!")
		return nil
	}
	if err == nil {
		if g.Players[pindex].Folded == true {
			log.Print("folded")
			return errors.New("Player has already folded and so cannot bet")
		}
		if bet > g.Players[pindex].Money {
			log.Print("No money")
			return errors.New("bet exceeds player's money")
		}
		if bet > g.Stakes.MaxBet {
			log.Print("Bet Too High")
			return errors.New("Bet exceeds maximum bet")
		}
		if bet < g.Stakes.MinBet {
			log.Print("Bet Too low")
			return errors.New("Bet is below the minimum bet")
		}
		balance := g.Current_Bet - g.Players[pindex].Bet
		log.Print(bet)
		fmt.Printf("%s owes a balance of %d and raises by %d \n", g.Players[pindex].Name, balance, bet)
		g.Players[pindex].Money -= (bet + balance)
		g.Pot += bet + (balance)
		log.Print(g.Pot)
		fmt.Printf("The pot is currently %d \n", g.Pot)
		g.Current_Bet += bet
		g.Players[pindex].Bet = g.Current_Bet
		g.ResetBetCounter()
		g.Current_Player = g.Next_Player()
		//g.check_if_betting_ends()
	}

	return nil
}

func (g *Game) Call(p_name string) error {

	if g.Phase != 0 && g.Phase != 2 && g.Phase != 4 {
		return errors.New("Game is not in a betting phase!")
	}
	if p_name != g.Current_Player {
		return nil
	}
	pindex, err := g.GetPlayerIndex(p_name)
	if err == nil {
		if g.Players[pindex].Folded == true {
			return errors.New("Player has already folded and cannot call.")
		}
		if g.Current_Bet > g.Players[pindex].Money {
			log.Print("Bet too high")
			return errors.New("bet exceeds player's money")
		}
		balance := g.Current_Bet - g.Players[pindex].Bet
		g.Players[pindex].Money -= balance
		fmt.Printf("Current game bet: %d, current bet of %s: %d, pays %d \n",
			g.Current_Bet, g.Players[pindex].Name, g.Players[pindex].Bet, balance)
		g.Players[pindex].Bet = g.Current_Bet
		g.Players[pindex].Called = true
		g.Pot += balance
		fmt.Printf("Bet counter before: %d \n", g.Bet_Counter)
		g.Bet_Counter -= 1
		if g.Bet_Counter < 0{
			g.Bet_Counter = 0
		}
		fmt.Printf("Bet counter after: %d \n", g.Bet_Counter)
		g.Current_Player = g.Next_Player()
		if g.Bet_Counter == 0 {
			//err := g.Check_if_end_of_bet_correct()
			//if err != nil {
			//	return err
			//}
			g.reset_bets()
			g.Next_Phase()
			fmt.Printf("Phase has been incremented to %d because the bet counter is zero \n", g.Phase)
		}
	} else {
		return err
	}
	return nil
}

func (g *Game) Fold(player_name string) error {
	if g.Phase != 0 && g.Phase != 2 && g.Phase != 4 {
		return errors.New("Game is not in a betting phase!")
	}
	if player_name != g.Current_Player {
		return nil
	}
	pindex, err := g.GetPlayerIndex(player_name)
	if err == nil {
		if g.Players[pindex].Folded == true {
			return errors.New("Player has already folded")
		} else {
			g.Players[pindex].Folded = true
			g.Current_Player = g.Next_Player()
			g.Bet_Counter -= 1
		}
		if g.Bet_Counter == 0 {
			err := g.Check_if_end_of_bet_correct()
			if err != nil {
				return err
			}
			g.Next_Phase()
			g.reset_bets()
			fmt.Printf("The game phase has been incremented to %d because the bet counter is 0 \n", g.Phase)
		}
	}
	return nil
}

func (g *Game) Check(player_name string) error {
	if g.Phase != 0 && g.Phase != 2 && g.Phase != 4 {
		return errors.New("Game is not in a betting phase!")
	}
	if player_name != g.Current_Player {
		return nil
	}
	pindex, err := g.GetPlayerIndex(player_name)
	if err == nil {
		if g.Current_Bet > g.Players[pindex].Bet {
			return errors.New("Player's cannot check unless her current bet is equal to the current bet of the game")
		}
		if g.Bet_Counter > 0 {
			g.Bet_Counter -= 1
		}
		if g.Bet_Counter == 0 {
			err := g.Check_if_end_of_bet_correct()
			if err != nil {
				return err
			}
			fmt.Printf("The game phase has been incremented to %d because the bet counter went to 0\n", g.Phase)
			g.Next_Phase()
			g.reset_bets()
			return nil
		}
		g.Current_Player = g.Next_Player()

	}
	return nil
}

func (g *Game) Discard(playerID string, cardIndexes ...int) error {
	log.Print("Discard being made")
	log.Print(playerID)
	if g.Phase != 1 && g.Phase != 3 {
		log.Print("Not in discard phase")
		return errors.New("Game is not in discard phase!")
	}
	if playerID != g.Current_Player {
		log.Print("274")
		return nil
	}
	check := getHighestInt(cardIndexes)
	pindex, err := g.GetPlayerIndex(playerID)
	if err == nil {
		log.Print("278")
		if g.Players[pindex].Discarded == true {
			for i := range g.Players {
				g.Players[i].Discarded = false
			}
			g.Next_Phase()
			fmt.Printf("The game phase has been incremented to %d because all players have discarded \n", g.Phase)
			return nil

			//g.Phase += 1
			//g.Current_Player = g.Players[g.Dealer_Token].Name
			//return nil
		log.Print("290")
		} else if check > len(g.Players[pindex].Hand) {
			log.Print("292")
			return errors.New("Index is out of range for player's hand")
		} else {
			discard := g.Card_Discard(g.Players[pindex].Name, cardIndexes)
			if discard == nil {
				log.Print("297")
				g.Redraw(g.Players[pindex].Name)
				g.Players[pindex].Discarded = true
				g.Current_Player = g.Next_Player()
				if g.check_if_discard_phase_complete() == true {
					fmt.Printf("The game phase has been incremented to %d because all players have discarded \n", g.Phase)
					g.Next_Phase()
				}
				return nil
			}
		}
	}
	return err
}

func (g *Game) Showdown() *Player {
	for i := 0; i < len(g.Players); i++ {
		if g.Players[i].Folded == false {
			g.Players[i].sort_hand_by_rank()
			g.Players[i].card_histogram()
		}
	}
	score_board := g.rank_hands()
	winner := g.DetermineWinner(score_board)
	//fmt.Printf("%s win a pot worth %d \n", winner.Name, pot)
	winner.Money += g.Pot
	return winner
}

func (game *Game) EndRound(){
		winner := game.Showdown()
		log.Print(winner)
		game.Seats[winner.Seat].Winner = true
		for i := range game.Players{
			var empty_hand []Card
			game.Players[i].Hand = empty_hand

		}
		<-time.After(8 * time.Second)
			log.Print("New round...")
			go game.NewRound(game.Dealer_Token)
}

