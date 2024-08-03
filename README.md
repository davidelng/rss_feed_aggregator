# RSS Feed Aggregator

## Used

- [Go](https://go.dev/)
- [SQLC](https://sqlc.dev/)
- [Goose](http://pressly.github.io/goose/)
- [PostgreSQL](https://www.postgresql.org/golan)

## Description

A restful API platform that scrapes RSS feeds and collects their posts periodically. It also handles user creation and authentication.

## Endpoints

All the authenticated endpoint needs a `Authentication: ApiKey XXX` in the headers, the key is given at user creation.

### Health status

- `GET /v1/healthz` report service status

### Users

- `POST /v1/users` creates a new user
- `GET /v1/users` (auth required) get current user
- `GET /v1/posts` (auth required) get posts from user saved feeds

### Feeds

- `GET /v1/feeds` get all feeds saved in db
- `POST /v1/feeds` (auth required) save a new feed to scrape

### Feed follows

- `GET /v1/feed_follows` (auth required) get followed feeds
- `POST /v1/feed_follows` (auth required) follow a feed
- `DELETE /v1/feed_follows` (auth required) remove follow to a feed

# Ideas for extending the project

- Support pagination of the endpoints that can return many items
- Support different options for sorting and filtering posts using query parameters
- Classify different types of feeds and posts (e.g. blog, podcast, video, etc.)
- Add a CLI client that uses the API to fetch and display posts, maybe it even allows you to read them in your terminal
- Scrape lists of feeds themselves from a third-party site that aggregates feed URLs
- Add support for other types of feeds (e.g. Atom, JSON, etc.)
- Add integration tests that use the API to create, read, update, and delete feeds and posts
- Add bookmarking or "liking" to posts
- Create a simple web UI that uses your backend API
