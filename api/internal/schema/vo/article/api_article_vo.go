package article

import (
	"time"
)

type ApiArticleVO struct {
	ID            string     `json:"id"`
	Title         string     `json:"title"`
	CategoryID    string     `json:"categoryID"`
	CategoryTitle string     `json:"categoryTitle"`
	CategoryName  string     `json:"categoryName"`
	Tags          string     `json:"tags"`
	Summary       string     `json:"summary"`
	Content       *string    `json:"content,omitempty"`
	CreatedAt     *time.Time `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
}
