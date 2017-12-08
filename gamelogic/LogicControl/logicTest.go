
package main

import(
	"math/rand"
	//"time"
	"fmt"
	"./gamelogic"
	"bufio"
  	"os"
  	"log"

)



func main(){
	gamelogic.Rand_init()
	//reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Welcome to Five Card Draw! (press 'enter' between messages to continue)")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	ante := 1
	min_bet := 0
	max_bet := 20


	game, gerr := gamelogic.GameInit(ante, min_bet, max_bet)
	if gerr != nil{
		fmt.Printf("Game object did not initialize! \n")
	}

	p1err := game.Join("Ashton", 100, 0)
	if p1err != nil{
		fmt.Printf("Player 1 was not added to the game! \n")

	}

	p2err := game.Join("Adam", 100, 1)
	if p2err != nil {
		fmt.Printf("Player 2 was not added to the game! \n")
		log.Fatal(p2err)
	}

	p3err := game.Join("Matthew", 100, 2)
	if p3err != nil {
		fmt.Printf("Player 3 was not added to the game! \n")
		log.Fatal(p3err)
	}

	dealterToken := 0

	fmt.Printf("1 \n")
	newErr := game.NewRound(dealterToken)
	if newErr != nil{
		log.Fatal(newErr)
	}
	fmt.Printf("2 \n")
	fmt.Printf("Dealer token: %d \n", game.Dealer_Token)
	fmt.Printf("Current Player is %s \n", game.Get_current_player_name())
	for{
		err, winner := game.Winner_check()
		if err != nil{
			log.Fatal(err)
		}
		if winner != nil{
			fmt.Printf("A winner is %s \n", winner.Name)
			break
		}
		if (game.Phase == 0 || game.Phase == 2 || game.Phase == 4){
			player := game.Get_current_player_name()
			decision := rand.Float32()
			if decision < 0.25{
				fmt.Printf("Bet \n")
				raise := rand.Intn(5)
				err := game.Bet(player, raise)
				if err != nil{
					log.Fatal(err)
				}
			}else if decision < 0.88{
				pindex, err := game.GetPlayerIndex(player)
				if err != nil{
					log.Fatal(err)
				}
				if game.Players[pindex].Bet == game.Current_Bet{
					fmt.Printf("Check \n")
					err = game.Check(player)
					if err != nil{
						log.Fatal(err)
					}
				}else{
					fmt.Printf("Call \n")
					err := game.Call(player)
					if err != nil{
						log.Fatal(err)
					}
				}
			}else{
				fmt.Printf("Fold \n")
				err := game.Fold(player)
				if err != nil{
					log.Fatal(err)
				}	
			}
		}else if (game.Phase == 1 || game.Phase == 3){
			player := game.Get_current_player_name()
			num_discard := rand.Intn(4)
			fmt.Printf("%s will discard %d cards \n", player, num_discard)
			var discard []int
			for i := 0; i <= num_discard; i++{
				discard = append(discard, i)
			}
			err := game.Discard(player, discard)
			if err != nil{
				log.Fatal(err)
			}
		}else{
			//game.Phase == 5
			winner := game.Showdown()
			fmt.Printf("A winner is %s", winner.Name)
			break
			}
		
	}
}

	/*player := game.Get_current_player_name()

	err := game.Bet(player, 7)
	fmt.Printf("%s bets %d \n", player, 7)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("3 \n")

	player = game.Get_current_player_name()
	err = game.Call(player)
	fmt.Printf("%s calls \n", player)

	if err != nil{
		log.Fatal(err)
	}	
	fmt.Printf("4 \n")

	player = game.Get_current_player_name()
	err = game.Bet(player, 4)
	fmt.Printf("%s bets 4 \n", player)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("5 \n")

	player = game.Get_current_player_name()
	err = game.Call(player)
	fmt.Printf("%s calls \n", player)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("6 \n")

	player = game.Get_current_player_name()
	err = game.Call(player)
	fmt.Printf("%s calls \n", player)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("7 \n")
	player = game.Get_current_player_name()
	err = game.Call(player)
	fmt.Printf("%s calls \n", player)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("8 \n")
	player = game.Get_current_player_name()
	err = game.Bet(player, 4)
	fmt.Printf("%s bests 4", player)
	if err != nil{
		log.Fatal(err)
	}

}
*/