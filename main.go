package readwise_sqlite_exporter

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ethanholz/readwise-go"

	sqlc "github.com/franklinhu/readwise-sqlite-exporter/gen_readwise_sqlc"

	_ "embed"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

func parseTime(s string) time.Time {
	t, _ := time.Parse(time.RFC3339, s)
	return t
}

func readwiseBookToCreateBookParams(book readwise.Book) sqlc.CreateBookParams {
	return sqlc.CreateBookParams{
		ID:              int64(book.ID),
		Title:           book.Title,
		Author:          book.Author,
		Category:        book.Category,
		Source:          book.Source,
		NumHighlights:   int64(book.NumHighlights),
		LastHighlightAt: parseTime(book.LastHighlightAt),
		Updated:         parseTime(book.Updated),
		CoverImageUrl:   book.CoverImageURL,
		HighlightsUrl:   book.HighlightsURL,
		SourceUrl:       book.SourceURL,
	}
}

func readwiseHighlightToCreateHighlightParams(
	highlight readwise.Highlight,
) (sqlc.CreateHighlightParams, error) {
	tags, err := json.Marshal(highlight.Tags)
	if err != nil {
		return sqlc.CreateHighlightParams{}, err
	}
	return sqlc.CreateHighlightParams{
		ID:            int64(highlight.ID),
		Text:          highlight.Text,
		Note:          highlight.Note,
		Location:      int64(highlight.Location),
		LocationType:  highlight.LocationType,
		HighlightedAt: parseTime(highlight.HighlightedAt),
		BookID:        int64(highlight.BookID),
		Url:           highlight.URL,
		Color:         highlight.Color,
		Updated:       parseTime(highlight.Updated),
		Tags:          string(tags),
	}, nil
}

func writeBooksToSqlite(
	ctx context.Context,
	books []readwise.Book,
	queries *sqlc.Queries,
) error {
	for _, book := range books {
		_, err := queries.CreateBook(ctx, readwiseBookToCreateBookParams(book))
		if err != nil {
			return err
		}
	}
	return nil
}

func writeHighlightsToSqlite(
	ctx context.Context,
	highlights []readwise.Highlight,
	queries *sqlc.Queries,
) error {
	for _, highlight := range highlights {
		params, err := readwiseHighlightToCreateHighlightParams(highlight)
		if err != nil {
			return err
		}

		_, err2 := queries.CreateHighlight(ctx, params)
		if err2 != nil {
			return err2
		}
	}
	return nil
}

func SetupDDL(ctx context.Context, db sqlc.DBTX) error {
	_, err := db.ExecContext(ctx, ddl)
	return err
}

func ExportReadwiseToSqlite(ctx context.Context, db sqlc.DBTX) error {
	queries := sqlc.New(db)

	readwiseClient := readwise.New()
	bookList, err := readwiseClient.GetBookList()
	if err != nil {
		return *err
	}

	writeBooksToSqlite(ctx, bookList.Results, queries)

	highlights, err := readwiseClient.GetHighlightList()
	if err != nil {
		return *err
	}
	writeHighlightsToSqlite(ctx, highlights.Results, queries)
	return nil
}
