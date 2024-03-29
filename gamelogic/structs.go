package gamelogic

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	//	"sort"
	"strings"
	"time"
)

func Init_card_cat() ([]string, []string) {
	cardTypes := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "jack", "queen", "king", "ace"}
	suites := []string{"hearts", "spades", "clubs", "diamonds"}
	return cardTypes, suites
}

func getIndex(array []string, item string) int {
	for i := 0; i < len(array); i++ {
		if array[i] == item {
			return i
		}
	}
	return -1
}

func Rand_init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

/*classes:   */

type Game struct {
	Name           string     `json:"-"`
	Slug           string     `json:"-"`
	Stakes         GameStakes `json:"-"`
	Phase          int        `json:"-"`
	Pot            int        `json:"pot"`
	Deck           []Card     `json:"-"`
	Seats          [6]*Seat   `json:"seats"`
	Players        []*Player  `json:"-"`
	Sitters        []*Player  `json:"-"`
	Current_Player string     `json:"-"`
	Current_Bet    int        `json:"-"`
	Bet_Counter    int        `json:"-"`
	Dealer_Token   int        `json:"-"`
	Timer          time.Timer `json:"-"`
	Winner         *Player    `json:"-"`
}

/* --> do this with a set of constants **
Phases:
	0 -> betting 1
	1 -> draw 1
	2 -> betting 2
	3 -> draw 2
	4 -> betting 2
	5 -> showdown
*/
type GameStakes struct {
	Ante   int
	MaxBet int
	MinBet int
}

type Seat struct {
	Number   int     `json:"number"`
	Occupied bool    `json:"occupied"`
	Occupier *Player `json:"player,omitempty"`
	Winner   bool    `json: "Winner"`
}

type Player struct {
	Username    string  `json:"username"`
	Name        string  `json:"-"`
	PictureSlug string  `json:"pictureslug"`
	Money       int     `json:"money"`
	Hand        []Card  `json:"hand"`
	Folded      bool    `json:"folded"`
	Called      bool    `json:"-"`
	Discarded   bool    `json:"-"`
	Bet         int     `json:"-"`
	Hand_Rank   int     `json:"-"`
	Card_Hist   [14]int `json:"-"`
	Seat        int     `json:"-"`
}

/*
type Card struct {
	Face string `json:"face"`
	Suit string `json:"suit"`
	Rank int    `json:"-"`
}
*/

/*
func getIndex(array []string, item string) int {
	for i := 0; i < len(array); i++ {
		if array[i] == item {
			return i
		}
	}
	return -1
}
*/

func GameInit(ante int, min_bet int, max_bet int) (*Game, error) {
	game := new(Game)

	if game == nil {
		return nil, errors.New("Game failed to initiate.")
	}

	game.Stakes.Ante = ante
	game.Stakes.MinBet = min_bet
	game.Stakes.MaxBet = max_bet
	game.Dealer_Token = -1

	for i := 0; i < 6; i++ {
		var seat Seat
		seat.Number = i + 1
		seat.Occupied = false
		seat.Winner = false
		game.Seats[i] = &seat
	}

	return game, nil
}

func (p *Player) pay_bet(amount int, pot int) int {
	p.Money -= amount
	pot += amount
	fmt.Printf("%s places %d dollars into the pot \n", p.Name, amount)
	fmt.Printf("The pot now contains %s dollars \n", pot)
	return pot
}

func (p *Player) stay_in(difference int) bool {
	reader := bufio.NewReader(os.Stdin)
	if difference > p.Money {
		p.Folded = true
		fmt.Print("%s is unable to meet the raised bet and is out of the game \n", p.Name)
		return false
	}
	fmt.Printf("The bet has been rased by %d \n", difference)
	fmt.Printf("Will %s stay in the game? (Y/N) \n", p.Name)
	var input string
	input, err := reader.ReadString('\n')
	fmt.Println(err)
	input = strings.Replace(input, "\r\n", "", -1)
	stay := false
	if input == "N" || input == "n" {
		stay = false
	} else if input == "Y" || input == "y" {
		stay = true
	}
	/* show_fun() letting player know that he or she is out of the game */
	/* stay = show_func(difference) player p gets a pop up asking if he or she wishes to keep up with
	the latest bet in order to remain in the game */
	if stay == false {
		p.Folded = true
		return false
	} else {
		return true
	}
}

func (p *Player) show_hand() {
	fmt.Printf("%s's Hand: \n", p.Name)
	for i, crd := range p.Hand {
		fmt.Printf("%d %s of %s \n", i, crd.Face, crd.Suit)
	}
}

func (p *Player) sort_hand_by_rank() {
	fmt.Printf("About to sort %s hand by rank \n", p.Name)
	hand := p.Hand
	/*sort.Slice(hand, func(i, j int) bool {
		return hand[i].Rank < hand[j].Rank
	})*/
	p.Hand = hand
}

func (p *Player) find_four_of_kind_rank() int {
	for k, v := range p.Card_Hist {
		if v == 4 {
			fmt.Printf("For of kind rank: %d \n", k)
			return k
		}
	}
	return 0
}

func (p *Player) find_three_of_kind_rank() int {
	for k, v := range p.Card_Hist {
		if v == 3 {
			return k
		}
	}
	return 0
}

func (p *Player) best_pair() int {
	for k := len(p.Card_Hist) - 1; k >= 0; k-- {
		if p.Card_Hist[k] == 2 {
			return k
		}
	}
	return 0
}

func (p *Player) second_best_pair() int {
	for k := range p.Card_Hist {
		if p.Card_Hist[k] == 2 {
			return k
		}
	}
	return 0
}

/*func (p *Player)discarded_hand(discard_index []int){

	discard := make([]Card, 5)
	for _, d := range discard_index {
		discard = append(discard, p.Hand[d])
	}

	discarded_hand := make([]Card, 5)
	check := true
	for _, c := range p.hand{
		for _, d := range p.hand{
			if c.Face == d.Face && c.Suit == d.Suit{
				check = false
			}
		if check == true{
			discarded_hand = append(discarded_hand, c)
			}
		}
	}
} */

//func (p *Player) remove_card(Card){
//	index := getIndex(p.Hand, card)
//	p.Hand = append(p.Hand[:index], p.Hand[index+1:]...)
//}

type Card struct {
	Face string `json: "face"`
	Suit string `json: "suit"`
	Rank int    `json: "rank"`
}

func newCard(face string, suit string, cardTypes []string) *Card {
	crd := new(Card)
	crd.Face = face
	crd.Suit = suit
	rank := getIndex(cardTypes, face)
	crd.Rank = rank
	return crd
}
