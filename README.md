# Group 15 - Five Card Draw

This is the repository for the term project of group 15. It will be an implemetation of poker (specifically five card draw).

To view our site, simply `vagrant up` and visit `localhost:8000/poker/`.

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
* Websockets: https://github.com/gorilla/websocket (not yet)

## Server Log

* `vagrant ssh`
* `sudo su`
* `tmux attach -t server` (or just `tmux a`)
* TODO: set it up so the server log writes to /log or something
