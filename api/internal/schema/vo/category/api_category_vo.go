package category

import (
	"time"
	"toolbox/pkg/enums/status"
)

type ApiCategoryVO struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Title        string        `json:"title"`
	DocSort      string        `json:"docSort"`
	ArticleCount int64         `json:"articleCount"`
	Pid          string        `json:"pid"`
	Sort         int           `json:"sort"`
	Status       status.Status `json:"status"`
	CreatedAt    *time.Time    `json:"createdAt"`
	UpdatedAt    *time.Time    `json:"updatedAt"`
}

type CountResultVO struct {
	CategoryID   int64
	ArticleCount int64
}
