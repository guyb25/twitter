# Twitter API

## Routes
> POST /user/register

register a new user.   
example:
```json
{
"username": "guy",
"password": "123"
}
```

> GET /user/request/:username

gets information about a user. in the future should have authentication and be a POST route 
to provide sensitive information like session ID.  
example: GET `localhost:8080/user/request/guy`

> POST /user/update

updates information about a user. Will use the userID to specify which user to change, and will change all
the other fields. can't mutate tweets. 
example:
```json
{  
"id": "31c343b9-a120-4e4c-b23e-6c29ae2d827c",  
"username": "guy2",  
"password": "123",  
"followers": [],  
"following": ["40d5a2a2-e644-44f0-bc0f-a4580dfafd13", "0f99cb9a-3773-4502-8872-8fdb921967d1"]  
}
```

> POST /tweet/create

creates a new tweet.  
example:
```json
{
    "userId": "40d5a2a2-e644-44f0-bc0f-a4580dfafd13",
    "text": "Abu yoyo"
}
```

> POST /tweet/feed/view

views the user feed. will show the tweets created by users that are followed by userId in a sorted way.
example:
```json
{
    "UserId": "31c343b9-a120-4e4c-b23e-6c29ae2d827c"
}
```