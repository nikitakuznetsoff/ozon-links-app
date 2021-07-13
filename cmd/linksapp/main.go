package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"os"

	"github.com/nikitakuznetsoff/ozon-links-app/internal/handlers"
	"github.com/nikitakuznetsoff/ozon-links-app/internal/repository"
)

const (
	host = "localhost"
	port = ":6000"
	dbURI = "postgres://nick:pass@db:5432/linksdb"
)

func main() {
	dsn := dbURI
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	repo := repository.CreateRepo(conn)
	handler := handlers.LinksHandler{Repo: repo, Host: host + port}

	http.HandleFunc("/short", handler.ShortLink)
	http.HandleFunc("/long", handler.LongLink)
	fmt.Println("Starting server at", port)
	log.Fatal(http.ListenAndServe(port, nil))
}