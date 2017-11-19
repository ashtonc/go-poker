
/*(only class attributes needed for the game logic will be included here) */
package gamelogic

import(
	"math/rand"
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
 	remaining := check_num_players_remaining(players)
 	if remaining < 2 {
 		winner := find_winner(players)
 		return winner
 	}
 	/* first draw */ 
 	deck = redraw(players, deck)
 	/* second round of betting */
 	pot = betting_round(players, minBet, maxBet, pot)
 	remaining = check_num_players_remaining(players)
 	if remaining < 2 {
 		winner := find_winner(players)
 		return winner
 	}
 	/* second draw */ 
 	deck= redraw(players, deck)
 	/* Third and final round of betting */
 	pot = betting_round(players, minBet, maxBet, pot)
 	remaining = check_num_players_remaining(players)
 	if remaining < 2 {
 		winner := find_winner(players)
 		return winner
 	}
 	/* sort hands by rank to prepare for hand comparisons */
 	for i := 0; i < len(players); i++{
 		if players[i].Folded == false{
 			players[i].sort_hand_by_rank()
 			fmt.Printf("Sorted hand: \n")
 			for _, crd := range players[i].Hand{
 				fmt.Printf("%s of %s \n", crd.Face, crd.Suit)
 			}
 		}
 	}
 	score_board :=rank_hands(players)	
 	winner := showdown(players, score_board)
 	fmt.Printf("%s win a pot worth %d \n", winner.Name, pot)
 	winner.Money += pot
 	return winner

 		
 }

 func createDeck(cardTypes []string, suites []string)[]Card{
	/* create deck, adding each card looping through type and suite */
	deck := make([]Card, 52)
	count := 0
	for _, t := range cardTypes{
		for _, s := range suites{
			crd := newCard(t, s, cardTypes)
			deck[count] = *crd
			count++

		}
	}
	//fmt.Printf("Deck test: \n")
	//for _, d := range deck{
	//	fmt.Printf("%s of %s ", d.Face, d.Suit)
	//}
	return deck
}


func shuffle(d []Card)[]Card{
	/*....   randomly re-order the array */
	for i := len(d) - 1; i > 0; i-- {
		selection := rand.Intn(i + 1)
		d[i], d[selection] = d[selection], d[i]
	}
	return d
}

func draw(d []Card)Card{
	crd := d[0]
	return crd
}


 func check_num_players_remaining(players []Player)int{
 	remaining := 0
 	for _, p := range players{
 		if p.Folded == false{
 			remaining ++
 		}
 	}
 	fmt.Printf("%d players remaining in this game\n", remaining)
 	return remaining
 }

