CREATE TABLE users (
    username VARCHAR(30) NOT NULL CHECK (username = LOWER(username)),
    password CHAR(60) NOT NULL,
    email VARCHAR(100) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    join_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(username)
);