## Overview

Poker is a commonly-played card game in which players try to assemble the "best" possible hand with the cards that they are given. Poker is most frequently played with a 52-card deck and up to 10 players. There are also multiple variants of poker, with the our project focusing on a variant called five-card draw.

Five-card draw is a variant of poker where players are dealt five cards. From here, up to five rounds (a full rotation of actions) take place before the match is finished. There are two types of rounds: discard rounds and betting rounds. In discard rounds, players can discard and replace cards in their hand. In betting rounds, players place bets, with the winner of the round collecting all money that was bet. After an inital betting round, players alternate between discard and betting rounds until only one player remains or three betting rounds have been completed.

In bettings rounds, each player takes turns deciding whether they want to wager as much as is already on the board (checking or calling), wager more than what is currently on the board (raising), or forfeit their hand and any chance of winning the match (folding). Calls are done to match an existing bet so that the player can stay in for another set of turns, while a raise is done to try and increase the amount of money at stake on the board. Folding is done if a player does not want to put any more money into the round.

The victory condition of the game is to obtain a hand that is considered stronger than the hands of the other players by the end of the match. Hands are ranked in order of strength from least to greatest by the following measure: high card, one pair of the same rank, two pairs of two different ranks, three-of-a-kind, straight (five cards of sequential rank), flush (five cards of the same suit), full house (a three-of-a-kind and a pair), four-of-a-kind, straight flush (a straight combined with a flush), and royal flush (a flush where the cards are numbered 10 to Ace).

When all of the rounds have finished, the remaining players reveal their hands, and the player with the best hand takes all of the money that has been wagered.


## Technologies / Implementation

The backend for our project is written in Go. Go's standard library ships with a built-in web server and database management tools. We used several libraries that extend this functionality:

* github.com/lib/pq allowed us to connect and use postgres as our backend database server.
* github.com/gorilla/mux adds functionality to the router included in Go's standard library. It improved the flexibility of our URL structure and use to define routes.
* github.com/gorilla/sessions was used alongside github.com/gorilla/securecookie for session management.
* github.com/gorilla/websockets was used to establish a real-time connection with clients so that they can play the game without refreshing their browser or long-polling.
* github.com/gorilla/csrf was used for a clean solution to cross-site request forgery.

Go was chosen because it has strong support for concurrency and asynchronous connections, something that is very useful for real-time games. Some of our code ends of being more verbose (for example, there is no built-in ORM), but it also gives us greater control over the functionality of the website.

Each page was first mocked up in HTML, then converted to a Go template. Go's templating language is quite powerful and allows us to prevent cross-site scripting attacks by commenting out dangerous expressions. Javascript is used to establish the websockets connection and update the page. We used jQuery to simplify the syntax and to add animations to the cards. It is possible to play the game without Javascript, users just need to keep track of the timer on their own and refresh the page to see updates. 

Nginx is set up to reverse-proxy the Go webserver and serve static content. The only static content are stylesheets (we use normalize.css to unify the look) and the javascript files for the game.


## Security

Go was designed to be very security-conscious. Their crypto library (with support for scrypt) is widely used and well-reviewed. Go's database functions are all designed with SQL injection in mind, allowing us to safely write functions as follows:

	func GetUser(env *models.Env, username string) (*models.User, error) {
		var user models.User
		sql := `SELECT name, email, picture FROM user WHERE username=$1;`
		row := env.Database.QueryRow(sql, username)
		err := row.Scan(&user.Name, &user.Email, &user.Picture)
		return &user, err
	}

Because we have this freedom, writing functions to cover up the lack of an ORM has been made much simpler.

Authorization middleware is used in tandem with our session management tools to ensure that users always have permission to view each page they view.


## Features

In our application, we have included a number of features that simulate five-card draw.

*User Profiles*: An authentication system is in place so that users can properly log in to play five-card draw and log out when they are done with their session. Users should be able to register an account, and those who create an account will have their own user profile in which they can display and edit basic information such as their name and profile picture. Each user account will also start off by default with a set amount of money which they can use to participate in matches. In addition, each player is associated with lifetime statistics such as wins and losses which will be displayed in the leaderboard as described below.

*Lobby System*: Our application displays a list of existing lobbies and permits the user to create a lobby of their own, as shown in Figure 1. The list of lobbies is not a static page and constantly updates itself to show users what rooms are available and how many players are in each room. Websockets are used here so that the server can send updates to users.

*Gameplay*: The gameplay itself is split up into several aspects, including the game logic, the timer, and the AI. In the overview, we discussed the manner in which five-card draw is played, and in our application we have written code in Go that directly handles the sequence of turns, actions that can be taken (i.e. discard rounds and wager rounds), and determining who the victor of a match is. A timer also prevents players from taking an indefinite period of time to make a move, and an AI will be able to follow the rules of five-card draw at a bare minimum. Figure 2 shows what a match of five-card draw would look like in our application.

**
Each player joins a table, where each table has six seats.
Each game has a specific ante (blinds might be added later), a minimum bet value, and maximum bet value. 
At the beginning of a round, each player pays the ante value, which is added to the pot. Next, each player is dealt five cards face down. Users can see the cards that they have been dealt, but only the back of other players cards.
After all players have examined their hands, betting begins starting from the dealer's left. 
Each player may choose to check (pass), call the current bet by matching the bet of the previous player, raise the current bet, or fold, whereby the player cuts her losses and withdraws from that round.
Betting continues to go around the table until each player has either passed or folded. All bets are added to the pot.
If more than one player remains after the first round of betting, each player may choose to discard a subset of the cards in her hand.
The dealer then deals new cards to replace those which were discarded. Upon examining the updated hands, players begin a new round of betting as before.
If more than one player remans after the second round of betting, a second discard and replace phase occurs, followed by a third and final betting phase.
After the final bets are made, the pot is awarded to the player with the strongest hand.
**

*Multiplayer*: Five-card draw is a game meant to be played with multiple other people, so being able to connect with others would be an expected part of our application. At the moment, we support up to 6 simultaneous players per game.

*Leaderboard*: A leaderboard exists that keeps track of user statistics. In this case, user statistics refer to information collected about a player's performance during their account's lifetime, such as the number of wins, losses, draws, folds, and amount of money held by that user. These statistics are obtained after the conclusion of each round, and the leaderboard is updated to reflect the changes in the player's performance.
