
/*(only class attributes needed for the game logic will be included here) */
package gamelogic

import(
	//"math/rand"
//	"time"
	"bufio"
  	"fmt"
  	"os"
  	"strings"
  	//"sort"
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
 	//reader := bufio.NewReader(os.Stdin)
 	cardTypes, suites := Init_card_cat()
 	pot := 0
 	deck := createDeck(cardTypes, suites)
 	deck = shuffle(deck)
 	fmt.Printf("Deck test: \n")
	for _, d := range deck{
		fmt.Printf("%s of %s ", d.Face, d.Suit)
	}
 	/*maybe show() some shuffle gif animation
 	/* each player pays the ante (may later swich to 'blind') */
 	for i := 0; i < len(players); i++{
 		players[i].Money -= ante
 		pot += ante
 		fmt.Printf("%s pays %d for ante \n", players[i].Name, ante)
	}
 
 	/*first round dealing */
 	fmt.Printf("The dealer suffles the cards and begins dealing... \n")
 	bufio.NewReader(os.Stdin).ReadBytes('\n')

 	d := 0
 	for d < 5{
 		for i := 0; i < len(players); i++{
 			card := draw(deck)
 			deck = deck[1:]
 			players[i].Hand = append(players[i].Hand, card)
 			fmt.Printf(" %s is delt a %s of %s \n ", players[i].Name, card.Face, card.Suit)
 		}
 		d++
 	}
 	/*first round betting */
 	pot = betting_round(players, minBet, maxBet, pot)
 	if len(players) == 1 {
 		winner := players[0]
 		return &winner
 	}
 	/* first draw */ 
 	deck = redraw(players, deck)
 	/* second round of betting */
 	pot = betting_round(players, minBet, maxBet, pot)
 	if len(players) == 1 {
 		winner := players[0]
 		return &winner
 	}
 	/* second draw */ 
 	deck= redraw(players, deck)
 	/* Third and final round of betting */
 	pot = betting_round(players, minBet, maxBet, pot)
 	/* sort hands by rank to prepare for hand comparisons */
 	for i := 0; i < len(players); i++{
 		if players[i].Folded == false{
 			players[i].sort_hand_by_rank()
 		}
 	}	
 	winner := showdown(players)
 	return winner
 		
 }



func place_bet_test(p Player, current int, min_bet int, max_bet int) int{
	reader := bufio.NewReader(os.Stdin)
	p.show_hand()
	fmt.Printf("%s, please place your bet \n", p.Name)
	fmt.Printf("%s has %d dollars and current bet is %d \n", p.Name, p.Money, current)
	fmt.Printf("Input -1 to fold, 0 to call or the amount you wish to raise \n")
	var bet int
    //_, err := fmt.Scanf("%d", &bet)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\r\n", "", -1)
	fmt.Printf("Input = %s", input)
	bet, err := strconv.Atoi(input)
	fmt.Println(bet)
	fmt.Println(err)
	return bet

}

 	

func betting_round(players []Player, minBet int, maxBet int, pot int)int{
	//reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(players); i++{
		players[i].Bet = 0
	}
	bet := 0
	for i := 0; i < len(players); i++{
 		if len(players) == 1{
 			return pot
 		}
 		bet = bet + players[i].place_bet(bet, minBet, maxBet)
 		pot = players[i].pay_bet(bet, pot)
 		players[i].Bet = bet
 		for j := 0; j < len(players); j++{
 			fmt.Printf("If raised, other players will need to match the bet of %d \n", bet)
 			if bet > players[j].Bet{
 				difference := bet - players[j].Bet
 				stay := players[j].stay_in(difference)
 				if stay == false{
 					fmt.Printf("%s is folding \n", players[j].Name)
 					players[j].Folded = true
 				}
 				if stay == true{
 					fmt.Printf("%s is staying ing the game \n", players[j].Name)
 					pot = players[j].pay_bet(difference, pot)
 				}
 			}
 		}
 		//bet += amount
 		fmt.Printf("bet is currently %d \n", bet)
 	}
 	fmt.Printf("Returning from betting round \n")
 	return pot
}


func stringToIntSlice(initial string) []int{
 	strs := strings.Split(initial, " ")
    ary := make([]int, len(strs))
    for i := range ary {
        ary[i], _ = strconv.Atoi(strs[i])
    }
    return ary
}

func redraw(players []Player, deck []Card)[]Card{
	reader := bufio.NewReader(os.Stdin)
	/*Each player may discard cards */
	for i := 0; i < len(players); i++{
 		if players[i].Folded == false{
 			/* remove = p.show_func() ask player which cards to remove return array of cards to be removed 
 			the array may be empty */
 			players[i].show_hand()
 			fmt.Printf("Would you like to discard any cards (y/n)? \n")
 			//var choice string
 			choice, _ := reader.ReadString('\n')
 			choice = strings.Replace(choice, "\r\n", "", -1)
 			if choice == "n"{
 				continue
 			}
 			fmt.Printf("Which cards would you like to discard? \n")
 			var input string
 			//_, err := fmt.Scanf("%s\n", &input)
 			input, err := reader.ReadString('\n')
 			fmt.Printf("err: %s", err)
 			input = strings.Replace(input, "\r\n", "", -1)
 			fmt.Printf("input: %s", err)
 			discard_index := stringToIntSlice(input)
 			fmt.Printf("Discard index %v: \n", discard_index)
 			players[i].Hand = discarded_hand(players[i].Hand, discard_index)	
 			}
 		}
 	/* Deal new cards to players */
 	fmt.Printf("The dealer will now deal new cards to the players... (press 'enter')\n")
 	bufio.NewReader(os.Stdin).ReadBytes('\n')

 	for i := 0; i < len(players); i++{
 		if players[i].Folded == false{
 			//hand_size := 5 - len(discard_index)
 			fmt.Printf("%s's hand len after discard: %d", players[i].Name, len(players[i].Hand))
 			replace := 5 - len(players[i].Hand) 
 			for j := 0; j < replace; j++{
 				card := draw(deck)
 				deck = deck[1:]
 				fmt.Printf("%s draws a %s of %s \n", players[i].Name, card.Face, card.Suit)
 				players[i].Hand = append(players[i].Hand, card)
 			}
 		}
 	}
 	return deck
}

func discarded_hand(hand []Card, discard_index []int)[]Card{
	//temp_hand := make([]Card, 5)
	var temp_hand []Card
	for _, t := range discard_index {
		//temp_hand[i] = hand[t]
		temp_hand = append(temp_hand, hand[t])
	}
	fmt.Printf("Print temp_hand\n")
	for _, t := range temp_hand{
		fmt.Println(t)
	}
	var discarded_hand []Card
	//discarded_hand := make([]Card, 5)
	index := 0
	for _, c := range hand{
		fmt.Printf("%s of %s \n", c.Face, c.Suit)
		 		check := true
		for _, d := range temp_hand{
			fmt.Printf("%s of %s \n", d.Face, d.Suit)
			if c.Face == d.Face && c.Suit == d.Suit{
				fmt.Printf("Do not include %s of %s \n", c.Face, c.Suit )
				check = false
			}
		}
		if check == true{
			fmt.Printf("Include %s of %s \n", c.Face, c.Suit)
			discarded_hand = append(discarded_hand, c)
			index ++
			}
		}
	fmt.Printf("Discarded Hand:\n")
	for _, c := range discarded_hand{
		fmt.Println(c)
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
	for i := 0; i < len(players); i++{
		if players[i].Folded == false{
			hand := players[i].Hand
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