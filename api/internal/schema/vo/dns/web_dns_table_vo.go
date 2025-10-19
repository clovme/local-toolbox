package dns

import (
	"time"
	"toolbox/pkg/enums/status"
)

type WebDNSTableVO struct {
	ID        string        `json:"id"`
	Protocol  string        `json:"protocol"`
	Domain    string        `json:"domain"`
	IP        string        `json:"ip"`
	Status    status.Status `json:"status"`
	Port      string        `json:"port"`
	CreatedAt *time.Time    `json:"createdAt"`
	UpdatedAt *time.Time    `json:"updatedAt"`
}

type WebSaveDataVO struct {
	InsertCount int   `json:"insertCount"`
	UpdateCount int   `json:"updateCount"`
	DeleteCount int64 `json:"deleteCount"`
}
