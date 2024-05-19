package kankalib

import (
	"time"
)

type Attributes struct {
	Data []Data    `json:"data,omitempty"`
	Sync time.Time `json:"sync,omitempty"`
}
type Data struct {
	APIKey       string    `json:"api_key,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	CreatedBy    any       `json:"created_by,omitempty"`
	DefaultOrder int       `json:"default_order,omitempty"`
	EntityID     int       `json:"entity_id,omitempty"`
	ID           int       `json:"id,omitempty"`
	IsPinned     bool      `json:"is_pinned,omitempty"`
	IsPrivate    bool      `json:"is_private,omitempty"`
	IsStar       bool      `json:"is_star,omitempty"`
	Name         string    `json:"name,omitempty"`
	Parsed       string    `json:"parsed,omitempty"`
	Type         string    `json:"type,omitempty"`
	TypeID       int       `json:"type_id,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	UpdatedBy    any       `json:"updated_by,omitempty"`
	Value        any       `json:"value,omitempty"`
}