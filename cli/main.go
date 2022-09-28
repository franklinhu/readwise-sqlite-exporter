package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	exporter "github.com/franklinhu/readwise-sqlite-exporter"
)

var sqlitePath string

func init() {
	now := time.Now()
	flag.StringVar(
		&sqlitePath,
		"sqlite_path",
		fmt.Sprintf("readwise-%s-%d.db", now.Format("2006-01-02"), now.Unix()),
		"Output path of the SQLite3 database file. Defaults to readwise-yyyy-mm-dd-unixts",
	)
}

func main() {
	flag.Parse()

	ctx := context.Background()

	db, err := sql.Open("sqlite3", sqlitePath)
	if err != nil {
		log.Fatalf("Error opening sqlite db: %v\n", err)
	}

	exporter.SetupDDL(ctx, db)
	exporter.ExportReadwiseToSqlite(ctx, db)
}
