package repository

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/nikitakuznetsoff/ozon-links-app/internal/models"
)

type PostgreDB struct {
	conn *pgx.Conn
}

func CreateRepo(db *pgx.Conn) *PostgreDB {
	return &PostgreDB{conn: db}
}

func (db *PostgreDB) GetByLink(url string) (*models.Link, error) {
	link := &models.Link{}
	err := db.conn.
		QueryRow(context.Background(), "select id, url from links where url = $1", url).
		Scan(&link.ID, &link.Address)
	if err != nil {
		return nil, err
	}
	return link, err
}

func (db *PostgreDB) GetByID(id int) (*models.Link, error) {
	link := &models.Link{}
	err := db.conn.
		QueryRow(context.Background(), "select id, url from links where id = $1", id).
		Scan(&link.ID, &link.Address)
	if err != nil {
		return nil, err
	}
	return link, err
}

func (db *PostgreDB) Set(link string) error {
	_, err := db.conn.
		Exec(context.Background(), "insert into links (url) values ($1)", link)
	if err != nil {
		return err
	}
	return nil
}
