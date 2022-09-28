CREATE TABLE IF NOT EXISTS books (
	id INTEGER PRIMARY KEY,
	title TEXT NOT NULL,
	author TEXT NOT NULL,
	category TEXT NOT NULL,
	source TEXT NOT NULL,
	num_highlights INTEGER NOT NULL,
	last_highlight_at DATE NOT NULL,
	updated DATE NOT NULL,
	cover_image_url TEXT NOT NULL,
	highlights_url TEXT NOT NULL,
	source_url TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS highlights (
  id INTEGER PRIMARY KEY,
  text TEXT NOT NULL,
  note TEXT NOT NULL,
  location INTEGER NOT NULL,
  location_type TEXT NOT NULL,
  highlighted_at DATE NOT NULL,
  book_id INTEGER NOT NULL,
  url TEXT NOT NULL,
  color TEXT NOT NULL,
  updated DATE NOT NULL,
  tags TEXT NOT NULL,
  FOREIGN KEY(book_id) REFERENCES books(id)
);
