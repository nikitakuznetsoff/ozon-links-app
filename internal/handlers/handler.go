package handlers

import (
	"github.com/nikitakuznetsoff/ozon-links-app/internal/repository"
)

type LinksHandler struct {
	Repo	repository.LinksRepository
}