CREATE DATABASE devbook;

\c devbook;

DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
	id serial NOT NULL,
	"name" varchar(256) NOT NULL,
	nick varchar(30) NOT NULL unique,
	email varchar(256) NOT NULL unique,
	"password" varchar(256) NOT NULL,
    createdin timestamp default current_timestamp,
	CONSTRAINT users_pk PRIMARY KEY (id)
);

CREATE TABLE followers (
	user_id int not null,
	CONSTRAINT fk_user
	foreign key (user_id)
	references users(id)
	on delete cascade,

	follower_id int not null,
	CONSTRAINT fk_follower
	foreign key (follower_id)
	references users(id)
	on delete cascade,	

	CONSTRAINT followers_pk primary key (user_id, follower_id)
);

CREATE TABLE posts (
	id serial NOT NULL,
	CONSTRAINT posts_pk PRIMARY KEY (id),
	title varchar(50) NOT NULL unique,
	content varchar(500) NOT NULL unique,

	autor_id int not null,
	CONSTRAINT fk_autor
	foreign key (autor_id)
	references users(id)
	on delete cascade,

	likes int default 0,
    createdin timestamp default current_timestamp
)