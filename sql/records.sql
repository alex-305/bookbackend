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
INSERT INTO reviews(username, volumeID, content, rating) VALUES('ted', 'In8mDwAAQBAJ', 'it was pretty good', 6);
INSERT INTO reviews(username, volumeID, content, rating) VALUES('ted', 'ho-rEAAAQBAJ', 'i didnt like it at all', 3);
INSERT INTO reviews(username, volumeID, content, rating) VALUES('ted', 'SxQSAAAAYAAJ', '"The Complete Works of Captain John Smith (1580-1631)" is a fascinating dive into early American history. As someone who loves history, I found Smiths firsthand accounts of Jamestown and the Chesapeake Bay incredibly vivid and adventurous. The editors did a great job with annotations and the intro, which really helped make sense of the old-timey language and context. However, some parts felt really outdated, especially Smiths descriptions of Native American societies, which are clearly biased. It’s a mix of thrilling tales and historical insight, but you have to take some of it with a grain of salt. Despite that, its a must-read for history buffs wanting to understand the early days of American colonization and Smiths role in it. Just be prepared for some 17th-century perspectives that might not sit well with modern readers.', 8);
INSERT INTO reviews(username, volumeID, content, rating) VALUES('ted', 'wrOQLV6xB-wC', 'best book ive ever read', 10);
INSERT INTO reviews(username, volumeID, content, rating) VALUES('ted', 'GgWgDAEACAAJ', '', 5);
INSERT INTO reviews(username, volumeID, content) VALUES('ted', 'cyrMu-gkGQQC', 'idk i cant tell if i love it or hate it');
INSERT INTO reviews(username, volumeID, content, rating) VALUES('ted', '8nt0-BeF2rUC', 'not sure why this young fella was so upset but I wasnt a fan', 3);
INSERT INTO reviews(username, volumeID, content, rating) VALUES('ted', '4Fs_DwAAQBAJ', 'This sucked ngl didnt like it at all', 2);
INSERT INTO reviews(username, volumeID, content, rating) VALUES('ted', 'U68lAAAAMAAJ', 'Loved this book so good', 8);
INSERT INTO reviews(username, volumeID, content, rating) VALUES('ted', '9XzIAwAAQBAJ', 'mid', 5);
INSERT INTO reviews(username, volumeID, content, rating) VALUES('ted', 'Dba3EAAAQBAJ', 'bro I think im transcended', 5);

--Comments
INSERT INTO comments(username, reviewid, content) VALUES('john','1', 'Yeah I thought it was pretty good as well') RETURNING commentid;
INSERT INTO comments(username, reviewid, content) VALUES('fred','1', 'Idk tbh I didnt really enjoy it') RETURNING commentid;
INSERT INTO comments(username, reviewid, content) VALUES('ned','1', 'This was genuinely one of the greatest books I have ever read @fred, fuck you.') RETURNING commentid;
INSERT INTO comments(username, reviewid, content) VALUES('allen','1', 'I havent read this yet but @ned got me thinking I have to read this for sure.') RETURNING commentid;
INSERT INTO comments(username, reviewid, content) VALUES('adam','1', 'Pretty good doesnt do it justice but alright') RETURNING commentid;