func find_winner(players []Player)*Player{
	//function assumes only one players remains in the game
	for _, p := range players{
		if p.Folded == false{
			return &p
		}
	}
	p := players[0]		//just to make go happy
	return &p
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
 	strs := strings.Split(initial," ")
 	fmt.Printf("Len of int-string: %d \n", len(strs))
 	fmt.Printf("%s \n", strs)
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
 			fmt.Printf("err: %s \n", err)
 			input = strings.Replace(input, "\r\n", "", -1)
 			fmt.Printf("input: %s \n", input)
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

func (p *Player)card_histogram(){
	/*p.Card_Hist = map[string]int{
		"2":	0,
		"3":	0,
		"4":	0,
		"5":	0,
		"6":	0,
		"7":	0,
		"8":	0,
		"9":	0,
		"10":	0,
		"Jack":	0,
		"Queen":0,
		"King":	0,
		"Ace":	0, 
	} */
	for _, crd := range p.Hand{
		p.Card_Hist[crd.Rank]++
	}

}

func max_dict(dict map[string]int){
	max := 1
	for _, v := range dict{
		if v > max{
		max = v
		}
	}
}

/*func (p *Player)rank_hand(){
	first 	:= 1
	second  := 1
	for k, v := range p.Card_Hist{
		if v > 1{
			if first !=1{
				second = first

			}
		}
	}
	first_tier := 1
	second_tier := 1
	larg_group := 0
	small_group := 0
	for r := 13; r > 0; r--{
		//if r > first_tier
		if first_tier != 1{
			second_tier := first_tier
			two_rank := three_rank  
		}
		first_tier = x
		large_group = x 
	}
} */


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

func rank_hands(players []Player)map[string][]int{
	fmt.Printf("About to begin ranking hands...\n")
	score_category_map := make(map[string][]int)

	//	"straight_flush": []Player,
	//	"four_of_a_kind": []Player,
	//	"full_house": []Player,
	//	"flush": []Player,
	///	"straight": []Player,
	//"three_of_a_kind": []Player,
	//	"two_pairs": []Player,
	//	"pair": []Player,
	//	"nothing": []Player,
	//}
	for i := 0; i < len(players); i++{
		if players[i].Folded == false{
			hand := players[i].Hand
			if check_straight_flush(hand){
				//score_category_map["straight_flush"]++
				score_category_map["straight_flush"] = append(score_category_map["straight_flush"], i )
				//player.Hand_Rank = 1
			}else if check_four_of_a_kind(hand){
				//score_category_map["four_of_a_kind"]++
				score_category_map["four_of_a_kind"] = append(score_category_map["four_of_a_kind"], i)
				//player.Hand_Rank = 2
			}else if check_full_house(hand){
				//score_category_map["full_house"]++
				score_category_map["full_house"] = append(score_category_map["full_house"], i)
				//player.Hand_Rank = 3
			}else if check_flush(hand){
				//score_category_map["flush"]++
				score_category_map["flush"] = append(score_category_map["flush"], i)
				//player.Hand_Rank = 4
			}else if check_stright(hand){
				//score_category_map["straight"]++
				score_category_map["straight"] = append(score_category_map["straight"], i)
				//player.Hand_Rank = 5
			}else if check_three_of_a_kind(hand){
				//score_category_map["three_of_a_kind"]++
				score_category_map["three_of_a_kind"] = append(score_category_map["three_of_a_kind"], i)
				//player.Hand_Rank = 6
			}else if check_two_pairs(hand){
				//score_category_map["two_pairs"] ++
				score_category_map["two_pairs"] = append(score_category_map["two_pairs"], i)
				//player.Hand_Rank = 7
			}else if check_pair(hand){
				//score_category_map["pair"]++
				score_category_map["pair"] = append(score_category_map["pair"], i)
				//player.Hand_Rank = 8
			}else{
				//score_category_map["nothing"]++
				score_category_map["nothing"] = append(score_category_map["nothing"], i)
				//player.Hand_Rank = 9
			}
		/* not yet complete - will modify return values so that it is easier to determine the winner
			when the best two hands belong to the same category */	
		}
	}
	return score_category_map
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
	if (hand[4].Rank - hand[0].Rank) == 4{
		return true
		}
	return false
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

/*func get_player_index_from_name(players []Player, name string)int{
	for i, p := range players{
		if p.Name == name{
		return i
		}	
	}
	return 0
} */

func showdown(players []Player, score_board map[string][]int)*Player{
	fmt.Printf("About to begin showdown \n")
	//hand_rank := 1
	var winner Player 
	for key, value := range score_board{
		if len(value) == 1{
			return &players[value[0]]
		}
		if len(value) > 1{
			best := 0
			if key == "straight_flush"{
				for i := 0; i < len(value); i++{
					if players[value[i]].Hand[0].Rank > best{
						best = players[value[i]].Hand[0].Rank
						winner = players[value[i]]
					} 
				}
			}else if key == "four_of_a_kind"{
				for i := 0; i < len(value); i++{
					rank := players[value[i]].find_four_of_kind_rank()
					if rank > best{
						winner = players[value[i]]
					}
				}
			}else if key == "full_house"{
				for i := 0; i < len(value); i++{
					rank := players[value[i]].find_three_of_kind_rank()
					if rank > best{
						winner = players[value[i]]
					}
				}
			}else if key == "three_of_a_kind"{
				for i := 0; i < len(value); i++{
					rank := players[value[i]].find_three_of_kind_rank()
					if rank > best{
						winner = players[value[i]]
					}
				}
			}else if key == "two_pairs"{
				best2 := 0
				for i := 0; i < len(value); i++{
					rank := players[value[i]].best_pair()
					if rank == best && best > 0{
						for j := 0; j < len(value); j++{
							rank2 := players[value[j]].second_best_pair()
							if rank2 > best2{
								winner = players[value[i]]
								}		
							}
						}else if rank > best && rank > best2{
							winner = players[value[i]]
						}
					}
			}else if key == "pair"{
				for i := 0; i < len(value); i++{
					rank := players[value[i]].best_pair()
					if rank > best{
						winner = players[value[i]]
						}
					}
			}else if key == "nothing"{
				for i := 0; i < len(value); i++{
					rank := players[value[i]].highest_card()
						if rank > best{
							winner = players[value[i]]
						}
					}
				}
			}
		}
	return &winner 
}



