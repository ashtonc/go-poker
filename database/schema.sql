CREATE TABLE player (
	id SERIAL PRIMARY KEY,
	username VARCHAR(32),
	description TEXT,
	password_hash VARCHAR(256),
	password_salt VARCHAR(256)
);

CREATE TABLE player_session (
	id SERIAL PRIMARY KEY,
	player_id INTEGER REFERENCES player (id),
	token VARCHAR(256),
	expiry_time TIMESTAMP
);

CREATE TABLE game (
	id SERIAL PRIMARY KEY,
	name TEXT,
	description TEXT,
	start_time TIMESTAMP,
	finish_time TIMESTAMP
);

CREATE TABLE game_round (
	id SERIAL PRIMARY KEY,
	game_id INTEGER REFERENCES game (id),
	stakes BIGINT,
	big_blind BIGINT,
	small_blind BIGINT,
	pot BIGINT
);

CREATE TABLE round_seat (
	id SERIAL PRIMARY KEY,
	player_id INTEGER REFERENCES player (id),
	position INTEGER
);

CREATE TABLE round_win (
	id SERIAL PRIMARY KEY,
	round_seat_id INTEGER REFERENCES round_seat (id),
	total_won BIGINT
);

CREATE TABLE round_action (
	id SERIAL PRIMARY KEY,
	round_seat_id INTEGER REFERENCES round_seat (id),
	pass BOOLEAN
);
