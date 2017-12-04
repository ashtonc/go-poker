/*(only class attributes needed for the game logic will be included here) */
package gamelogic

<<<<<<< HEAD
import(
	_"time"
	_"bufio"
  	_"os"  
  	"errors"
 // 	"time"
  	"fmt"
  	"math/rand"
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

func (g *Game) Join(name string, buyin int, seatNumber int) error {
	if g.Seats[seatNumber].Occupied == true {
		return errors.New("Seat is already Occupied")
	}
	if buyin > g.Stakes.Ante*100 {
		return errors.New("Buyin exceeds the limit for this table")
	}
<<<<<<< HEAD
	if buyin < g.Stakes.Ante*50{
=======
	if buyin < g.Stakes.Ante*50 {
>>>>>>> ecf4b2a38403a8c5b2c26814ca83d6a37f402e76
		return errors.New("Buyin is too low for this table")
	}
	player := new(Player)
	player.Name = name
	player.Money = buyin
	player.Seat = seatNumber
	player.Discarded = false
	player.Folded = false
	player.Called = false
	g.Sitters = append(g.Players, *player)
	g.Seats[seatNumber].Occupied = true
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

func (g *Game) NewRound(players []Player, ante int, minBet int, maxBet int, dealterToken int) error {
	for _, p := range g.Sitters {
		g.Players = append(g.Players, p)
	}
	if len(g.Players) < 2 {
		return errors.New("A round of Poker requires at least two players")
	}
	g.Phase = 1
	g.Bet_Counter = -1
	g.Stakes.MaxBet = maxBet
	g.Stakes.MinBet = minBet
	g.Dealer_Token = dealterToken + 1
	g.Current_Player = g.Players[dealterToken].Name
	cardTypes, suites := Init_card_cat()
	g.Pot = 0
	// ant played by each player
	for i := 0; i < len(g.Players); i++ {
		g.Players[i].Money -= ante
		g.Pot += ante
	}
	//Create deck, shuffle cards, deal cards to players
	g.Deck = createDeck(cardTypes, suites)
	g.Deck = shuffle(g.Deck)
	d := 0
	 	for d < 5{
	 		for i := 0; i < len(players); i++{
	 			card := draw(g.Deck)
	 			g.Deck = g.Deck[1:]
	 			g.Players[i].Hand = append(g.Players[i].Hand, card)
	 			fmt.Printf(" %s is delt a %s of %s \n ", g.Players[i].Name, card.Face, card.Suit)
	 		}
	 		d++
	 	}
	}

// Eachetting round lasts until each player has either: (a) folded (b) called
func (g *Game) Bet(p_name string, bet int) error {
	pindex, error := g.GetPlayerIndex(p_name)
	if error == nil {
		if g.Players[pindex].Folded == true {
			return errors.New("Player has already folded and so cannot bet")
		}
		if bet > g.Players[pindex].Money {
			return errors.New("bet exceeds player's money")
		}
		if bet > g.Stakes.MaxBet {
			return errors.New("Bet exceeds maximum bet")
		}
		if bet < g.Stakes.MinBet {
			return errors.New("Bet is below the minimum bet")
		}
		if bet <= g.Current_Bet {
			return errors.New("New bet must be geater than the current bet")
		}
		g.Players[pindex].Money -= bet
		g.Pot += bet
		g.Current_Bet = bet
		g.ResetBetCounter()
		g.Current_Player = g.Next_Player()
		//g.check_if_betting_ends()
	}

	return nil
}

func (g *Game) ResetBetCounter() {
	players_in := 0
	for i := range g.Players {
		if g.Players[i].Folded == false {
			players_in += 1
		}
		g.Bet_Counter = players_in
	}
}

func (g *Game) Call(p_name string) error {
	pindex, error := g.GetPlayerIndex(p_name)
	if error == nil {

		if g.Players[pindex].Folded == true {
			return errors.New("Player has already folded and cannot call.")
		}
		if g.Current_Bet > g.Players[pindex].Money {
			return errors.New("bet exceeds player's money")
		}
		g.Players[pindex].Money -= g.Current_Bet
		g.Players[pindex].Called = true
		g.Pot += g.Current_Bet
		g.Current_Player = g.Next_Player()
		g.Bet_Counter -= 1
		if g.Bet_Counter == 0 {
			g.Next_Phase()
		}

		return nil
	} else {
		return error
	}
}

func (g *Game) Next_Player() string {
	pindex, error := g.GetPlayerIndex(g.Current_Player)
	if error == nil {
		current := g.Players[pindex]
		for i := range g.Players {
			if i == len(g.Players)-1 {
				return g.Players[0].Name
			} else {
				if g.Players[i].Name == current.Name {
					return g.Players[i+1].Name
				}
			}
		}
		g.SetTimer(10)
	}

	return "-1"
}

func (g *Game) Next_Phase() error {
	if g.Phase > 4 {
		return errors.New("The game phase cannot be further incremented \n")
	}
	g.Phase += 1
	g.Current_Player = g.Players[g.Dealer_Token].Name
	return nil
}

func (g *Game) Fold(player_name string) error {
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
			g.Next_Phase()
		}
	}
	return nil
}

func (g *Game) Check(player_name string) error {
	pindex, err := g.GetPlayerIndex(player_name)
	if err == nil {
		if g.Current_Bet > g.Players[pindex].Bet {
			return errors.New("Player's cannot check unless her current bet is equal to the current bet of the game")
		}
		if g.Bet_Counter > 0 {
			g.Bet_Counter -= 1
		}
		if g.Bet_Counter == 0 {
			g.Next_Phase()
		}
	}
	return nil
}

//Call when (phase == 2 or phase == 4) and g.Current_Player.seat == g.Dealer_Token)
func (g *Game) check_if_winner() string {
	remaining := check_num_players_remaining(g.Players)
	if remaining == 1 {
		winner := find_winner(g.Players)
		return winner
	}
	return "-1"
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

func find_winner(players []Player) string {
	//function assumes only one players remains in the game
	for _, p := range players {
		if p.Folded == false {
			return p.Name
		}
	}
	p := players[0] //just to make go happy
	return p.Name
}

func (g *Game) check_for_winner() string {
	remaining := check_num_players_remaining(g.Players)
	if remaining < 2 {
		winner := find_winner(g.Players)
		return winner
	}
	return "-1"
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

func (g *Game) Discard(playerID string, cardIndexes []int) error {
	check := getHighestInt(cardIndexes)
	pindex, err := g.GetPlayerIndex(playerID)
	if err == nil {
		if g.Players[pindex].Discarded == true {
			for i := range g.Players {
				g.Players[i].Discarded = false
				g.Next_Phase()
				return nil
			}
			g.Phase += 1
			g.Current_Player = g.Players[g.Dealer_Token].Name
			return nil
		}
		if check > len(g.Players[pindex].Hand) {
			return errors.New("Index is out of range for player's hand")
		}
		discard := g.Card_Discard(g.Players[pindex].Name, cardIndexes)
		if discard == nil {
			g.Redraw(g.Players[pindex].Name)
			g.Players[pindex].Discarded = true
			g.Next_Player()
			return nil
		}
	}
	return nil
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
