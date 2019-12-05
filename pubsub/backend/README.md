# PubSub backend

## Go commands

```bash
# Install packages
go get ./...

# Compile & run backend and load 50 messages from PostgreSQL
go build && ./backend 50
```

## Env

```
# PostgreSQL
PG_HOST=''
PG_PORT=''
PG_USER=''
PG_PASSWORD=''
PG_DATABASE=''

# Redis
REDIS_HOST=''
REDIS_PASSWORD=''
REDIS_DATABASE=''
```