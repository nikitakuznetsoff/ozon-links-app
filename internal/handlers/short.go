package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/jackc/pgconn"
	"io/ioutil"
	"net/http"

	"github.com/nikitakuznetsoff/ozon-links-app/internal/transfomer"
)


func (handler *LinksHandler) ShortLink(w http.ResponseWriter, r *http.Request) {
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

	err = handler.Repo.Set(req.URL)
	if err != nil {
		pgerr, ok := err.(*pgconn.PgError)
		// Postgres error with code 23505 == db already contains a link
		if ok && pgerr.Code == "23505" {
			fmt.Println("link already in db")
		} else {
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}
	}

	link, err := handler.Repo.GetByLink(req.URL)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	shortLink := transfomer.Encode(link.ID)
	resp, err := json.Marshal(map[string]string{"url": handler.Host + "/" + shortLink})
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}