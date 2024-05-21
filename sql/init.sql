CREATE SEQUENCE revID
START WITH 238328
MINVALUE 238328;
CREATE SEQUENCE comID
START WITH 238328
MINVALUE 238328;

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
    join_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    followercount BIGINT NOT NULL DEFAULT 0,
    PRIMARY KEY(username)
);

CREATE TABLE reviews (
    username VARCHAR(30) NOT NULL,
    volumeID VARCHAR(30) NOT NULL,
    reviewID VARCHAR(20) DEFAULT genAlphaNum(NEXTVAL('revID')),
    content VARCHAR(3000) NOT NULL DEFAULT '',
    rating SMALLINT NOT NULL DEFAULT 0,
    post_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    likecount BIGINT NOT NULL DEFAULT 0,

    --Constraints
    CONSTRAINT chk_rating_change CHECK (rating >= 0 AND rating <= 10),
    CONSTRAINT fk_username FOREIGN KEY(username) REFERENCES users(username) ON DELETE CASCADE,

    PRIMARY KEY(reviewID)
);

CREATE TABLE comments (
    reviewID VARCHAR(20) NOT NULL,
    commentID VARCHAR(20) NOT NULL DEFAULT genAlphaNum(NEXTVAL('comID')),
    content VARCHAR(1000) NOT NULL,
    username VARCHAR(30) NOT NULL,
    post_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    likecount BIGINT NOT NULL DEFAULT 0,

    --Constraints
    CONSTRAINT fk_reviewid FOREIGN KEY(reviewID) REFERENCES reviews(reviewID) ON DELETE CASCADE,
    CONSTRAINT fk_username FOREIGN KEY(username) REFERENCES users(username) ON DELETE CASCADE,

    PRIMARY KEY(commentID)
);

CREATE TABLE user_likes_review (
    username VARCHAR(30) NOT NULL,
    reviewID VARCHAR(20) NOT NULL,

    --Constraints
    CONSTRAINT fk_reviewid FOREIGN KEY(reviewID) REFERENCES reviews(reviewID) ON DELETE CASCADE,
    CONSTRAINT fk_username FOREIGN KEY(username) REFERENCES users(username) ON DELETE CASCADE,

    PRIMARY KEY(username, reviewID)
);

CREATE TABLE user_likes_comment (
    username VARCHAR(30) NOT NULL,
    commentID VARCHAR(20) NOT NULL,

    --Constraints
    CONSTRAINT fk_commentID FOREIGN KEY(commentID) REFERENCES comments(commentID) ON DELETE CASCADE,
    CONSTRAINT fk_username FOREIGN KEY(username) REFERENCES users(username) ON DELETE CASCADE,

    PRIMARY KEY(username, commentID)
);

CREATE TABLE user_follows_user (
    follower VARCHAR(30) NOT NULL,
    followed VARCHAR(30) NOT NULL,

    --Constraints
    CONSTRAINT fk_follower FOREIGN KEY(follower) REFERENCES users(username) ON DELETE CASCADE,
    CONSTRAINT fk_followed FOREIGN KEY(followed) REFERENCES users(username) ON DELETE CASCADE,

    PRIMARY KEY(follower, followed)
);

--Indexes
CREATE INDEX VolumeReviewIndex ON reviews(volumeID);
CREATE INDEX UserReviewIndex ON reviews(username);