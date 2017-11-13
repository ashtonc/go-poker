
/*(only class attributes needed for the game logic will be included here) */
package gamelogic

import(
	//"math/rand"
//	"time"
	"bufio"
  	"fmt"
  	"os"
  	"strings"
  	"sort"
  	"strconv"
)




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

 func Game(players []Player, ante int, minBet int, maxBet int, dealerToken int)*Player{
 	/* additional arguments might be added to function to denote display specifiations 
 	/* show() main game page
 	/* create and shuffle deck of cards */
 	cardTypes, suites := Init_card_cat()
 	pot := 0
 	deck := createDeck(cardTypes, suites)
 	shuffle(deck)
 	/*maybe show() some shuffle gif animation
 	/* each player pays the ante (may later swich to 'blind') */
 	for _, p := range players{
 		p.Money -= ante
 	}
 	/*display some text noting that each player is paying the ante and update the money in view
 	}
 	/*first round dealing */
 	d := 0
 	for d < 5{
 		for _, p := range players{
 			card := draw(deck)
 			p.Hand = append(p.Hand, card)
 				/*show card being added to player p - if p is "home" player show card face up at appriate 
 				position. Else, have card added to one of the "other" players.  */
 		}
 		d++
 	}
 	/*first round betting */
 	betting_round(players, minBet, maxBet, pot)
 	if len(players) == 1 {
 		winner := players[0]
 		return &winner
 	}
 	/* first draw */ 
 	redraw(players, deck)
 	/* second round of betting */
 	betting_round(players, minBet, maxBet, pot)
 	if len(players) == 1 {
 		winner := players[0]
 		return &winner
 	}
 	/* second draw */ 
 	redraw(players, deck)
 	/* Third and final round of betting */
 	betting_round(players, minBet, maxBet, pot)
 	/* sort hands by rank to prepare for hand comparisons */
 	for _, p := range players{
 		if p.Folded == false{
 			sort_hand_by_rank(p.Hand)
 		}
 	} 	
 	winner := showdown(players)
 	return winner
 		
 }

func sort_hand_by_rank(hand []Card){
	sort.Slice(hand[:], func(i, j int) bool {
    return hand[i].Rank < hand[j].Rank
	})
}



func place_bet_test(p Player, current int, min_bet int, max_bet int) int{
	reader := bufio.NewReader(os.Stdin)
	show_hand(p.Hand)
	fmt.Printf("Please place your bet")
	fmt.Printf("You have %d dollats and current bet is %d", p.Money, current)
	fmt.Printf("Input -1 to fold, 0 to call or the amount you wish to raise")
	input, _ := reader.ReadString('\n')
	bet, _ := strconv.Atoi(input)
	return bet


}

 func (p *Player) place_bet(current int, max_bet int, min_bet int) int {
 	//options := []string {"call", "fold", "raise"}
 	//if current == 0{
 	//	options := append(options, "check")
 	//}
 	if p.Money < current{
 		p.Folded = true
 		return current
 		fmt.Printf("You need to have %d dollars to stay in the game and only have %d", current, p.Money)
 		fmt.Printf("You have no choice but to fold")
 	}
 	fmt.Printf("Place bet for player %s", p.Name)
 	show_hand(p.Hand)
 	value := place_bet_test(*p, current, min_bet, max_bet)
 	/* 
 	value = function(options, max_bet, min_bet)
 	function should show (or call something to show) the appropriate player a pop-up or something with 
 	the options listed and ok button if call, return 0, if raise, return the amount added to bet,
 	if fold, return -1. Do not let player bet more than his current money or the maximum bet*/
 	if value == -1{
 		p.Folded = true
 		return current
 	}
    amount := current + value 
    return amount
}

 

 func (p *Player) pay_bet(amount int, pot int){
 	p.Money -= amount 
 	pot += amount
 }

 func (p *Player) stay_in(difference int) bool {
 	reader := bufio.NewReader(os.Stdin)
 	if difference > p.Money{
 		p.Folded = true
 		return false
 		fmt.Print("%s unable to meet the raised bet and is out of the game", p.Name)
 	}
	fmt.Printf("The bet has been rased by %d", difference)
	fmt.Printf("Will you stay in the game? (Y/N")
	input, _ := reader.ReadString('\n')
	var stay bool
	if input == "N"{
		stay = false
	}else{stay = true
	}

 	/* show_fun() letting player know that he or she is out of the game */
 	
 	/* stay = show_func(difference) player p gets a pop up asking if he or she wishes to keep up with
 	the latest bet in order to remain in the game */
 	if stay == false{
 		p.Folded = true
 		return false
 	}else{
 		return true
 	}
 }
	


func betting_round(players []Player, minBet int, maxBet int, pot int){
	for _, p := range players{
		p.Bet = 0
	}
	bet := 0
 	for _, p := range players{
 		if len(players) == 1{
 			return
 		}
 		amount := p.place_bet(bet, minBet, maxBet)

 		for _, q := range players{
 			if amount > q.Bet{
 				difference := amount - q.Bet
 				q.stay_in(difference)
 				q.pay_bet(difference, pot)
 			}
 		}
 		bet = amount
 	}
}

func show_hand(hand []Card){
	for i, crd := range hand{
		fmt.Printf("%d %s of %s", i, crd.Face, crd.Suit)
	}
}

