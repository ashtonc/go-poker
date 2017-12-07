INSERT INTO account (username, name, email, picture_slug, description, password_hash) VALUES
	('adam', 'Adam Labecki', 'adam@email.com', 'picture1.png', 'Description goes here...', '$2a$04$Hca3lE0CmRxwluksV5y59eCvW7MNRppOp493fiS0asRceHUJL5.Wy'),
	('ashton', 'Ashton Charbonneau', 'ashton@email.com', 'picture2.png', 'Description goes here...', '$2a$04$Hca3lE0CmRxwluksV5y59eCvW7MNRppOp493fiS0asRceHUJL5.Wy'),
	('matthew', 'Matthew Tan', 'matthew@email.com', 'picture3.png', 'Description goes here...', '$2a$04$Hca3lE0CmRxwluksV5y59eCvW7MNRppOp493fiS0asRceHUJL5.Wy'),
	('clayton', 'Clayton Jian', 'clayton@email.com', 'picture4.png', 'Description goes here...', '$2a$04$Hca3lE0CmRxwluksV5y59eCvW7MNRppOp493fiS0asRceHUJL5.Wy'),
	('rimple', 'Rimpledeep Chahal', 'rimple@email.com', 'picture5.png', 'Description goes here...', '$2a$04$Hca3lE0CmRxwluksV5y59eCvW7MNRppOp493fiS0asRceHUJL5.Wy'),
	('greg', 'Greg Baker', 'greg@email.com', 'picture6.png', 'Description goes here...', '$2a$04$Hca3lE0CmRxwluksV5y59eCvW7MNRppOp493fiS0asRceHUJL5.Wy');

INSERT INTO game_status (id, description) VALUES
	(1, 'open'),
	(2, 'closed');

INSERT INTO game_stakes (id, ante, min_bet, max_bet) VALUES
	(1, 10, 10, 1000),
	(2, 100, 100, 10000),
	(3, 1, 1, -1);

INSERT INTO game (id, name, slug, game_status, stakes) VALUES
	(1, 'No Limit 1', 'no-limit-1', 1, 3),
	(2, 'Closed 10', 'closed-10', 2, 1),
	(3, 'Limit 10', 'limit-10', 1, 1),
	(4, 'Limit 100', 'limit-100', 1, 2);
