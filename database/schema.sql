CREATE TABLE user (
	id SERIAL PRIMARY KEY,
	username VARCHAR(32),
	name VARCHAR(256),
	email VARCHAR(128),
	picture INT,
	description TEXT,
	password_hash VARCHAR(256),
	password_salt VARCHAR(256)
);

CREATE TABLE session (
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
	pot BIGINT
);

CREATE TABLE game_player (
	id SERIAL PRIMARY KEY,
	player_id INTEGER REFERENCES player (id),
	position INTEGER,
	total_won BIGINT
);

CREATE TABLE round_action (
	id SERIAL PRIMARY KEY,
	round_seat_id INTEGER REFERENCES round_seat (id),
	pass BOOLEAN
);
