CREATE TABLE player (  
	id SERIAL PRIMARY KEY,
	username VARCHAR(32),
	description TEXT,
	password_hash VARCHAR(256),
	password_salt VARCHAR(256),
);

CREATE TABLE player_session (
	player_id integer REFERENCES player (id),
);

CREATE TABLE game (  
	id SERIAL PRIMARY KEY,
	name
	description
	start_time
	finish_time
);

CREATE TABLE game_round (  
	id SERIAL PRIMARY KEY,
	big_blind
	small_blind
	pot
);

CREATE TABLE round_win (
	id SERIAL PRIMARY KEY,
	round_seat_id integer REFERENCES round_seat (id),
	total_won
);

CREATE TABLE round_seat (
	id SERIAL PRIMARY KEY,
	player_id integer REFERENCES player (id),
	position
	card1
	card2
	card3
	card4
	card5
);

CREATE TABLE round_action (  
	id SERIAL PRIMARY KEY,
	round_seat_id integer REFERENCES round_seat (id),
	pass
	bet
	trade
);
