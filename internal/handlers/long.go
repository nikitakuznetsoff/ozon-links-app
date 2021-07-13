package handlers

import (
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/nikitakuznetsoff/ozon-links-app/internal/transfomer"
)

func (handler *LinksHandler) LongLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "incorrect method", http.StatusBadRequest)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "request body reading error", http.StatusInternalServerError)
		return
	}

	req := &struct{URL string `json:"url"`}{}
	err = json.Unmarshal(body, req)
	if err != nil || req.URL == "" {
		http.Error(w, "incorrect request body", http.StatusBadRequest)
		return
	}
	// Checking that the link contains the current host
	if strings.Index(req.URL, handler.Host) != 0 {
		http.Error(w, "incorrect link", http.StatusBadRequest)
		return
	}
	// URL without host
	url := req.URL[len(handler.Host)+1:]
	linkID, err := transfomer.Decode(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	link, err := handler.Repo.GetByID(linkID)
	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "unknown url", http.StatusNotFound)
		} else {
			http.Error(w, "database error", http.StatusInternalServerError)
		}
		return
	}

	resp, err := json.Marshal(map[string]string{"url": link.Address})
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}