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

func main() {
	dbURL := "postgres://nick:pass@db:5432/linksdb"
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	repo := repository.CreateRepo(conn)
	handler := handlers.LinksHandler{Repo: repo, Host: "localhost:6000"}

	http.HandleFunc("/short", handler.ShortLink)
	http.HandleFunc("/long", handler.LongLink)
	fmt.Println("Starting server at :6000")
	log.Fatal(http.ListenAndServe(":6000", nil))
}