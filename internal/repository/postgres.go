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

func (db *PostgreDB) GetByLink(link string) (*models.Link, error) {
	shortLink := &models.Link{}
	err := db.conn.
		QueryRow(context.Background(), "select id, link from shortlinks where link = $1", link).
		Scan(&shortLink.ID, &shortLink.Address)
	if err != nil {
		return nil, err
	}
	return shortLink, err
}

func (db *PostgreDB) GetByID(id int) (*models.Link, error) {
	shortLink := &models.Link{}
	err := db.conn.
		QueryRow(context.Background(), "select id, link from shortlinks where id = $1", id).
		Scan(&shortLink.ID, &shortLink.Address)
	if err != nil {
		return nil, err
	}
	return shortLink, err
}

func (db *PostgreDB) Set(link string) (int64, error) {
	result, err := db.conn.
		Exec(context.Background(), "insert into shortlinks (link) values ($1)", link)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected(), nil
}
