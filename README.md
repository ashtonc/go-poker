# Group 15 - Five Card Draw

This is the repository for the term project of group 15. It is an implementation of poker (specifically five-card-draw).

To view our site, simply clone the respository, `vagrant up`, and visit `localhost:8000/poker/`.

The view the server log, `vagrant ssh` into the server, `sudo su` to switch to root, and `tmux a` to attach the server console.

## Features

Our site doesn't currently work very well. While it opens and there are a few things that can be done, it is far from a working site.

### Working Features

Our site has the following working features:

* Users can register and log in.
* Users can edit their account and view public information about other accounts.
* There exists a game lobby where instantiated games (currently just taken from the database) can be viewed or joined.
* Each button on the game page makes a request that updates the state of the game, then the user is redirected back to the game page. Consider this a win for users that have Javascript disabled.
* The game template will show the current state of the game, including player cards and pots, if the game object that is passed in is in the middle of a game.

### Hidden/Broken Features

Our site has a number of things that are in the code but are not reflected on the actual site itself.

* There exists a leaderboard (/poker/leaderboard/) but it has been taken off the navigation bar because game statistics are not currently saved in the database. Theoretically, it shows users sorted by their total cash (or perhaps best hand). Currently, there are no entries, and it simply says that there is nobody on the leaderboard.
* Working game logic. Using api calls to individual instantiations of a game, it is possible to run through a full round of five-card draw. This isn't well reflected on the game page because the code used to update the page with information about the game hasn't been written.
* Websocket connection. It is possible to send a JSON representation of the game state to clients, and possible for the server to receive JSON representation of game moves, but this code was taken out because the page could not be updated using this information.
* Users have an image associated with their account that should be displayed on the game page. The form for submitting that image was not finished, but most of the supporting code exists to handle those images.
* Not particularly valuable, but the site root can be changed from `/poker` to something else if desired.
* Bcrypt runs a bit slow on the virtual machine.

### Missing features

Many things were not completed for this submission.

* Page updates. Given a JSON game state, the client should be able to update the page without refreshing.
* Complete websocket methods. Though it is possible to send a complete gamestate to the client through our websocket connection, the client page hasn't been updated such that the buttons on the page would send their requests or interactions.
* Proper css. Our pages look pretty messy right now.
* Some proper form validation and some error handling is missing.
* A number of optimizations haven't really been made (proper caching, minification, etc).
* Some security features are lacking, most notably CSRF. SQL injection and XSS should not be an issue, however.
* Proper server log fixing. Some errors are printed without context, which isn't helpful.
* Player stats. Joining a game should reduce the players wallet by the buyin amount, and finishing the game should update the the players total amount of money and potentially their best historical hand.
* Proper page headers. Different pages should have different headers (vary would be useful, for example).
* Proper error handling. 404 pages should be styled, and some pages should have different http responses than they currently have.

## Authorization

It is easy to register for an account, but there also exist a few premade accounts that can be accessed:

* user: ashton; pw: 470
* user: adam; pw: 470
* user: matthew; pw: 470
* user: clayton; pw: 470
* user: rimple; pw: 470
* user: greg; pw: 470

## External Libraries

We used the following external libraries:

* Routing: https://github.com/gorilla/mux
* Database: https://github.com/lib/pq
* Cookies: https://github.com/gorilla/securecookie
* Bcrypt: https://godoc.org/golang.org/x/crypto/bcrypt
* Websockets: https://github.com/gorilla/websocket
* UUID generation for websockets: https://github.com/satori/go.uuid
* Convenient JSON parser: https://github.com/tidwall/gjson

We also referenced several code bases:

* Golang websocket drawing app: https://outcrawl.com/realtime-collaborative-drawing-go/
* HTML5 playing cards: https://github.com/pakastin/deck-of-cards
