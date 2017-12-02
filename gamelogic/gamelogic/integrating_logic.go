
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

func (g *Game) SetTimer(time int){
  g.Timer = time.NewTimer(time.Second * time)
}


func (g *Game)GetPlayerIndex(name string)(int, error){
	for i, p := range(g.Players){
		if p.Name == name{
			return i, nil 
			}
		}
	return 0, error.New("There is no player of that name") 
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
	g.Sitters = append(g.Players, player)
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
 	g.GameStakes.MaxBet = maxBet
 	g.GameStakes.MinBet = minBet
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

func (g *Game)Bet(p_name string, bet int)error{
	pindex = g.GetPlayerIndex(p_name)
	p = g.Players[pindex]
	if bet > p.Money{
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
	p.Money -= bet
	g.Pot += bet
	g.Current_bet = bet
	g.Current_player = g.Next_Player()

	return nil
	}
}


func (g *Game)Call(p_name string)error{
	pindex = g.GetPlayerIndex(p_name)
	p = g.Players[pindex]
	if g.Current_bet > p.Money{
		return error.New("bet exceeds player's money")
	}
	p.Money -= g.Current_bet
	g.Pot += g.Current_bet
	g.Current_player = g.Next_Player()
	g.Timer := time.NewTimer(time.Second * 10)
	return nil
}


func (g *Game)Next_Player()string{
	pindex = g.GetPlayerIndex(g.Current_player)
	current = g.Players[pindex]
	for i := range g.Players{
		if i == len(g.Players) - 1 {
			g.Phase ++
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

//func (g *Game)Check(p_name string)error{
//	p = players[p_name]
//}

func (g *Game)Fold(player_name string)error{
	pindex = g.GetPlayerIndex(player_name)
	p = g.Players[pindex]
	if p.Folded == true {
		return error.New("Player has already folded")
	}else{
		p.Folded = true
		g.Current_player = g.Next_Player()
	}
	return nil
}


func (g *Game)Check(player_name, string)error{
	pindex = g.GetPlayerIndex(player_name)
	p = g.Players[pindex]
	if g.Current_bet > p.Bet{
		return error.New("Player's cannot check unless her current bet is equal to the current bet of the game")
	}
	return nil
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


func (g *Game)Discard(playerID stirng, cardIndexes []int) error{
	check = getHighestInt(cardIndexes)
	pindex = g.GetPlayerIndex(playerID)
	p = g.Players[pindex]
	if check > len(p.Hand){
		return error.New("Index is out of range for player's hand")
	}
	p.Card_Discard(card_indexes []int)
	return nil
}


func (p *Player)Card_Discard(card_indexes []int)error{
	var temp_hand []Card
	for i, n := range card_indexes{
		if n > len(p.Hand){
			return error.New("Index out of range of player's hand")
		}
	}
	for _, t := range card_indexes {
		//temp_hand[i] = hand[t]
		temp_hand = append(temp_hand, p.Hand[t])
	}
	var discarded_hand []Card
	//discarded_hand := make([]Card, 5)
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
	return nil
	}










