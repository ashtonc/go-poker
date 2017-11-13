/* (only class attributes needed for the game logic will be included here) */
package logic

import (
	"math/rand"
	"time"
)

var cardTypes = [13]string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
var suites    = [4]string{"hearts", "spades", "clubs", "diamonds"}

func getIndex(array []string, item string) int {
	for i := 0; i < len(array); i++ {
		if array[i] == item {
			return i
		}
	}
	return -1
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

/*classes:   */

type Player struct { /* A more complete player struct will likely be someplace else in repo */
	name   string
	money  float32
	hand   []Card
	folded bool /*default value false */
	bet    int
}

type Globals struct {
	bets map[Player.name]int /*A map containing the bets of all players */
}

type Card struct {
	kind string
	suit string
	rank int
}

func newCard(kind string, suit string) Card {
	crd := new(Card)
	crd.kind = kind
	crd.suit = suit
	rank := getIndex(kind) + 1
	return crd
}

type Deck struct {
	cards [52]Card
}

func createDeck() Deck {
	/* create deck, adding each card looping through type and suite */
	deck := new(Deck)
	for _, t := range cardTypes {
		for _, s := range suites {
			crd = newCard(t, s)
			r := getIndex(t) + 1
			deck = append(deck, crd)
		}
	}
	return deck
}

func (d *Deck) shuffle() {
	/*....   randomly re-order the array */
	N = len(d)
	for i := 0; i < N; i++ {
		selection := rand.Intn(N - i)
		d[selection], d[i] = d[i], d[selection]
	}
}

func (d *Deck) draw() Card {
	crd := d[0]
	d = d[1:]
	return crd

}

type Dealer struct {
}

/* In casino play the first betting round begins with the player to the left of the big blind,
and subsequent rounds begin with the player to the dealer's left. Home games typically use an ante;
the first betting round begins with the player to the dealer's left, and the second round begins with
the player who opened the first round.

Play begins with each player being dealt five cards, one at a time, all face down. The remaining deck
 is placed aside, often protected by placing a chip or other marker on it. Players pick up the cards and
  hold them in their hands, being careful to keep them concealed from the other players, then a round
  of betting occurs.

If more than one player remains after the first round, the "draw" phase begins. Each player specifies
 how many of their cards they wish to replace and discards them. The deck is retrieved, and each player
 is dealt in turn from the deck the same number of cards they discarded so that each player again has
 five cards.

A second "after the draw" betting round occurs beginning with the player to the dealer's left or else
beginning with the player who opened the first round (the latter is common when antes are used instead
 of blinds). This is followed by a showdown, if more than one player remains, in which the player with
 the best hand wins the pot. */

func game(players []Player, ante int, minBet int, maxBet int, dealerToken int) {
	/* additional arguments might be added to function to denote display specifiations
	/* show() main game page
	/* create and shuffle deck of cards */
	pot = 0
	deck := createDeck()
	deck.shuffle()
	/*maybe show() some shuffle gif animation
	/* each player pays the ante (may later swich to 'blind') */
	for p := range len(players) {
		player.money -= ante
	}
	/*display some text noting that each player is paying the ante and update the money in view
	}
	/*first round dealing */
	for d := range 5 {
		for _, p := range len(players) {
			card = deck.draw()
			p.hand = append(p.hand, card)
			/*show card being added to player p - if p is "home" player show card face up at appriate
			position. Else, have card added to one of the "other" players.  */
		}
	}
	/*first round betting */
	betting_round(players)
	if len(players) == 1 {
		return players[0]
	}
	/* first draw */
	draw(players)
	/* second round of betting */
	betting_round(players)
	if len(players) == 1 {
		return players[0]
	}
	/* second draw */
	draw(players)
	/* Third and final round of betting */
	betting_round(players)
	/* sort hands by rank to prepare for hand comparisons */
	for p := range len(players) {
		if player.folded == false {
			sort_hand_by_rank(p.hand)
		}
	}
	showdown(players)

}

func sort_hand_by_rank(hand []Card) {
	sort.Slice(hand[:], func(i, j int) bool {
		return hand[i].rank < hand[j].rank
	})
}

func (p *Player) place_bet(current int, max_bet int, min_bet int) int {
	options = []string{"call", "fold", "raise"}
	if current == 0 {
		options = append(options, "check")
	}
	if player.money < current {
		p.folded = true
		return current
	}
	/*
		value = function(options, max_bet, min_bet)
		function should show (or call something to show) the appropriate player a pop-up or something with
		the options listed and ok button if call, return 0, if raise, return the amount added to bet,
		if fold, return -1. Do not let player bet more than his current money or the maximum bet*/
	if value == -1 {
		p.folded = true
		return current
	}
	amount = current + value
	return amount
}

func (p *Player) pay_bet(amount int) {
	player.money -= amount
}

func (p *Player) stay_in(difference) bool {
	if differnce > p.money {
		p.folded = true
		return false
		/* show_fun() letting player know that he or she is out of the game */
	}
	/* stay = show_func(difference) player p gets a pop up asking if he or she wishes to keep up with
	the latest bet in order to remain in the game */
	if stay == false {
		p.folded = true
		return false
	} else {
		p.money -= difference
		return true
	}
}

func (p *Player) remove_card(card) {
	index = getIndex(p.hand, card)
	p.hand = append(p.hand[:index], p.hand[index+1:]...)
}

func betting_round(players []Player) {
	bet = 0
	for i, p := range len(players) {
		if len(players) == 1 {
			return
		}
		amount := p.place_bet(bet)
		for q := range i {
			if amount > q.bet {
				p.stay_in(amount - q.bet)
				p.pay_bet(amount - q.bet)
			}
		}
		bet = amount
	}
}

func draw(players []Player) {
	/*Each player may discard cards */
	for i, p := range len(players) {
		if p.folded == false {
			/* remove = p.show_func() ask player which cards to remove return array of cards to be removed
			the array may be empty */
			/* This statement doesn't work without a remove
						for _, r := remove {
			 				p.remove_card(r)
			 			}
			*/
		}
	}
	/* Deal new cards to players */
	for i, p := range len(players) {
		if p.folded == false {
			hand_size = len(p)
			for hand_size < 5 {
				card = deck.draw()
				p.hand = append(p.hand, card)
			}
		}
	}
}

/* hand ranking :
	straight flush
	four of a kind
	full house
	straight
	flush
	three of a kind
	two pairs
	nothing
when determining the winner: consider hand rank of each player. If no tie of rank select winner
if two or more players are tied for top rank, call a second function that compares the hands based on
rank of the individual cards */

func showdown(players []Player) {
	score_category_map := map[string]int{
		"straight_flush":  0,
		"four_of_a_kind":  0,
		"full_house":      0,
		"flush":           0,
		"straight":        0,
		"three_of_a_kind": 0,
		"two_pairs":       0,
		"pair":            0,
		"nothing":         0,
	}

	for i, p := range player {
		if p.folded == false {
			hand = p.hand

			// A switch statement would work nicely here
			if check_straight_flush() {
				score_category_map["straight_flush"]++
			} else if check_four_of_a_kind() {
				score_category_map["four_of_a_kind"]++
			} else if check_full_house() {
				score_category_map["full_house"]++
			} else if check_flush() {
				score_category_map["flush"]++
			} else if check_stright() {
				score_category_map["straight"]++
			} else if check_three_of_a_kind() {
				score_category_map["three_of_a_kind"]++
			} else if check_two_pairs() {
				score_category_map["two_pairs"]++
			} else if check_pair() {
				score_category_map["pair"]++
			}
			/* not yet complete - will modify return values so that it is easier to determine the winner
			when the best two hands belong to the same category */

		}
	}
}

func check_flush(hand []Card) {
	suit := hand[0].suit
	for c := range len(p.hand) {
		if c.suite != suit {
			return false
		}
	}
	return true
}

func check_stright(hand []Card) {
	for i := 1; i < len(hand); i++ {
		if hand[i] != hand[i-1]+1 {
			return false
		}
	}
	return true
}

func check_straight_flush(hand, []card) {
	flush := check_flush(hand)
	straight := check_stright(hand)
	if flush == true && straight == true {
		return true
	} else {
		return false
	}
}

func check_four_of_a_kind(hand []Card) {
	if hand[0].rank == haad[3].rank[3] || hand[1].rank == hand[4].rank {
		return true
	} else {
		return false
	}
}

func check_full_house(hand []Card) {
	if hand[0].rank == hand[1].rank && hand[2].rank == hand[4].rank {
		return true
	}
	if hand[0].rank == hand[2].rank && hand[3].rank == hand[4].rank {
		return true
	} else {
		return false
	}
}

func check_three_of_a_kind(hand []Card) {
	if hand[0].rank == hand[2].rank || hand[1].rank == hand[3].rank || hand[2].rank == hand[4].rank {
		return true
	} else {
		return false
	}
}

/*A lower ranked unmatched card + 2 cards of the same rank + 2 cards of the same rank
2 cards of the same rank + a middle ranked unmatched card + 2 cards of the same rank
2 cards of the same rank + 2 cards of the same rank + a higher ranked unmatched card
*/

func check_two_pairs(hand []Card) {
	if hand[1].rank == hand[2].rank && hand[3].rank == hand[4].rank {
		return true
	}
	if hand[0].rank == hand[1].rank && hand[3].rank == hand[4].rank {
		return true
	}
	if hand[0].rank == hand[1].rank && hand[2].rank == hand[3].rank {
		return true
	}
	return false
}

/* this function must be used after the others, as it does not make sure that the hand has
nothing greater than two of a kind */
func check_pair(hand []Card) {
	try := 0
	for try < len(hand)-1 {
		for i := try + 1; i < len(hand); i++ {
			if hand[try].rank == hand[i].rank {
				return true
			}
		}
		try++
	}
}
