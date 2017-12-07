package gamelogic

import (
	_ "bufio"
	"errors"
	"fmt"
	"log"
	"math/rand"
	_ "os"
	_ "time"
)

/*func (g *Game)SetTimer(amount int){
	timer := time.NewTimer(time.Second * 10)
 	g.Timer = *timer
} */

func (g *Game) GetPlayerIndex(name string) (int, error) {
	for i, p := range g.Players {
		if p.Name == name {
			return i, nil
		}
	}
	return -1, errors.New("There is no player of that name")
}

func (g *Game) GetSitterIndex(name string) (int, error) {
	for i, p := range g.Sitters {
		if p.Name == name {
			return i, nil
		}
	}
	return 0, errors.New("There is no sitter of that name")
}

func (g *Game) Get_current_player_name() string {
	current, err := g.GetPlayerIndex(g.Current_Player)
	if err != nil {
		log.Fatal(err)
	}
	player := g.Players[current].Name
	return player
}

func createDeck(cardTypes []string, suites []string) []Card {
	/* create deck, adding each card looping through type and suite */
	deck := make([]Card, 52)
	count := 0
	for _, t := range cardTypes {
		for _, s := range suites {
			crd := newCard(t, s, cardTypes)
			deck[count] = *crd
			count++
		}
	}
	return deck
}

func shuffle(d []Card) []Card {
	/*....   randomly re-order the array */
	for i := len(d) - 1; i > 0; i-- {
		selection := rand.Intn(i + 1)
		d[i], d[selection] = d[selection], d[i]
	}
	return d
}

func draw(d []Card) Card {
	crd := d[0]
	return crd
}

func (g *Game) ResetBetCounter() {
	players_in := 0
	for i := range g.Players {
		//fmt.Printf("%s \n", g.Players[i].Name)
		if g.Players[i].Folded == false {
			players_in += 1
		}
	}
	g.Bet_Counter = players_in
	fmt.Printf("The bet counter was set to %d \n", players_in)
}

func (g *Game) Check_if_end_of_bet_correct() error {
	for i := range g.Players {
		if g.Players[i].Bet != g.Current_Bet && g.Players[i].Folded == false {
			return errors.New("Not all players made same bet at end of betting round!")
		}
	}
	return nil
}

func (g *Game) reset_bets() {
	for i := range g.Players {
		g.Players[i].Bet = 0
	}
	g.Current_Bet = 0
}

func (g *Game) Next_Player() string {
	pindex, error := g.GetPlayerIndex(g.Current_Player)
	if error == nil {
		current := pindex
		for {
			if current == len(g.Players)-1 {
				current = 0
				fmt.Printf("At end of table.. \n")
				if g.Players[current].Folded == false {
					fmt.Printf("Current player is now %s \n", g.Players[current].Name)
					return g.Players[0].Name
				}
			} else {
				//if g.Players[i].Name == current.Name {
				current += 1
				if g.Players[current].Folded == false {
					fmt.Printf("Current player is now %s \n", g.Players[current].Name)
					return g.Players[current].Name
				}
			}
		}
	}
	//g.SetTimer(10)
	return "-1"
}

func (g *Game) Next_Phase() error {
	if g.Phase > 4 {
		return errors.New("The game phase cannot be further incremented \n")
	}
	g.Phase += 1
	for i := range g.Players {
		g.Players[i].Discarded = false
		g.Players[i].Bet = 0
	}
	//fmt.Printf("Dealer token is: %d \n", g.Dealer_Token)
	g.Current_Player = g.Players[g.Dealer_Token].Name
	//current := g.Dealer_Token
	if g.Players[g.Dealer_Token].Folded == true {
		g.Current_Player = g.Next_Player()
	}
	g.ResetBetCounter()
	return nil
}

func (g *Game) check_if_discard_phase_complete() bool {
	for i := range g.Players {
		if g.Players[i].Folded == false && g.Players[i].Discarded == false {
			return false
		}
	}
	return true
}

func (g *Game) Winner_check() (error, *Player) {
	remaining := g.check_num_players_remaining()
	if remaining == 0 {
		return errors.New("There are zero players remaining! (something went wrong)"), nil
	}
	if remaining == 1 {
		err, winner := g.find_winner()
		if err == nil {
			return nil, winner
		}
	}
	return nil, nil
}

func (g *Game) check_num_players_remaining() int {
	remaining := 0
	for _, p := range g.Players {
		if p.Folded == false {
			remaining++
		}
	}
	return remaining
}

func (g *Game) find_winner() (error, *Player) {
	//function assumes only one players remains in the game
	for i, p := range g.Players {
		if p.Folded == false {
			return nil, g.Players[i]
		}
	}
	//just to make Go happy
	return errors.New("No winner was found!"), nil
}

func getHighestInt(array []int) int {
	highest := 0
	for _, v := range array {
		if v > highest {
			highest = v
		}
	}
	return highest
}

func (g *Game) Card_Discard(player string, card_indexes []int) error {
	pindex, err := g.GetPlayerIndex(player)
	if err == nil {
		var temp_hand []Card
		for _, n := range card_indexes {
			if n > len(g.Players[pindex].Hand) {
				return errors.New("Index out of range of player's hand")
			}
		}
		for _, t := range card_indexes {
			//temp_hand[i] = hand[t]
			temp_hand = append(temp_hand, g.Players[pindex].Hand[t])
		}
		var discarded_hand []Card
		index := 0
		for _, c := range g.Players[pindex].Hand {
			check := true
			for _, d := range temp_hand {
				if c.Face == d.Face && c.Suit == d.Suit {
					check = false
				}
			}
			if check == true {
				discarded_hand = append(discarded_hand, c)
				index++
			}
		}
		g.Players[pindex].Hand = discarded_hand
	}
	return nil
}

func (g *Game) Redraw(player string) {
	pindex, err := g.GetPlayerIndex(player)
	if err == nil {
		replace := 5 - len(g.Players[pindex].Hand)
		for j := 0; j < replace; j++ {
			card := draw(g.Deck)
			g.Deck = g.Deck[1:]
			fmt.Printf("%s draws a %s of %s \n", g.Players[pindex].Name, card.Face, card.Suit)
			g.Players[pindex].Hand = append(g.Players[pindex].Hand, card)
		}
	}
}

//Called when the last betting phase (phase 4) is over
