use twitterdb;

DROP TABLE IF EXISTS tweets;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS followers;

CREATE TABLE users (
  user_id INT NOT NULL AUTO_INCREMENT,
  user VARCHAR(255) NOT NULL,
  passhash VARCHAR(40) NOT NULL,
  email VARCHAR(255) NOT NULL,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  dob DATE,
  PRIMARY KEY (user_id)
);

CREATE TABLE tweets (
  tweet_id INT NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  tweet VARCHAR(140) NOT NULL,
  date_tweet DATETIME NOT NULL,
  PRIMARY KEY (tweet_id),
  FOREIGN KEY user_id(user_id) REFERENCES users(user_id) 
  ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE followers (
  id_user INT NOT NULL REFERENCES users (user_id),
  id_follower INT NOT NULL REFERENCES users (user_id),
  PRIMARY KEY (id_user, id_follower)
);

INSERT INTO users (user, passhash, email, first_name, last_name, dob) VALUES
("foo", "asdsad", "test@gmail.com", "bob", "bobbinson", "1991-01-01"),
("foo2", "asdsad", "test2@gmail.com", "bob2", "bobbinson2", "1992-01-01"),
("foo3", "asdsad", "test3@gmail.com", "bob3", "bobbinson3", "1993-01-01"),
("foo4", "asdsad", "test4@gmail.com", "bob4", "bobbinson4", "1994-01-01"),
("foo5", "asdsad", "test5@gmail.com", "bob5", "bobbinson5", "1995-01-01"),
("foo6", "asdsad", "test5@gmail.com", "bob6", "bobbinson6", "1996-01-01");

INSERT INTO tweets(user_id, tweet, date_tweet) VALUES
(1, "test tweet", "2001-01-01 22:00:00"),
(2, "test tweet2", "2002-01-01 22:00:00"),
(3, "test tweet3", "2003-01-01 22:00:00"),
(4, "test tweet4", "2004-01-01 22:00:00"),
(5, "test tweet5", "2005-01-01 22:00:00");

INSERT INTO followers(id_user, id_follower) VALUES
(5,1),
(4,1),
(3,1),
(2,1),
(6,1),
(2,5),
(4,5);