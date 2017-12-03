
/*(only class attributes needed for the game logic will be included here) */
package gamelogic

import(
	"math/rand"
	"time"
	"bufio"
  	"fmt"
  	"os"
  	"strings"
  	//"sort"
  	"strconv"
  	"errors"
)

func (g *Game)SetTimer(time int){
  g.Timer = time.NewTimer(time.Second*time)
}


func (g *Game)GetPlayerIndex(name string)(int, error){
	for i, p := range(g.Players){
		if p.Name == name{
			return i, nil 
			}
		}
	return -1, error.New("There is no player of that name") 
	}

func (g *Game)GetSitterIndex(name string)(int, error){
	for i, p := range(g.Sitter){
		if p.Name == name{
			return i, nil 
			}
		}
	return 0, error.New("There is no sitter of that name") 
	}

func (g *Game) Join(name string, buyin int, seatNumber int)error{
	if g.Seat[seatNumber].occupied == true {
		return error.New("Seat is already occupied")
	}
	if buyin > g.Ant*100{
		return error.New("Buyin exceeds the limit for this table")
	}
	if buyin < g.Ante+50{
		return error.New("Buyin is too low for this table")
	}
	player = new(Player)
	player.Name = name
	player.Money = buyin
	player.Seat = seatNumber
	player.Discard = false
	player.Folded = false
	player.Called = false
	g.Sitters = append(g.Players, *player)
	g.Seat[seatNumber].occupied = true
	return nil
}


func (g *Game) Leave(name string)error{
	index = g.GetSitterIndex(name)
	g.Sitters = append(g.Sitters[:index], g.Sitters[index+1:])
	p = g.Sitters[index]
	if g.Seats[p.Seat].occupied == false{
		return error.New("Seat already empty")
	}
	g.Seats[p.Seat].occupied = false
	if len(g.Players) > 0{
		index = g.GetPlayerIndex(name)
		g.Players[] = append(g.Players[:index], g.Sitters[index+1])
	}
	return nil
}
 
func (g *Game) NewRound(players []Player, ante int, minBet int, maxBet int, dealterToken int)error{
	for _, p in g.Sitters{
		g.Players = append(g.Players, *p)
	}
	if len(g.Players) < 2{
		return error.New("A round of Poker requires at least two players")
	}	
 	g.Phase = 1 
 	g.Bet_Counter = -1
 	g.GameStakes.MaxBet = maxBet
 	g.GameStakes.MinBet = minBet
 	g.Dealer_Token = dealterToken + 1 
 	g.Current_player = g.Players[dealterToken]
 	cardTypes, suites := Init_card_cat()
 	g.Pot = 0
 	// ant played by each player
 	for i := 0; i < len(g.Players); i++{
 		g.players[i].Money -= ante
 		g.Pot += ante
 	}
 	//Create deck, shuffle cards, deal cards to players
 	g.Deck = createDeck(cardTypes, suites)
 	g.Deck = shuffle(deck)
	d := 0
	 	for d < 5{
	 		for i := 0; i < len(players); i++{
	 			card := draw(g.Deck)
	 			g.Deck = g.Deck[1:]
	 			g.Players[i].Hand = append(g.players[i].Hand, card)
	 			fmt.Printf(" %s is delt a %s of %s \n ", g.Players[i].Name, card.Face, card.Suit)
	 		}
	 		d++
	 	}
	 g.SetTimer(11)
	 return nil
	}

// Eachetting round lasts until each player has either: (a) folded (b) called 
func (g *Game)Bet(p_name string, bet int)error{
	pindex = g.GetPlayerIndex(p_name)
	if g.Players[pindex].Folded == true{
		return erro.New("Player has already folded and so cannot bet")
	}
	if bet > g.Players[pindex].Money{
		return error.New("bet exceeds player's money")
	}
	if bet > g.GameStakes.MaxBet{
		return error.New("Bet exceeds maximum bet")
	}
	if bet < g.GameStakes.MinBet{
		return error.New("Bet is below the minimum bet")
	}
	if bet <= g.Current_bet{
		return error.New("New bet must be geater than the current bet")
	}	
	g.Players[pindex].Money -= bet
	g.Pot += bet
	g.Current_bet = bet
	g.ResetBetCounter()
	g.Current_player = g.Next_Player()
	if g.check_if_betting_ends(){
		g.Phase += 1
	}

	return nil
	}
}

func (g *Game)ResetBetCounter(){
	players_in := 0
	for i := range(g.Players){
		if g.Players[i].Folded == false{
			players_in += 1
		}
	g.Bet_Counter = players_in
	}
}

