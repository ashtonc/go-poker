# Group 15 - Five Card Draw

This is the repository for the term project of group 15. It will be an implemetation of poker (specifically five card draw).

To view our site, simply `vagrant up` and visit `localhost:8000/poker/`.

## Project Checkpoint

While our site doesn't completely work, here is a list of what is finished:

* Vagrant provisioning, including nginx and postgres
* Basic routing
* Basic URL structure
* Basic templates
* Basic page designs
* Poker game logic (in the testing phase)

What remains to be finished:

* Cohesive models for each game
* Full database integration
* Registration and authentication on the website
* Connection between the game logic and client page
* Websockets integration sending json representations of the game state to the client to prevent clients from having to refresh the page

## Technologies

* Routing: https://github.com/gorilla/mux
* Database: https://github.com/lib/pq
* Session management: https://github.com/gorilla/sessions (not yet)
* Websockets: https://github.com/gorilla/websocket (not yet)

## Server Log

* `vagrant ssh`
* `sudo su`
* `tmux attach -t server`
* TODO: set it up so the server log writes to /log or something
