package entity

import (
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

type Product struct {
	ID             uuid.UUID  `json:"id"`
	Name           string     `json:"name" binding:"required"`
	Description    string     `json:"description" binding:"required"`
	Price          float64    `json:"price" binding:"required"`
	Stock          int        `json:"stock" binding:"required" `
	Category       string     `json:"category" binding:"required"`
	SKU            string     `json:"sku"`
	ImageURLs      []string   `json:"image_urls" binding:"required"` // Matches the correct JSON field name
	Discount       float64    `json:"discount"`
	IsActive       bool       `json:"is_active"`
	Tags           []string   `json:"tags" binding:"required"`
	AdditionalInfo string     `json:"additional_info"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}

// convert image urls to string
func (p *Product) ImageURLsToString() string {
	return strings.Join(p.ImageURLs, ",")
}

// parse string to image urls
func (p *Product) StringToImageURLs(input string) {
	p.ImageURLs = strings.Split(input, ",")
}

// convert tags to string
func (p *Product) TagsToString() string {
	return strings.Join(p.Tags, ",")
}

// parse string to tags
func (p *Product) StringToTags(input string) {
	p.Tags = strings.Split(input, ",")
}
