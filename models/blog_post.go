package models

import (
	"time"

	"github.com/lib/pq"
)

// BlogPost represents a post in the blogging platform.
type BlogPost struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	Category  string         `json:"category"`
	Tags      pq.StringArray `gorm:"type:text[]" json:"tags"` // Uses PostgreSQL text array
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}
