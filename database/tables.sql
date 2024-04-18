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