INSERT INTO users ("name",nick,email, "password") VALUES
	 ('User 1','user_1','user_1@test.com','$2a$10$11fJXi.7k24y/Iz07m7zXe5lxchhFTyZKCttUDFK8YwMZmzyVJ9Mm'),
	 ('User 2','user_2','user_2@test.com','$2a$10$11fJXi.7k24y/Iz07m7zXe5lxchhFTyZKCttUDFK8YwMZmzyVJ9Mm'),
	 ('User 3','user_3','user_3@test.com','$2a$10$11fJXi.7k24y/Iz07m7zXe5lxchhFTyZKCttUDFK8YwMZmzyVJ9Mm'),
	 ('User 4','user_4','user_4@test.com','$2a$10$11fJXi.7k24y/Iz07m7zXe5lxchhFTyZKCttUDFK8YwMZmzyVJ9Mm'),
	 ('Thiago Eleuterio','JKMTHIAGO','thiago@teste.com','$2a$10$tIDPpYjaYnbtrOh0NLSBAe4URqcQuRGJJTSB0JAP4Jz3z9csOcgvO');

INSERT INTO followers (user_id,follower_id) VALUES
	 (1,2),
	 (3,1),
	 (1,3),
	 (4,2),
	 (3,5),
	 (2,5),
	 (1,5);

INSERT INTO posts (title, "content", autor_id) VALUES
	('Title nº2', 'This is a test message for the post nº2 ', 5)
	('Title nº3', 'This is a test message for the post nº3 ', 2)
	('Title nº4', 'This is a test message for the post nº4 ', 4)
	('Title nº5', 'This is a test message for the post nº5 ', 1)
	('Title nº6', 'This is a test message for the post nº6 ', 1)
	('Title nº7', 'This is a test message for the post nº7 ', 1)
	('Title nº8', 'This is a test message for the post nº8 ', 3)
	('Title nº9', 'This is a test message for the post nº9 ', 1)