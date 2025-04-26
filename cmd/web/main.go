package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/aaron-imbrock/snippetbox/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "postgres://web:pass@localhost:5432/snippetbox", "PostgreSQL data source name")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	logger.Info("starting server", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*pgxpool.Pool, error) {
	ctx := context.Background()

	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
