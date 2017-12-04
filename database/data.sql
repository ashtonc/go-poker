INSERT INTO account (id, username, name, email, password, picture_slug, description, password_salt, password_hash) VALUES
	(1, 'adam', 'Adam Labecki', 'adam@email.com', 'nil', 'picture1.png', 'Description goes here...', 'salt', 'hash'),
	(2, 'ashton', 'Ashton Charbonneau', 'ashton@email.com', 'nil', 'picture2.png', 'Description goes here...', 'salt', 'hash'),
	(3, 'matthew', 'Matthew Tan', 'matthew@email.com', 'nil', 'picture3.png', 'Description goes here...', 'salt', 'hash'),
	(4, 'clayton', 'Clayton Jian', 'clayton@email.com', 'nil', 'picture4.png', 'Description goes here...', 'salt', 'hash'),
	(5, 'rimple', 'Rimpledeep Chahal', 'rimple@email.com', 'nil', 'picture5.png', 'Description goes here...', 'salt', 'hash'),
	(6, 'greg', 'Greg Baker', 'greg@email.com', 'nil', 'picture6.png', 'Description goes here...', 'salt', 'hash');

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
