# Group 15 - Five Card Draw

This is the repository for the term project of group 15. It will be an implementation of poker (specifically five-card-draw).

To view our site, simply clone the respository, `vagrant up`, and visit `localhost:8000/poker/`.

The view the server log, `vagrant ssh` into the server, `sudo su` to switch to root, and `tmux a` to attach the server console.

## To do

* Site
	* [**DONE**] Add games to the database
	* [**DONE**] View games from the lobby (fix it) - assigned a seat?
* Websockets
	* [**DONE**] Client can establish a connection to the server (websockets, js + go)
	* Client can send empty messages of each kind (js)
	* Server can recieve these messages (go)
	* Server can send game states (go) to all clients
	* Client can recieve game states (js)
	* Page is updated based on those game states (js)
	* Client can send content in their messages (js)
* Games
	* Updated database after each round
	* Make sure the game methods work?
	* Design it properly
* Fixes
	* Store password hashes instead of plain text (wtf)
	* Maybe leaderboard (less important)
	* Make registration work
	* Make sessions work

## Project Checkpoint

The majority of the work that has been completed on our website is not yet integrated together. We have designs that aren't yet visible, we have game logic that can't yet be interacted with, we have pages that can't be visited, etc. With that said, here is a list of what has been worked on:

* Vagrant provisioning, including nginx and postgres
* Basic routing and our basic URL scheme
* Basic templates for all pages (some are empty)
* Basic page designs (see the pages and designs directories)
* Basic database schema (missing space for information about each game)
* Poker game logic (being tested)

What remains to be finished:

* Integrating the design CSS into the actual pages
* Cohesive models for each game so that it can be managed in memory
* Full database integration (currently only viewing the user pages right now)
* Registration and authentication on the website
* Creation of the game page template (will be a nightmare)
* Connection between the game page and the game logic 
* Websockets integration sending json representations of the game state to the client to prevent clients from having to refresh the page

The following pages can be viewed on our website:

* `/poker/`: the home page
* `/poker/game/`: not yet finished
* `/poker/user/{any username with a-z, A-Z, 0-9, -, _, .}`: will currently redirect you to the home page because it cannot find that user in the database yet
* `/poker/user/{any username with a-z, A-Z, 0-9, -, _, .}/edit`: filled with dummy data for now but you can get the idea
* `/poker/leaderboard/`: example page is up
* `/poker/login/`: example page is up
* `/poker/register/`: example page is up

Repository information:

* `database` contains files that connect to the postgres database and retrieve data from that database
* `design` (and `design/pages`) contains our mockups of each page that will exist
* `gamelogic` contains the poker logic that will run the game on the server
* `handlers` contains the handlers for each individual page
* `models` will contain the structs/models for the other packages
* `sfx` contains preliminary sound effects to be integrated into the game
* `static` contains static content server at `/poker/assets/` on the server
* `templates` contains our templates
* `poker.go` is the main file that is run
* The other files are related to provisioning

## Technologies

* Routing: https://github.com/gorilla/mux
* Database: https://github.com/lib/pq
* Session management: https://github.com/gorilla/sessions (not yet)
* Websockets: https://github.com/gorilla/websocket
