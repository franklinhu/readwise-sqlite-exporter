# readwise-sqlite-exporter

Exports Readwise books and highlights to a SQLite3 database

## Usage

```
READWISE_KEY=1234 go run github.com/franklinhu/readwise-sqlite-exporter/cli
```

The Readwise client assumes the API key is passed as the `READWISE_KEY` environment variable.

## Development

### sqlc

This project uses [sqlc](https://github.com/kyleconroy/sqlc) to generate the wrappers around `database/sql`.
You'll need to install sqlc if you intend to modify the schema or available queries.

```
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
sqlc generate
```

## TODO

- [ ] Readwise pagination (client only fetches a single page)

