## Blog Aggregator gator (CLI)

Command-line RSS/Atom blog aggregator written in Go with PostgreSQL and sqlc.

### Prerequisites

- Go (matching version in `go.mod`)
- PostgreSQL 13+
- sqlc (optional, for regenerating DB code)

### Setup

1. Create a PostgreSQL database and note its connection URL (example):
   - `postgres://USER:PASSWORD@localhost:5432/blog_aggregator?sslmode=disable`
2. In the project root, create `.gatorconfig.json`:

```json
{
  "db_url": "postgres://USER:PASSWORD@localhost:5432/blog_aggregator?sslmode=disable",
  "current_user_name": ""
}
```

3. Apply schema migrations in order:
   - `sql/schema/001_users.sql`
   - `sql/schema/002_feeds.sql`
   - `sql/schema/003_ feed_follows.sql`
   - `sql/schema/004_posts.sql`

You can run them with psql, e.g.:

```bash
psql "$DB_URL" -f sql/schema/001_users.sql
psql "$DB_URL" -f sql/schema/002_feeds.sql
psql "$DB_URL" -f "sql/schema/003_ feed_follows.sql"
psql "$DB_URL" -f sql/schema/004_posts.sql
```

### Run

From the project root:

```bash
go run . <command> [args]
```

### Commands

- `register <username>`: Create a new user and set as current.
- `login <username>`: Set current user (must already exist).
- `users`: List users; marks the current user.
- `reset`: Delete all rows from users, feeds, and feed_follows.
- `feeds`: List all feeds with owners.
- `addfeed <name> <url>`: Create a feed and auto-follow it. Requires login.
- `follow <url>`: Follow an existing feed by URL. Requires login.
- `following`: List feeds followed by the current user. Requires login.
- `unfollow <url>`: Unfollow a feed by URL. Requires login.
- `browse`: Print followed feeds for the current user. Requires login.
- `agg`: Fetch and print a sample RSS feed (dev utility).

Examples:

```bash
go run . register alice
go run . login alice
go run . addfeed "WagsLane" https://www.wagslane.dev/index.xml
go run . feeds
go run . follow https://www.wagslane.dev/index.xml
go run . following
go run . browse
go run . unfollow https://www.wagslane.dev/index.xml
```

### Configuration Notes

- The app reads `.gatorconfig.json` from the current working directory at runtime.
- On certain actions, the config may also be written to a user path on this machine; ensure your primary source of truth is the project-root `.gatorconfig.json` when running via `go run .`.

### Development

- sqlc config: `sqlc.yaml` (inputs in `sql/schema` and `sql/queries`, output to `internal/database`).
- Regenerate DB code after changing SQL:

```bash
sqlc generate
```

### Troubleshooting

- If you see “not logged in”, run `login <username>` first.
- If a user already exists on register, use `login` instead.
- Ensure `db_url` in `.gatorconfig.json` points to a reachable Postgres instance.
