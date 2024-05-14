CREATE SEQUENCE base62id
    INCREMENT BY 1
    START WITH 1
    MINVALUE 1
    NO MAXVALUE
    CACHE 1;

CREATE OR REPLACE FUNCTION genAlphaNum(num BIGINT)
RETURNS VARCHAR(255) AS $$

DECLARE 

    base_62 VARCHAR(62) := '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz';
    result VARCHAR(20) := '';
    remainder INTEGER;

BEGIN
    IF num = 0 THEN
        RETURN 0;
    END IF;

    WHILE num > 0 LOOP
        remainder := num % 62;
        result := substr(base_62, remainder + 1, 1) || result;
        num := num/62;
    END LOOP;

    RETURN result;
END;

$$ LANGUAGE plpgsql;

CREATE TABLE users (
    username VARCHAR(30) NOT NULL CHECK (username = LOWER(username)),
    password CHAR(60) NOT NULL,
    email VARCHAR(100) NOT NULL,
    description VARCHAR(500) NOT NULL DEFAULT '',
    join_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(username)
);

CREATE TABLE reviews (
    username VARCHAR(30) NOT NULL,
    worksID VARCHAR(30) NOT NULL,
    reviewID VARCHAR(20) DEFAULT genAlphaNum(NEXTVAL('base62id')),
    content VARCHAR(500),
    rating SMALLINT,
    post_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    --Potentially should make an index on username

    --Constraints
    CONSTRAINT chk_rating_change CHECK (rating >= 0 AND rating <= 10),
    CONSTRAINT fk_username FOREIGN KEY(username) REFERENCES users(username) ON DELETE CASCADE,

    PRIMARY KEY(reviewID)
);

CREATE TABLE replies (
    reviewID VARCHAR(20) NOT NULL,
    replyID VARCHAR(20) DEFAULT genAlphaNum(NEXTVAL('base62id')),
    content VARCHAR(500) NOT NULL,
    username VARCHAR(30) NOT NULL,

    --Constraints
    CONSTRAINT fk_reviewid FOREIGN KEY(reviewID) REFERENCES reviews(reviewID) ON DELETE CASCADE,
    CONSTRAINT fk_username FOREIGN KEY(username) REFERENCES users(username) ON DELETE CASCADE,

    PRIMARY KEY(replyID)
);

CREATE TABLE user_likes_reviews (
    username VARCHAR(30) NOT NULL,
    reviewID VARCHAR(20) NOT NULL,

    --Constraints
    CONSTRAINT fk_reviewid FOREIGN KEY(reviewID) REFERENCES reviews(reviewID) ON DELETE CASCADE,
    CONSTRAINT fk_username FOREIGN KEY(username) REFERENCES users(username) ON DELETE CASCADE,

    PRIMARY KEY(username, reviewID)
);

CREATE TABLE user_likes_reply (
    username VARCHAR(30) NOT NULL,
    replyID VARCHAR(20) NOT NULL,

    --Constraints
    CONSTRAINT fk_reviewid FOREIGN KEY(replyID) REFERENCES replies(replyID) ON DELETE CASCADE,
    CONSTRAINT fk_username FOREIGN KEY(username) REFERENCES users(username) ON DELETE CASCADE,

    PRIMARY KEY(username, replyID)
);

CREATE TABLE user_follows_user (
    follower VARCHAR(30) NOT NULL,
    followed VARCHAR(30) NOT NULL,

    --Constraints
    CONSTRAINT fk_follower FOREIGN KEY(follower) REFERENCES users(username) ON DELETE CASCADE,
    CONSTRAINT fk_followed FOREIGN KEY(followed) REFERENCES users(username) ON DELETE CASCADE,

    PRIMARY KEY(follower, followed)
);