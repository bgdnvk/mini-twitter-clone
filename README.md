# Mini Twitter Clone  
  
  
1. Write a database schema for creating the tables in a MariaDB/MySQL database with users, tweets, and followers.  

1.1 A user has a username, password, email, first name, last name, and age.  

1.2 A tweet has text, publication date, and is posted by a user.  

1.3 A user can be followed by other users.  

The tables may contain additional structures that provide better space efficiency, security, or performance.

2. Write REST API for the tweet. It should be able to return the tweets by users followed by a specific user. The result must include username, first name, last name, tweet text, and publication date and support pagination.

3. Write a utility function for a User object. You will be implementing three different functions, each with the following function signature: func ([]User) bool.  

3.1  AtLeastTwice is a function that returns true if there is a person who is at least twice as old as any other person in the list, otherwise, the function returns false.  

3.2  ExactlyTwice is a function that returns true if there is a person who is exactly twice as old as any other person in the list, otherwise the function returns false.  

3.3  ConstrainedExactlyTwice is a function that behaves like ExactlyTwice, but input age values are guaranteed to always be within the range 18 to 80, and this function must perform extremely well (can be over-optimized for performance, considering time and space complexity).  