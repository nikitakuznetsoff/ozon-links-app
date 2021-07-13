package repository

import (
	"github.com/nikitakuznetsoff/ozon-links-app/internal/models"
)

type LinksRepository interface {
	GetByLink(url string) (*models.Link, error)
	GetByID(id int) (*models.Link, error)
	Set(url string) error
}