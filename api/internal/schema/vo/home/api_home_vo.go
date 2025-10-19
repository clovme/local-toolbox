package home

import (
	"time"
	"toolbox/pkg/enums/icon"
	"toolbox/pkg/enums/position"
)

type ApiHomeVO struct {
	ID        string            `json:"id"`
	Icon      string            `json:"icon"`
	Name      string            `json:"name"`
	Sort      int               `json:"sort"`
	IconType  icon.Icon         `json:"iconType"`
	Position  position.Position `json:"position"`
	CreatedAt *time.Time        `json:"createdAt"`
	UpdatedAt *time.Time        `json:"updatedAt"`
}
