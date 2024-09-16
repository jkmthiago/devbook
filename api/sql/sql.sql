CREATE DATABASE IF NOT EXISTS devbook

USE devbook;

DROP TABLE IF EXISTS users

CREATE TABLE users (
	id serial NOT NULL,
	"name" varchar NOT NULL,
	nick varchar NOT NULL unique,
	email varchar NOT NULL,
	"password" varchar NOT NULL,
    createdin timestamp default current_timestamp,
	CONSTRAINT users_pk PRIMARY KEY (id)
);