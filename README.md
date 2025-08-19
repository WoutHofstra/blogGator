# blogGator


![BlogGator Logo](bloggator.png)


Welcome! This is one of the biggest projects I have made during my studies at
boot.dev. The purpose of this CLI tool i have made is to aggregate RSS feeds and 
view their posts. This tool can be used by multiple people on one device, because
of the login and registering functions i have made.

Usage:

First, install postgre and go if you don't have them already
Then install this CLI with 'go install https://github.com/WoutHofstra/blogGator'
In your home directory, you have to include a file with the following structure:

{  
  "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"  
}

for 'username' use your own name and for 'database' use your database name, which you have made in postgreSQL

Commands:

register: Takes an username as input and registers it as a new user, after this you will have to run login to be logged in  
login: Logs you into your account  
addfeed: takes a name for the new feed and an URL, and adds it to the database  
follow: takes a feed URL and makes the current user follow it  
unfollow: also takes a feed URL, but unfollows it  
reset: resets the whole database  
users: shows all of the registered users  
feeds: shows all of the added feeds  
following: shows all of the feeds the current user is following  
agg: takes a time as input, and turns on the aggregator showing 1 page per cycle. This time should be in a format like "1s", "1m", "1h", etc.   

