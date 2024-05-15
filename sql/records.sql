--Users
INSERT INTO users(username, password, email) VALUES('ted', '$2a$10$NxMO/aiddukYPXCM7/riCeKETTOUYEcGyBfMDxBpwxVZIU/eOPE12', 'ted@mail.com');
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
INSERT INTO reviews(username, volumeID, content, rating) VALUES('ted', 'In8mDwAAQBAJ', 'it was pretty good', 6) RETURNING reviewID;
INSERT INTO reviews(username, volumeID, content, rating) VALUES('ted', 'ho-rEAAAQBAJ', 'i didnt like it at all', 3) RETURNING reviewID;
INSERT INTO reviews(username, volumeID, content, rating) VALUES('ted', 'wrOQLV6xB-wC', 'best book ive ever read', 10) RETURNING reviewID;
INSERT INTO reviews(username, volumeID, content, rating) VALUES('ted', 'GgWgDAEACAAJ', '', 5) RETURNING reviewID;
INSERT INTO reviews(username, volumeID, content) VALUES('ted', 'OL15110516W', 'idk i cant tell if i love it or hate it') RETURNING reviewID;
INSERT INTO reviews(username, volumeID, content, rating) VALUES('ted', '8nt0-BeF2rUC', 'not sure why this young fella was so upset but I wasnt a fan', 3) RETURNING reviewID;
INSERT INTO reviews(username, volumeID, content, rating) VALUES('ted', '4Fs_DwAAQBAJ', 'This sucked ngl didnt like it at all', 2) RETURNING reviewID;
INSERT INTO reviews(username, volumeID, content, rating) VALUES('ted', 'U68lAAAAMAAJ', 'Loved this book so good', 8) RETURNING reviewID;

--Comments
INSERT INTO comments(username, reviewid, content) VALUES('john','1', 'Yeah I thought it was pretty good as well') RETURNING commentid;
INSERT INTO comments(username, reviewid, content) VALUES('fred','1', 'Idk tbh I didnt really enjoy it') RETURNING commentid;
INSERT INTO comments(username, reviewid, content) VALUES('ned','1', 'This was genuinely one of the greatest books I have ever read @fred, fuck you.') RETURNING commentid;
INSERT INTO comments(username, reviewid, content) VALUES('allen','1', 'I havent read this yet but @ned got me thinking I have to read this for sure.') RETURNING commentid;
INSERT INTO comments(username, reviewid, content) VALUES('adam','1', 'Pretty good doesnt do it justice but alright') RETURNING commentid;
