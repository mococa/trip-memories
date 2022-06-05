## Welcome to TripMemories API

### About

This is a simple Go webserver that serves the Flutter App a back-end.

It stores `trips` and `users` to a DynamoDB, using AWS-SDK.

I tried to implement principles of SOLID architecture that I've been learning for the last month

The Dynamo follows the single table design, which means that we have only one table for everything.

### Getting Started:

```bash
make dev
```

### Routes:

#### Trips

`GET /trips` - Brings information about a certain trip

- Needs 2 query params: `user_pk` and `trip_sk`

`GET /trips/list` - Brings all trips from a certain user

- Needs 1 query param: `user_pk`

`POST /trips` - Stores a trip into DynamoDB

- Needs one JSON body:

```json
{
  "name": "string",
  "description": "string",
  "done_at": "string"
}
```

#### Authentication

`GET /auth` - Brings information about a certain user

- Needs 1 query param: `user_pk`

`POST /auth` - Stores a user into DynamoDB

- Needs one JSON body:

```json
{
  "google_id": "000000000000000",
  "name": "Luiz Felipe",
  "email": "luizfelipesmoureau@gmail.com",
  "picture": "https://lh3.googleusercontent.com/..."
}
```
