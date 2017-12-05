/*(only class attributes needed for the game logic will be included here) */
package gamelogic

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
	if buyin < g.Stakes.Ante*50 {
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
	g.Seats[seatNumber].Occupier = player
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

func (g *Game) NewRound(ante int, minBet int, maxBet int, dealterToken int) error {
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
	 		for i := 0; i < len(g.Players); i++{
	 			card := draw(g.Deck)
	 			g.Deck = g.Deck[1:]
	 			g.Players[i].Hand = append(g.Players[i].Hand, card)
	 			fmt.Printf(" %s is delt a %s of %s \n ", g.Players[i].Name, card.Face, card.Suit)
	 		}
	 		d++
	 	}
	 	return nil
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
		//g.SetTimer(10)
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

func (g *Game) Fold(player_name string)error{
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

func (g *Game)post_bet_winner_check() (error, *Player){
	remaining := g.check_num_players_remaining()
	if remaining == 0{
		return errors.New("There are zero players remaining! (something went wrong)"), nil
	}
	if remaining == 1{
		err, winner := g.find_winner()
		if err == nil{
			return nil, winner
		}
	}
	return nil, nil 
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


func (g *Game) check_num_players_remaining() int {
	remaining := 0
	for _, p := range g.Players {
		if p.Folded == false {
			remaining++
		}
	}
	return remaining
}

func (g *Game)find_winner() (error, *Player) {
	//function assumes only one players remains in the game
	for i, p := range g.Players {
		if p.Folded == false {
			return nil, &g.Players[i]
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

//Called when the last betting phase (phase 4) is over
func (g *Game)Showdown()*Player{
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
func (p *Player)card_histogram() {
	for _, crd := range p.Hand {
		p.Card_Hist[crd.Rank]++
	}
	fmt.Printf("Card Hist of %s: \n", p.Name)
	//for _, i := range p.Card_Hist {
	//	fmt.Printf(" %d \n", i)
	//}
}


func (g *Game)rank_hands() map[string][]int {
	//fmt.Printf("About to begin ranking hands...\n")
	score_category_map := make(map[string][]int)
	var value []int
	for rank := 0; rank < 9; rank++ {
		score_category_map["straight_flush"] = value
		score_category_map["four_of_a_kind"] = value
		score_category_map["full_house"] = value
		score_category_map["flush"] = value
		score_category_map["straight"] = value
		score_category_map["three_of_a_kind"] = value
		score_category_map["two_pairs"] = value
		score_category_map["pair"] = value
		score_category_map["nothing"] = value
	}
	for i := 0; i < len(g.Players); i++ {
		if g.Players[i].Folded == false {
			hand := g.Players[i].Hand
			if check_straight_flush(hand) {
				//score_category_map["straight_flush"]++
				score_category_map["straight_flush"] = append(score_category_map["straight_flush"], i)
				fmt.Printf("%s has a straight flush\n", g.Players[i].Name)
			} else if check_four_of_a_kind(hand) {
				//score_category_map["four_of_a_kind"]++
				score_category_map["four_of_a_kind"] = append(score_category_map["four_of_a_kind"], i)
				fmt.Printf("%s has a four of a kind\n", g.Players[i].Name)
			} else if check_full_house(hand) {
				//score_category_map["full_house"]++
				score_category_map["full_house"] = append(score_category_map["full_house"], i)
				fmt.Printf("%s has a full house\n", g.Players[i].Name)
			} else if check_flush(hand) {
				fmt.Printf("%s has a flush\n", g.Players[i].Name)
				score_category_map["flush"] = append(score_category_map["flush"], i)
				//player.Hand_Rank = 4
			} else if check_stright(hand) {
				//score_category_map["straight"]++
				score_category_map["straight"] = append(score_category_map["straight"], i)
				fmt.Printf("%s has a straight\n ", g.Players[i].Name)
			} else if check_three_of_a_kind(hand) {
				//score_category_map["three_of_a_kind"]++
				score_category_map["three_of_a_kind"] = append(score_category_map["three_of_a_kind"], i)
				fmt.Printf("%s has a three of a kind \n", g.Players[i].Name)
			} else if check_two_pairs(hand) {
				//score_category_map["two_pairs"] ++
				score_category_map["two_pairs"] = append(score_category_map["two_pairs"], i)
				fmt.Printf("%s has a two pairs\n ", g.Players[i].Name)
			} else if check_pair(hand) {
				//score_category_map["pair"]++
				score_category_map["pair"] = append(score_category_map["pair"], i)
				fmt.Printf("%s has one pair \n", g.Players[i].Name)
			} else {
				//score_category_map["nothing"]++
				score_category_map["nothing"] = append(score_category_map["nothing"], i)
				fmt.Printf("%s has nothing \n", g.Players[i].Name)
			}
			/* not yet complete - will modify return values so that it is easier to determine the winner
			when the best two hands belong to the same category */
		}
	}
	return score_category_map
}

func check_flush(hand []Card) bool {
	suit := hand[0].Suit
	for _, c := range hand {
		if c.Suit != suit {
			return false
		}
	}
	return true
}

func check_stright(hand []Card) bool {
	for i := 1; i < len(hand); i++ {
		if hand[i].Rank != hand[i-1].Rank+1 {
			return false
		}
	}
	return true
}

func check_straight_flush(hand []Card) bool {
	flush := check_flush(hand)
	straight := check_stright(hand)
	if flush == true && straight == true {
		return true
	} else {
		return false
	}
}

func check_four_of_a_kind(hand []Card) bool {
	if hand[0].Rank == hand[3].Rank || hand[1].Rank == hand[4].Rank {
		return true
	} else {
		return false
	}
}

func check_full_house(hand []Card) bool {
	if hand[0].Rank == hand[1].Rank && hand[2].Rank == hand[4].Rank {
		return true
	}
	if hand[0].Rank == hand[2].Rank && hand[3].Rank == hand[4].Rank {
		return true
	} else {
		return false
	}
}

func check_three_of_a_kind(hand []Card) bool {
	if hand[0].Rank == hand[2].Rank || hand[1].Rank == hand[3].Rank || hand[2].Rank == hand[4].Rank {
		return true
	} else {
		return false
	}
}

/*A lower ranked unmatched card + 2 cards of the same rank + 2 cards of the same rank
2 cards of the same rank + a middle ranked unmatched card + 2 cards of the same rank
2 cards of the same rank + 2 cards of the same rank + a higher ranked unmatched card
*/

func check_two_pairs(hand []Card) bool {
	if hand[1].Rank == hand[2].Rank && hand[3].Rank == hand[4].Rank {
		return true
	}
	if hand[0].Rank == hand[1].Rank && hand[3].Rank == hand[4].Rank {
		return true
	}
	if hand[0].Rank == hand[1].Rank && hand[2].Rank == hand[3].Rank {
		return true
	}
	return false
}

/* this function must be used after the others, as it does not make sure that the hand has
nothing greater than two of a kind */
func check_pair(hand []Card) bool {
	try := 0
	for try < len(hand)-1 {
		for i := try + 1; i < len(hand); i++ {
			if hand[try].Rank == hand[i].Rank {
				return true
			}
		}
		try++
	}
	return false
}

/*func get_player_index_from_name(players []Player, name string)int{
	for i, p := range players{
		if p.Name == name{
		return i
		}
	}
	return 0
} */

func (g *Game)DetermineWinner(score_board map[string][]int) *Player {
	//fmt.Printf("About to begin showdown \n")
	fmt.Printf("score_board: \n")
	//for key, value := range score_board {
	//	fmt.Printf("%s: %v \n", key, value)
	//}
	best := 0
	var winner Player
	if len(score_board["straight_flush"]) == 1 {
		win := score_board["straight_flush"][0]
		winner = g.Players[win]
		//return winner
	}else if len(score_board["straight_flush"]) > 1 {
		contenders := score_board["straight_flush"]
		for c := 0; c < len(contenders); c++ {
			if g.Players[c].Hand[4].Rank > best {
				best = g.Players[c].Hand[4].Rank
				winner = g.Players[contenders[c]]
			}
		}
	}else if len(score_board["four_of_a_kind"]) == 1 {
		win := score_board["four_of_a_kind"][0]
		winner = g.Players[win]
	} else if len(score_board["four_of_a_kind"]) > 1 {
		contenders := score_board["four_of_a_kind"]
		for c := 0; c < len(contenders); c++ {
			if g.Players[contenders[c]].find_four_of_kind_rank() > best {
				best = g.Players[contenders[c]].find_four_of_kind_rank()
				winner = g.Players[contenders[c]]
			}
		}
	} else if len(score_board["full_house"]) == 1 {
		win := score_board["full_house"][0]
		winner = g.Players[win]
	} else if len(score_board["full_house"]) > 1 {
		contenders := score_board["full_house"]
		for c := 0; c < len(contenders); c++ {
			if g.Players[contenders[c]].find_three_of_kind_rank() > best {
				best = g.Players[contenders[c]].find_three_of_kind_rank()
				winner = g.Players[contenders[c]]
			}
		}
	} else if len(score_board["flush"]) == 1 {
		win := score_board["full_house"][0]
		winner = g.Players[win]
	} else if len(score_board["flush"]) > 1 {
		contenders := score_board["flush"]
		win := find_best_nothing(contenders, g.Players)
		winner = g.Players[win]

	} else if len(score_board["straight"]) == 1 {
		win := score_board["straight"][0]
		winner = g.Players[win]
	} else if len(score_board["straight"]) > 1 {
		contenders := score_board["straight"]
		win := find_best_nothing(contenders, g.Players)
		winner = g.Players[win]
	} else if len(score_board["three_of_a_kind"]) == 1 {
		win := score_board["three_of_a_kind"][0]
		winner = g.Players[win]
	} else if len(score_board["three_of_a_kind"]) > 1 {
		contenders := score_board["three_of_a_kind"]
		for c := 0; c < len(contenders); c++ {
			if g.Players[contenders[c]].find_three_of_kind_rank() > best {
				best = g.Players[contenders[c]].find_three_of_kind_rank()
				winner = g.Players[contenders[c]]
			}
		}
	} else if len(score_board["two_pairs"]) == 1 {
		win := score_board["two_pairs"][0]
		winner = g.Players[win]
	} else if len(score_board["two_pairs"]) > 1 {
		contenders := score_board["two_pairs"]
		best2 := 0
		for i := 0; i < len(contenders); i++ {
			rank := g.Players[contenders[i]].best_pair()
			fmt.Printf("%s best pair rank is %d \n", g.Players[contenders[i]].Name, rank)
			if rank > best {
				best = rank
				winner = g.Players[contenders[i]]
			} else if rank == best && best > 0 {
				fmt.Printf("tie detected \n")
				best = rank
				for j := 0; j < len(contenders); j++ {
					if g.Players[contenders[j]].best_pair() < best {
						continue
					}
					rank2 := g.Players[contenders[j]].second_best_pair()
					fmt.Printf("%s second best pair rank is %d \n", g.Players[contenders[j]].Name, rank2)
					if rank2 > best2 {
						best2 = rank2
						winner = g.Players[contenders[j]]
					} else if rank2 == best2 && best2 > 0 {
						fmt.Printf("Second tie!! \n")
						win := find_best_nothing(contenders, g.Players)
						winner = g.Players[win]
						fmt.Printf("current winner: %s \n", winner.Name)
						return &winner
					}
				}
			}
		}
	} else if len(score_board["pair"]) == 1 {
		win := score_board["pair"][0]
		winner = g.Players[win]
	} else if len(score_board["pair"]) > 1 {
		contenders := score_board["two_pairs"]
		fmt.Printf("Best is %d \n", best)
		for i := 0; i < len(contenders); i++ {
			fmt.Printf("i is %d \n", i)
			rank := g.Players[contenders[i]].best_pair()
			fmt.Printf("Rank is %d \n", rank)
			if rank > best {
				best = rank
				winner = g.Players[contenders[i]]
			} else if rank == best && best > 0 {
				index := find_best_nothing(contenders, g.Players)
				winner = g.Players[index]
				return &winner
			}
		}
	} else {
		contenders := score_board["nothing"]
		index := find_best_nothing(contenders, g.Players)
		winner = g.Players[index]
	}
	return &winner
}

func find_best_nothing(indexes []int, players []Player) int {

	for i := 4; i > 0; i-- {
		win_indx := 0
		highest := 0
		tie := false
		for _, j := range indexes {
			fmt.Printf("%s has a rank %d card \n", players[j].Name, players[j].Hand[i].Rank)
			if players[j].Hand[j].Rank > highest {
				if players[j].Hand[j].Rank == highest && highest > 0 {
					tie = true
				}
				win_indx = j
			}
		}
		if tie == false {
			return win_indx
		}
	}
	return -1
}