package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Nier704/url-shortener/internal/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type UrlHandler struct {
	DB *sql.DB
}

type UrlRequest struct {
	Url string `json:"url"`
}

func (h *UrlHandler) GetUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("r")

	if id == "" {
		log.Println("Invalid Id")
		http.Error(w, "Invalid Id", 400)
		return
	}

	sql := `
	UPDATE urls SET views = views + 1
	WHERE id = $1
	RETURNING url
	`

	var url string
	err := h.DB.QueryRowContext(r.Context(), sql, id).Scan(&url)
	if err != nil {
		log.Printf("Error executing context: %v", err)
		http.Error(w, "URL not found", 404)
		return
	}
	println(url)

	http.Redirect(w, r, url, http.StatusFound)
}

func (h *UrlHandler) CreateUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body UrlRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("Error decoding body: %v", err)
		http.Error(w, "Invalid Body", 500)
		return
	}

	newUrl := models.Url{
		ID:    uuid.New().String(),
		Url:   body.Url,
		Views: 0,
	}

	sql := `
	INSERT INTO urls (id, url, views)
	VALUES ($1, $2, $3)
	ON CONFLICT(url) DO NOTHING
	RETURNING id
	`

	var insertedId string
	err := h.DB.QueryRowContext(r.Context(), sql, newUrl.ID, newUrl.Url, newUrl.Views).
		Scan(&insertedId)

	if err != nil {
		log.Printf("Error executing context: %v", err)
		http.Error(w, "Error executing context", 500)
		return
	}

	res := map[string]string{
		"message": "created",
		"id":      insertedId,
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Error encoding response", 500)
	}
}

func NewUrlHandler(db *sql.DB) *UrlHandler {
	return &UrlHandler{
		DB: db,
	}
}