func (g *Game)Call(p_name string)error{
	pindex = g.GetPlayerIndex(p_name)
	if g.Players[pindex].Folded == true{
		return error.New("Player has already folded and cannot call.")
	}
	if g.Current_bet > g.Players[pindex].Money{
		return error.New("bet exceeds player's money")
	}
	g.Players[pindex].Money -= g.Current_bet
	g.Players[pindex].Called = true
	g.Pot += g.Current_bet
	g.Current_player = g.Next_Player()
	g.Bet_Counter -= 1
	if g.Bet_Counter == 0{
		g.Phase +=1
	}
	g.Timer := time.NewTimer(time.Second * 10)
	return nil
}

func (g *Game)Next_Player()string{
	pindex = g.GetPlayerIndex(g.Current_player)
	current = g.Players[pindex]
	for i := range g.Players{
		if i == len(g.Players) - 1 {
			return g.Players[0].Name
		}else{
			if g.Players[i].Name = current.Name{
			return g.Players[i+1].Name			
			}
		}
	}
	g.SetTimer(10)	
	return nil
}

func (g *Game)Next_Phase()error{
	if g.Phase > 4{
		return error.New("The game phase cannot be further incremented \n")
	}
	g.Phase +=1
	g.Current_player = g.Players[g.Dealer_Token]
}

func (g *Game)Fold(player_name string)error{
	pindex = g.GetPlayerIndex(player_name)
	if g.Players[pindex].Folded == true {
		return error.New("Player has already folded")
	}else{
		g.Players[pindex].Folded = true
		g.Current_player = g.Next_Player()
		g.Bet_Counter -= 1
	}
	if g.Bet_Counter == 0{
		g.Next_Phase()
	}
	return nil
}

func (g *Game)Check(player_name, string)error{
	pindex = g.GetPlayerIndex(player_name)
	if g.Current_bet > g.Players[pindex].Bet{
		return error.New("Player's cannot check unless her current bet is equal to the current bet of the game")
	}
	if g.Bet_Counter > 0{
		g.Bet_Counter -= 1
	}
	if g.Bet_Counter == 0{
		g.Next_Phase()
	}
	return nil
}

//Call when (phase == 2 or phase == 4) and g.Current_player.seat == g.Dealer_Token)
func (g *Game)check_if_winner()string{
	remaining := check_num_players_remaining(g.Players)
	if remaining == 1{
		winner := find_winner(g.players)
		return winner
	}
	return nil
}

func check_num_players_remaining(players []Player) int {
	remaining := 0
	for _, p := range players {
		if p.Folded == false {
			remaining++
		}
	}
	return remaining
}

func find_winner(players []Player)string{
	//function assumes only one players remains in the game
	for _, p := range players {
		if p.Folded == false {
			return p.Name
		}
	}
	p := players[0] //just to make go happy
	return p.Name
}

func (g *Game)check_for_winner(){
	remaining := check_num_players_remaining(g.Players)
	if remaining < 2 {
		winner := find_winner(players)
		return winner
	}
}


func getHighestInt(array []int){
	highest := 0
	for i, v := range(array){
		if v > highest{
			highest = v
		}
	}
	return highest
}


func (g *Game)Discard(playerID string, cardIndexes []int) error{
	check = getHighestInt(cardIndexes)
	pindex = g.GetPlayerIndex(playerID)
	if g.Players[pindex].Discarded == true{
		for _, p := range(g.Players){
			p.Players[pindex].Discarded = false
			g.Next_Phase()
			return nil
		}
		g.Phase += 1
		g.Current_player = g.Players[Dealer_Token]
		return nil 
	}
	if check > len(p.Hand){
		return error.New("Index is out of range for player's hand")
	}
	g.Players[pindex].Card_Discard(card_indexes []int)
	g.Players[pindex].Redraw()
	g.Players[pindex].Discarded = true
	g.Next_Player()
	return nil
}



func (g *Game)Card_Discard(player string, card_indexes []int)error{
	pindex := g.GetPlayerIndex(player_name)
	var temp_hand []Card
	for i, n := range card_indexes{
		if n > len(p.Hand){
			return error.New("Index out of range of player's hand")
		}
	}
	for _, t := range card_indexes {
		//temp_hand[i] = hand[t]
		temp_hand = append(temp_hand, g.Players[pindex].Hand[t])
	}
	var discarded_hand []Card
	index := 0
	for _, c := range hand{
		 	check := true
		for _, d := range temp_hand{
			if c.Face == d.Face && c.Suit == d.Suit{
				check = false
			}
		}
		if check == true{
			discarded_hand = append(discarded_hand, c)
			index ++
			}
		}
	g.Players[pindex].Hand = discarded_hand
	return nil
	}

func (g *Game)Redraw(player string){
	pindex := g.GetPlayerIndex(player_name)
	replace := 5 - len(g.Players[pindex].Hand)
	for j := 0; j < replace; j++ {
		card := draw(g.Deck)
		g.Deck = deck[1:]
		fmt.Printf("%s draws a %s of %s \n", g.Players[pindex].Name, card.Face, card.Suit)
		g.Players[pindex].Hand = append(g.Players[pindex].Hand, card)
		}
	}










