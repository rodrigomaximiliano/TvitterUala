# Technical Challenge for Ualá

> **TvitterUala – a microblogging platform similar to Twitter**

## Architecture and Components

- **Language:** Go  
- **Framework:** Fiber  
- **Structure:** Layered separation (`main.go`, `models`, `handlers`, `storage`)  
- **Storage:** In-memory (maps and slices)  
- **REST Endpoints:** For users, tweets, follows, and timeline  
- **Documentation:** This README describes the architecture, components, and usage  

## Description

TvitterUala is a sample application that simulates a simple version of Twitter. It allows users to:  
- Create a user account  
- Post short messages called "tweets"  
- Follow other users  
- View a timeline with their own tweets and tweets from users they follow  

The backend is developed in Go using the Fiber framework. All data is stored in memory (no disk persistence).  

---

## Project Structure

backend/
│
├── main.go # Initializes the server and defines main routes
├── models/
│ └── models.go # Data models: User, Tweet, Follow
├── handlers/
│ └── tweet_handlers.go # Endpoint logic (users, tweets, follow, timeline)
├── storage/
│ └── memory.go # In-memory storage (users, follows, indexed tweets)


---

## How does it work?

### 1. Create User

Allows registering a new user.

- **Method:** POST  
- **Route:** `/users`  
- **Example request body:**
  ```json
  {
    "id": "user1",
    "name": "Name"
  }

    Example response:

    {
      "id": "user1",
      "name": "Name"
    }

2. Post Tweet

Allows a user to post a short message (max 280 characters).

    Method: POST

    Route: /tweets

    Example request body:

{
  "user_id": "user1",
  "text": "Hello world"
}

Example response:

    {
      "id": "generated-uuid",
      "user_id": "user1",
      "text": "Hello world",
      "timestamp": "2024-06-01T12:34:56.789Z"
    }

3. Follow User

Allows a user to follow another user.

    Method: POST

    Route: /follow

    Example request body:

{
  "follower_id": "user1",
  "followee_id": "user2"
}

Example response:

    {
      "message": "followed"
    }

4. View Timeline

Returns tweets from the user and the users they follow, ordered from newest to oldest.
Supports pagination to avoid returning too many tweets at once.

    Method: GET

    Route: /timeline?user_id=user1&page=1&size=10

        user_id: ID of the user requesting their timeline (required)

        page: page number (optional, default 1)

        size: number of tweets per page (optional, default 10)

    Example response:

    {
      "page": 1,
      "size": 10,
      "total": 3,
      "timeline": [
        {
          "id": "uuid1",
          "user_id": "user2",
          "text": "Tweet from user2",
          "timestamp": "2024-06-01T12:34:56.789Z"
        },
        {
          "id": "uuid2",
          "user_id": "user1",
          "text": "My own tweet",
          "timestamp": "2024-06-01T12:30:00.000Z"
        }
      ]
    }

        The timeline includes tweets from the user and those they follow, ordered newest to oldest.

        You can change the page and size using the page and size parameters.

How is data stored?

    Users:
    Stored in memory in a map (key: user ID).

    Tweets:
    Stored in memory, indexed by user for quick lookup.

    Follows:
    Stored as a list of who follows whom.

    Note: Data is lost if the server stops because no database is used.

How to run the backend?

    Install dependencies:

go mod tidy

Run the backend:

    go run main.go

    You can test the endpoints using Postman, curl, or any HTTP client.

Full example flow

    Create two users (user1 and user2).

    user1 follows user2.

    Both post tweets.

    Querying user1’s timeline will show their own tweets and those of user2.

---
