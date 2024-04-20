package models

var QUERY = `
	CREATE TABLE IF NOT EXISTS users (
		user_id CHAR(36) NOT NULL PRIMARY KEY UNIQUE,
		nickname VARCHAR(20) NOT NULL UNIQUE,
		age INTEGER NOT NULL,
		gender TEXT CHECK (gender IN ('M', 'F')) NOT NULL,
		first_name VARCHAR(20) NOT NULL,
		last_name VARCHAR(20) NOT NULL,
		email VARCHAR(30) NOT NULL,
		password_hash BYTEA NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	
	CREATE TABLE IF NOT EXISTS posts (
		post_id CHAR(36) NOT NULL PRIMARY KEY UNIQUE,
		user_id CHAR(36) NOT NULL,
		title VARCHAR(50) NOT NULL UNIQUE,
		content TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(user_id)
	);
	
	CREATE TABLE IF NOT EXISTS comments (
		comment_id CHAR(36) NOT NULL PRIMARY KEY UNIQUE,
		user_id CHAR(36) NOT NULL,
		post_id CHAR(36) NOT NULL,
		content TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(user_id),
		FOREIGN KEY (post_id) REFERENCES posts(post_id)
	);
`
