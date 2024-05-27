CREATE SEQUENCE revID
START WITH 238328
MINVALUE 238328;
CREATE SEQUENCE comID
START WITH 238328
MINVALUE 238328;
CREATE SEQUENCE uID
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
    username VARCHAR(30) UNIQUE NOT NULL CHECK (username = LOWER(username)),
    userID VARCHAR(20) NOT NULL DEFAULT genAlphaNum(NEXTVAL('uID')),
    password CHAR(60) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    description VARCHAR(500) NOT NULL DEFAULT '',
    join_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    followercount BIGINT NOT NULL DEFAULT 0,
    followingcount BIGINT NOT NULL DEFAULT 0,
    PRIMARY KEY(userID)
);

CREATE TABLE reviews (
    userID VARCHAR(20) NOT NULL,
    volumeID VARCHAR(30) NOT NULL,
    reviewID VARCHAR(20) NOT NULL DEFAULT genAlphaNum(NEXTVAL('revID')),
    content VARCHAR(3000) NOT NULL DEFAULT '',
    rating SMALLINT NOT NULL DEFAULT 0,
    post_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    likecount BIGINT NOT NULL DEFAULT 0,

    --Constraints
    CONSTRAINT chk_rating_change CHECK (rating >= 0 AND rating <= 10),
    CONSTRAINT fk_userid FOREIGN KEY(userID) REFERENCES users(userID) ON DELETE CASCADE,

    PRIMARY KEY(reviewID)
);

CREATE TABLE comments (
    reviewID VARCHAR(20) NOT NULL,
    commentID VARCHAR(20) NOT NULL DEFAULT genAlphaNum(NEXTVAL('comID')),
    content VARCHAR(1000) NOT NULL,
    userID VARCHAR(20) NOT NULL,
    post_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    likecount BIGINT NOT NULL DEFAULT 0,

    --Constraints
    CONSTRAINT fk_reviewid FOREIGN KEY(reviewID) REFERENCES reviews(reviewID) ON DELETE CASCADE,
    CONSTRAINT fk_userid FOREIGN KEY(userID) REFERENCES users(userID) ON DELETE CASCADE,

    PRIMARY KEY(commentID)
);

CREATE TABLE user_likes_review (
    userid VARCHAR(20) NOT NULL,
    like_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    reviewid VARCHAR(20) NOT NULL,

    --Constraints
    CONSTRAINT fk_reviewid FOREIGN KEY(reviewID) REFERENCES reviews(reviewID) ON DELETE CASCADE,
    CONSTRAINT fk_userid FOREIGN KEY(userID) REFERENCES users(userID) ON DELETE CASCADE,

    PRIMARY KEY(userID, reviewID)
);

CREATE TABLE user_likes_comment (
    userid VARCHAR(20) NOT NULL,
    like_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    commentid VARCHAR(20) NOT NULL,

    --Constraints
    CONSTRAINT fk_commentID FOREIGN KEY(commentID) REFERENCES comments(commentID) ON DELETE CASCADE,
    CONSTRAINT fk_userid FOREIGN KEY(userID) REFERENCES users(userID) ON DELETE CASCADE,

    PRIMARY KEY(userID, commentID)
);

CREATE TABLE user_follows_user (
    followerid VARCHAR(20) NOT NULL,
    follow_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    followedid VARCHAR(20) NOT NULL,

    --Constraints
    CONSTRAINT fk_follower FOREIGN KEY(followerID) REFERENCES users(userID) ON DELETE CASCADE,
    CONSTRAINT fk_followed FOREIGN KEY(followedID) REFERENCES users(userID) ON DELETE CASCADE,
    CHECK (followerID <> followedID),

    PRIMARY KEY(followerID, followedID)
);

--Indexes
CREATE INDEX VolumeReviewIndex ON reviews(volumeID);
CREATE INDEX UserReviewIndex ON reviews(userID);