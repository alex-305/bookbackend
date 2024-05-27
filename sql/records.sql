--Users
INSERT INTO users(username, password, email, description) VALUES('ted', '$2a$10$NxMO/aiddukYPXCM7/riCeKETTOUYEcGyBfMDxBpwxVZIU/eOPE12', 'ted@mail.com', 'Hi my name is Ted and I am the primary testing account of bookdom.');
INSERT INTO users(username, password, email) VALUES('john', '$2a$10$NxMO/aiddukYPXCM7/riCeKETTOUYEcGyBfMDxBpwxVZIU/eOPE12', 'john@mail.com');
INSERT INTO users(username, password, email) VALUES('fred', '$2a$10$NxMO/aiddukYPXCM7/riCeKETTOUYEcGyBfMDxBpwxVZIU/eOPE12', 'fred@mail.com');
INSERT INTO users(username, password, email) VALUES('ned', '$2a$10$NxMO/aiddukYPXCM7/riCeKETTOUYEcGyBfMDxBpwxVZIU/eOPE12', 'ned@mail.com');
INSERT INTO users(username, password, email) VALUES('allen', '$2a$10$NxMO/aiddukYPXCM7/riCeKETTOUYEcGyBfMDxBpwxVZIU/eOPE12', 'allen@mail.com');
INSERT INTO users(username, password, email) VALUES('adam', '$2a$10$NxMO/aiddukYPXCM7/riCeKETTOUYEcGyBfMDxBpwxVZIU/eOPE12', 'adam@mail.com');
INSERT INTO users(username, password, email) VALUES('jacob', '$2a$10$NxMO/aiddukYPXCM7/riCeKETTOUYEcGyBfMDxBpwxVZIU/eOPE12', 'jacob@mail.com');
INSERT INTO users(username, password, email) VALUES('james', '$2a$10$NxMO/aiddukYPXCM7/riCeKETTOUYEcGyBfMDxBpwxVZIU/eOPE12', 'james@mail.com');
INSERT INTO users(username, password, email) VALUES('miles', '$2a$10$NxMO/aiddukYPXCM7/riCeKETTOUYEcGyBfMDxBpwxVZIU/eOPE12', 'miles@mail.com');
INSERT INTO users(username, password, email) VALUES('peter', '$2a$10$NxMO/aiddukYPXCM7/riCeKETTOUYEcGyBfMDxBpwxVZIU/eOPE12', 'peter@mail.com');
INSERT INTO users(username, password, email) VALUES('jane', '$2a$10$NxMO/aiddukYPXCM7/riCeKETTOUYEcGyBfMDxBpwxVZIU/eOPE12', 'jane@mail.com');

--Reviews
INSERT INTO reviews(userID, volumeID, content, rating) VALUES('1000', 'In8mDwAAQBAJ', 'it was pretty good', 6);
INSERT INTO reviews(userID, volumeID, content, rating) VALUES('1000', 'ho-rEAAAQBAJ', 'i didnt like it at all', 3);
INSERT INTO reviews(userID, volumeID, content, rating) VALUES('1000', 'SxQSAAAAYAAJ', '"The Complete Works of Captain John Smith (1580-1631)" is a fascinating dive into early American history. As someone who loves history, I found Smiths firsthand accounts of Jamestown and the Chesapeake Bay incredibly vivid and adventurous. The editors did a great job with annotations and the intro, which really helped make sense of the old-timey language and context. However, some parts felt really outdated, especially Smiths descriptions of Native American societies, which are clearly biased. It’s a mix of thrilling tales and historical insight, but you have to take some of it with a grain of salt. Despite that, its a must-read for history buffs wanting to understand the early days of American colonization and Smiths role in it. Just be prepared for some 17th-century perspectives that might not sit well with modern readers.', 8);
INSERT INTO reviews(userID, volumeID, content, rating) VALUES('1000', 'wrOQLV6xB-wC', 'best book ive ever read', 10);
INSERT INTO reviews(userID, volumeID, content, rating) VALUES('1000', 'GgWgDAEACAAJ', '', 5);
INSERT INTO reviews(userID, volumeID, content) VALUES('1000', 'cyrMu-gkGQQC', 'idk i cant tell if i love it or hate it');
INSERT INTO reviews(userID, volumeID, content, rating) VALUES('1000', '8nt0-BeF2rUC', 'not sure why this young fella was so upset but I wasnt a fan', 3);
INSERT INTO reviews(userID, volumeID, content, rating) VALUES('1000', '4Fs_DwAAQBAJ', 'This sucked ngl didnt like it at all', 2);
INSERT INTO reviews(userID, volumeID, content, rating) VALUES('1000', 'U68lAAAAMAAJ', 'Loved this book so good', 8);
INSERT INTO reviews(userID, volumeID, content, rating) VALUES('1000', '9XzIAwAAQBAJ', 'mid', 5);
INSERT INTO reviews(userID, volumeID, content, rating) VALUES('1000', 'Dba3EAAAQBAJ', 'bro I think im transcended', 5);

