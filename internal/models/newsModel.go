package models

import "github.com/gofrs/uuid"

type News struct {
	ID           uuid.UUID `json:"id"`
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []struct {
		ID     string `json:"id"`
		Source struct {
			Name string `json:"name"`
		} `json:"source"`
		Author      string `json:"author"`
		Title       string `json:"title"`
		Description string `json:"description"`
		URL         string `json:"url"`
		PublishedAt string `json:"publishedAt"`
	} `json:"articles"`
}
