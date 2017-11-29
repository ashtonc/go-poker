CREATE TABLE user (
	id SERIAL PRIMARY KEY,
	username VARCHAR(32),
	name VARCHAR(256),
	email VARCHAR(128),
	picture_slug VARCHAR(128),
	description TEXT,
	password_salt VARCHAR(256),
	password_hash VARCHAR(256)
);

CREATE TABLE session (
	id SERIAL PRIMARY KEY,
	token VARCHAR(256),
	expiry_time TIMESTAMP,
	user_id INTEGER REFERENCES user (id)
);

CREATE TABLE game_stakes (
	id SERIAL PRIMARY KEY,
	ante BIGINT,
	min_bet BIGINT,
	max_bet BIGINT
);

CREATE TABLE game_status (
	id SERIAL PRIMARY KEY,
	description TEXT
);

CREATE TABLE game (
	id SERIAL PRIMARY KEY,
	name VARCHAR(256),
	description TEXT,
	slug VARCHAR(64),
	stakes INTEGER REFERENCES game_stakes (id),
	status INTEGER REFERENCES game_status (id)
);

CREATE TABLE round_phase (
	id SERIAL PRIMARY KEY,
	description TEXT
);

CREATE TABLE game_round (
	id SERIAL PRIMARY KEY,
	pot BIGINT,
	game_id INTEGER REFERENCES game (id),
	phase INTEGER REFERENCES round_phase (id)
);

CREATE TABLE card (
	id SERIAL PRIMARY KEY,
	suit VARCHAR(8),
	rank VARCHAR(8)
);

CREATE TABLE round_hand (
	id SERIAL PRIMARY KEY,
	card_1 REFERENCES card (id),
	card_2 REFERENCES card (id),
	card_3 REFERENCES card (id),
	card_4 REFERENCES card (id),
	card_5 REFERENCES card (id)
);

CREATE TABLE game_player (
	id SERIAL PRIMARY KEY,
	position INTEGER,
	starting_cash BIGINT,
	final_cash BIGINT,
	hand INTEGER REFERENCES round_hand (id)
	user_id INTEGER REFERENCES user (id)
);

CREATE TABLE player_stats (
	user_id INTEGER REFERENCES user (id),
	best_hand INTEGER REFERENCES round_hand (id),
	total_hands INTEGER,
	total_cash BIGINT
);