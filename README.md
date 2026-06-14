# gator

Lightweight CLI RSS/blog aggregator written in Go with PostgreSQL.

## Run

```bash
go run . <command> [args...]
```

## Available Commands

- `register` - create a user
- `login` - switch active user
- `reset` - clear user data
- `users` - list users
- `addfeed` - add a feed URL
- `feeds` - list feeds
- `follow` - follow a feed
- `following` - list followed feeds
- `unfollow` - unfollow a feed
- `agg` - run feed aggregation/fetching