--Comments
INSERT INTO comments(userID, reviewid, content) VALUES('1001','1000', 'Yeah I thought it was pretty good as well') RETURNING commentid;
INSERT INTO comments(userID, reviewid, content) VALUES('1002','1000', 'Idk tbh I didnt really enjoy it') RETURNING commentid;
INSERT INTO comments(userID, reviewid, content) VALUES('1003','1000', 'This was genuinely one of the greatest books I have ever read @fred, fuck you.') RETURNING commentid;
INSERT INTO comments(userID, reviewid, content) VALUES('1004','1000', 'I havent read this yet but @ned got me thinking I have to read this for sure.') RETURNING commentid;
INSERT INTO comments(userID, reviewid, content) VALUES('1005','1000', 'Pretty good doesnt do it justice but alright') RETURNING commentid;

--Comment Likes
INSERT INTO user_likes_comment(userID, commentID) VALUES('1001','1000');
UPDATE comments SET likecount = likecount+1 WHERE commentID='1000';
INSERT INTO user_likes_comment(userID, commentID) VALUES('1002','1000');
UPDATE comments SET likecount = likecount+1 WHERE commentID='1000';
INSERT INTO user_likes_comment(userID, commentID) VALUES('1003','1000');
UPDATE comments SET likecount = likecount+1 WHERE commentID='1000';
INSERT INTO user_likes_comment(userID, commentID) VALUES('1004','1001');
UPDATE comments SET likecount = likecount+1 WHERE commentID='1001';
INSERT INTO user_likes_comment(userID, commentID) VALUES('1005','1004');
UPDATE comments SET likecount = likecount+1 WHERE commentID='1004';
INSERT INTO user_likes_comment(userID, commentID) VALUES('1006','1004');
UPDATE comments SET likecount = likecount+1 WHERE commentID='1004';
INSERT INTO user_likes_comment(userID, commentID) VALUES('1007','1004');
UPDATE comments SET likecount = likecount+1 WHERE commentID='1004';

--Review Likes
INSERT INTO user_likes_review(userID, reviewID) VALUES('1001','1000');
UPDATE reviews SET likecount = likecount+1 WHERE reviewID='1000';
INSERT INTO user_likes_review(userID, reviewID) VALUES('1002','1001');
UPDATE reviews SET likecount = likecount+1 WHERE reviewID='1001';
INSERT INTO user_likes_review(userID, reviewID) VALUES('1003','1001');
UPDATE reviews SET likecount = likecount+1 WHERE reviewID='1001';
INSERT INTO user_likes_review(userID, reviewID) VALUES('1007','1001');
UPDATE reviews SET likecount = likecount+1 WHERE reviewID='1001';
INSERT INTO user_likes_review(userID, reviewID) VALUES('1004','1003');
UPDATE reviews SET likecount = likecount+1 WHERE reviewID='1003';
INSERT INTO user_likes_review(userID, reviewID) VALUES('1004','1002');
UPDATE reviews SET likecount = likecount+1 WHERE reviewID='1002';
INSERT INTO user_likes_review(userID, reviewID) VALUES('100A','1003');
UPDATE reviews SET likecount = likecount+1 WHERE reviewID='1003';
INSERT INTO user_likes_review(userID, reviewID) VALUES('100A','1004');
UPDATE reviews SET likecount = likecount+1 WHERE reviewID='1004';
INSERT INTO user_likes_review(userID, reviewID) VALUES('100A','1005');
UPDATE reviews SET likecount = likecount+1 WHERE reviewID='1005';
INSERT INTO user_likes_review(userID, reviewID) VALUES('1007','1006');
UPDATE reviews SET likecount = likecount+1 WHERE reviewID='1006';
INSERT INTO user_likes_review(userID, reviewID) VALUES('1009','1002');
UPDATE reviews SET likecount = likecount+1 WHERE reviewID='1002';

--Follows
INSERT INTO user_follows_user(followerID, followedID) VALUES('1000','1001');
INSERT INTO user_follows_user(followerID, followedID) VALUES('1000','1002');
INSERT INTO user_follows_user(followerID, followedID) VALUES('1000','1003');
UPDATE users SET followercount = followercount+3 WHERE userID='1000';
INSERT INTO user_follows_user(followerID, followedID) VALUES('1001','1000');
UPDATE users SET followercount = followercount+1 WHERE userID='1001';
INSERT INTO user_follows_user(followerID, followedID) VALUES('100A','1000');
UPDATE users SET followercount = followercount+1 WHERE userID='100A';
INSERT INTO user_follows_user(followerID, followedID) VALUES('1005','1000');
UPDATE users SET followercount = followercount+1 WHERE userID='1005';
INSERT INTO user_follows_user(followerID, followedID) VALUES('1008','1000');
UPDATE users SET followercount = followercount+1 WHERE userID='1008';
INSERT INTO user_follows_user(followerID, followedID) VALUES('1007','1000');
UPDATE users SET followercount = followercount+1 WHERE userID='1007';
INSERT INTO user_follows_user(followerID, followedID) VALUES('1002','1000');
UPDATE users SET followercount = followercount+1 WHERE userID='1002';
INSERT INTO user_follows_user(followerID, followedID) VALUES('1004','1000');
UPDATE users SET followercount = followercount+1 WHERE userID='1004';
INSERT INTO user_follows_user(followerID, followedID) VALUES('1003','1000');
UPDATE users SET followercount = followercount+1 WHERE userID='1003';