CREATE TABLE IF NOT EXISTS users (
    user_id VARCHAR(36) PRIMARY KEY NOT NULL UNIQUE,
    nickname VARCHAR(20) NOT NULL UNIQUE,
    age INTEGER NOT NULL,
    gender CHAR(1) CHECK (gender IN ('M', 'F')) NOT NULL,
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(20) NOT NULL,
    email VARCHAR(30) NOT NULL,
    -- OPTIMIZE: Salt
    password_hash BYTEA NOT NULL,
    session TIMESTAMP
);

CREATE TABLE IF NOT EXISTS posts (
    post_id VARCHAR(36) PRIMARY KEY NOT NULL UNIQUE,
    user_id VARCHAR(36) NOT NULL,
    title VARCHAR(50) NOT NULL UNIQUE,
    content TEXT,
    created_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE IF NOT EXISTS comments (
    comment_id VARCHAR(36) PRIMARY KEY NOT NULL UNIQUE,
    user_id VARCHAR(36) NOT NULL,
    post_id VARCHAR(36) NOT NULL,
    content TEXT,
    created_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (post_id) REFERENCES posts(post_id)
);

CREATE TABLE IF NOT EXISTS messages (
    message_id VARCHAR(36) PRIMARY KEY NOT NULL UNIQUE,
    sender_id VARCHAR(36) NOT NULL,
    receiver_id VARCHAR(36) NOT NULL,
    content TEXT,
    created_at TIMESTAMP
)
