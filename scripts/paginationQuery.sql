USE twitterdb;

--example query
SELECT users.user_id,
    users.user,
    users.first_name,
    users.last_name,
    tweets.tweet,
    tweets.date_tweet
FROM users
    INNER JOIN tweets ON users.user_id = tweets.user_id
    INNER JOIN followers ON users.user_id = followers.id_user
WHERE followers.id_follower = 1
ORDER BY tweets.date_tweet DESC
LIMIT 2 OFFSET 1;