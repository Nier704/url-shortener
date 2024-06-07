package models

type Url struct {
	ID    string `json:"id"`
	Url   string `json:"url"`
	Views int    `json:"views"`
}