func stringToIntSlice(initial string) []int{
 	strs := strings.Split(initial, " ")
    ary := make([]int, len(strs))
    for i := range ary {
        ary[i], _ = strconv.Atoi(strs[i])
    }
    return ary
}

func redraw(players []Player, deck []Card){
	reader := bufio.NewReader(os.Stdin)
	/*Each player may discard cards */
 	for _, p := range players{
 		if p.Folded == false{
 			/* remove = p.show_func() ask player which cards to remove return array of cards to be removed 
 			the array may be empty */
 			hand := p.Hand
 			show_hand(hand)
 			fmt.Printf("Which cards would you like to discard?")
 			input, _ := reader.ReadString('\n')
 			discard_index := stringToIntSlice(input)
 			p.Hand = discarded_hand(p.Hand, discard_index)	
 			}
 		}
 	/* Deal new cards to players */
 	for _, p := range players{
 		if p.Folded == false{
 			hand_size := len(p.Hand)
 			for hand_size < 5{
 				card := draw(deck)
 				p.Hand = append(p.Hand, card)
 				hand_size ++
 			}
 		}
 	}
}

func discarded_hand(hand []Card, discard_index []int)[]Card{

	discard := make([]Card, 5)
	for _, d := range discard_index {
		discard = append(discard, hand[d])
	}

	discarded_hand := make([]Card, 5)
	check := true
	for _, c := range hand{
		for _, d := range hand{
			if c.Face == d.Face && c.Suit == d.Suit{
				check = false
			}
		if check == true{
			discarded_hand = append(discarded_hand, c)
			}
		}
	}
	return discarded_hand
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

func showdown(players []Player)*Player{
	score_category_map := map[string]int{
		"straight_flush": 0,
		"four_of_a_kind": 0,
		"full_house": 0,
		"flush": 0,
		"straight": 0,
		"three_of_a_kind": 0,
		"two_pairs": 0,
		"pair": 0,
		"nothing": 0,
	}

	for _, p := range players{
		if p.Folded == false{
			hand := p.Hand
			if check_straight_flush(hand){
				score_category_map["straight_flush"]++
			}else if check_four_of_a_kind(hand){
				score_category_map["four_of_a_kind"]++
			}else if check_full_house(hand){
				score_category_map["full_house"]++
			}else if check_flush(hand){
				score_category_map["flush"]++
			}else if check_stright(hand){
				score_category_map["straight"]++
			}else if check_three_of_a_kind(hand){
				score_category_map["three_of_a_kind"]++
			}else if check_two_pairs(hand){
				score_category_map["two_pairs"] ++
			}else if check_pair(hand){
				score_category_map["pair"]++
			}else{
				score_category_map["nothing"]++
			}
		/* not yet complete - will modify return values so that it is easier to determine the winner
			when the best two hands belong to the same category */
			
		}
		
	}
	p := players[0]
	return &p
}

func check_flush(hand []Card)bool{
	suit := hand[0].Suit
	for _, c := range hand{
		if c.Suit != suit{
			return false
		}
	}
	return true
}

func check_stright(hand []Card)bool{
	for i := 1; i < len(hand); i++{
		if hand[i].Rank != hand[i-1].Rank+1{
			return false
		}
	}
	return true
}

func check_straight_flush(hand []Card)bool{
	flush := check_flush(hand)
	straight := check_stright(hand)
	if flush == true && straight == true{
		return true
	}else{
		return false
	}
}

func check_four_of_a_kind(hand []Card)bool{
	if hand[0].Rank == hand[3].Rank || hand[1].Rank == hand[4].Rank{
		return true
	}else{
  		return false
  	}
}


func check_full_house(hand []Card)bool{
	if hand[0].Rank == hand[1].Rank && hand[2].Rank == hand[4].Rank{
		return true
	}
	if hand[0].Rank == hand[2].Rank && hand[3].Rank == hand[4].Rank{
		return true
	}else{
		return false
	}
}

func check_three_of_a_kind(hand []Card)bool{
	if hand[0].Rank == hand[2].Rank || hand[1].Rank == hand[3].Rank || hand[2].Rank == hand[4].Rank{
		return true
	}else{
		return false
	}
}


/*A lower ranked unmatched card + 2 cards of the same rank + 2 cards of the same rank
2 cards of the same rank + a middle ranked unmatched card + 2 cards of the same rank
2 cards of the same rank + 2 cards of the same rank + a higher ranked unmatched card
*/


func check_two_pairs(hand []Card)bool{
	if hand[1].Rank == hand[2].Rank && hand[3].Rank == hand[4].Rank{
		return true
	}
	if hand[0].Rank == hand[1].Rank && hand[3].Rank == hand[4].Rank{
		return true
	}
	if hand[0].Rank == hand[1].Rank && hand[2].Rank == hand[3].Rank{
		return true
	}
	return false
}

/* this function must be used after the others, as it does not make sure that the hand has
nothing greater than two of a kind */
func check_pair(hand []Card)bool{
	try := 0
	for try < len(hand)-1{
		for i := try+1; i < len(hand); i++{
			if hand[try].Rank == hand[i].Rank{
				return true
			}
		}
		try ++
	}
	return false
}