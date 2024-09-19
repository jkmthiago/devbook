INSERT INTO users ("name",nick,email,"password",createdin) VALUES
	 ('User 1','user_1','user_1@test.com','$2a$10$11fJXi.7k24y/Iz07m7zXe5lxchhFTyZKCttUDFK8YwMZmzyVJ9Mm','2024-09-18 14:49:23.901167'),
	 ('User 2','user_2','user_2@test.com','$2a$10$11fJXi.7k24y/Iz07m7zXe5lxchhFTyZKCttUDFK8YwMZmzyVJ9Mm','2024-09-18 14:49:23.901167'),
	 ('User 3','user_3','user_3@test.com','$2a$10$11fJXi.7k24y/Iz07m7zXe5lxchhFTyZKCttUDFK8YwMZmzyVJ9Mm','2024-09-18 14:49:23.901167'),
	 ('User 4','user_4','user_4@test.com','$2a$10$11fJXi.7k24y/Iz07m7zXe5lxchhFTyZKCttUDFK8YwMZmzyVJ9Mm','2024-09-18 14:49:23.901167'),
	 ('Thiago Eleuterio','JKMTHIAGO','thiago@teste.com','$2a$10$tIDPpYjaYnbtrOh0NLSBAe4URqcQuRGJJTSB0JAP4Jz3z9csOcgvO','2024-09-18 18:47:21.524407');

INSERT INTO followers (user_id,follower_id) VALUES
	 (1,2),
	 (3,1),
	 (1,3),
	 (4,2),
	 (3,5),
	 (2,5),
	 (1,5);