# go-project-rrs-aggregator-backend

## Packages

- github.com/joho/godotenv
- github.com/go-chi/chi
- github.com/go-chi/cors

## Tools

- Thunder client
- github.com/sqlc-dev/sqlc/cmd/sqlc@latest

```bash
sqlc version
# print: v1.25.0
```

- github.com/pressly/goose/v3/cmd/goose@latest

```bash
goose -version
# print: goose version: v3.18.0
```

## Commands

```bash
goose postgres postgres://postgres:postgres@localhost:5432/go-project-local up
goose postgres postgres://postgres:postgres@localhost:5432/go-project-local down
```
