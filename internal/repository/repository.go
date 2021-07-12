package repository

import (
	"github.com/nikitakuznetsoff/ozon-links-app/internal/models"
)

type LinksRepository interface {
	GetByLink(link string) (*models.Link, error)
	GetByID(id int) (*models.Link, error)
	Set(link string) (int64, error)
}