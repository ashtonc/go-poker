package gamelogic

import(
	_"time"
	_"bufio"
  	_"os"  
 // 	"time"
  	"fmt"
)


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
	var winner *Player
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
						return winner
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
				return winner
			}
		}
	} else {
		contenders := score_board["nothing"]
		index := find_best_nothing(contenders, g.Players)
		winner = g.Players[index]
	}
	return winner
}

func find_best_nothing(indexes []int, players []*Player) int {

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