-- name: GetBook :one
SELECT * FROM books
WHERE id = ? LIMIT 1;

-- name: GetMostRecentlyUpdatedBook :one
SELECT * FROM books
ORDER BY updated DESC
LIMIT 1;

-- name: ListBooks :many
SELECT * FROM books;

-- name: CreateBook :one
INSERT INTO books (
	id, title, author, category, source, num_highlights, last_highlight_at,
	updated, cover_image_url, highlights_url, source_url
) VALUES (
	?, ?, ?, ?, ?, ?, ?,
	?, ?, ?, ?
)
RETURNING *;

-- name: CreateHighlight :one
INSERT INTO highlights (
  id, text, note, location, location_type, highlighted_at, book_id, url,
  color, updated, tags
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?,
  ?, ?, ?
)
RETURNING *;
