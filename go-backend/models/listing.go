package models

import "time"

type Listing struct {
	ID            int64     `json:"id"`
	UserID        int64     `json:"user_id"`
	CategoryID    int       `json:"category_id"`
	Title         string    `json:"title"`
	Description   *string   `json:"description,omitempty"`
	Price         float64   `json:"price"`
	ConditionType string    `json:"condition_type"`
	Location      *string   `json:"location,omitempty"`
	Status        string    `json:"status"`
	Views         int       `json:"views"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ListingWithDetails struct {
	Listing
	CategoryName string   `json:"category_name"`
	CategoryType string   `json:"category_type"`
	SellerName   string   `json:"seller_name"`
	SellerEmail  string   `json:"seller_email"`
	Images       []string `json:"images,omitempty"`
}
