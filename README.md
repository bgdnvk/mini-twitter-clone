# Mini Twitter Clone  
  
  
## Disclaimer
This is an introductory project to look at Golang a bit.  
This project is not production ready and it has bad practices like the way pagination works or how we interact with the database. It was mainly made for the blog post that you will find here:  
https://bognov.tech/introduction-to-golang-build-a-mini-twitter-clone#heading-bonus-ukrainian-meme

Through the codebase you will find different "prints" that are used to debug the project, feel free to play with them.

## What's missing for production ready?
You should add a proper logger, configuration, middleware, a different way to handle the data, perhaps an ORM, and a better way to handle the pagination, as well as a better API.  

Note that you can make yourself vulnerable to SQL ingections if you copy and paste the code. This has been made for learning purposes.