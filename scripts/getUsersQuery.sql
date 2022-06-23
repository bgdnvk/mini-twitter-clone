USE twitterdb;

SELECT user, dob, TIMESTAMPDIFF(YEAR, dob, CURDATE()) AS age 
FROM users
ORDER BY age DESC;