package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"
	// _ indicates a blank identifier
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger *slog.Logger
}

func main() {
	// Other flag values: flag.Int(), flag.Bool(), flag.Float64() and flag.Duration()
	addr := flag.String("addr", ":4000", "HTTP network address")
	// dsn = username:password@protocol(address)?param=value
	dsn := flag.String("dsn", "web:pass1234@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// defer allows db.Close() to run before the main funciton exits
	// However, a graceful shutdown will be useful
	defer db.Close()

	app := &application {
		logger: logger,
	}

	logger.Info("starting server", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
