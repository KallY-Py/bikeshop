package models

import "time"

type Report struct {
	ID             int64     `json:"id"`
	ReporterID     int64     `json:"reporter_id"`
	ListingID      *int64    `json:"listing_id,omitempty"`
	ReportedUserID *int64    `json:"reported_user_id,omitempty"`
	Reason         string    `json:"reason"`
	Status         string    `json:"status"` // pending, investigating, resolved
	CreatedAt      time.Time `json:"created_at"`

	// Joined fields for admin view
	ReporterName  string `json:"reporter_name,omitempty"`
	ReporterEmail string `json:"reporter_email,omitempty"`
	TargetName    string `json:"target_name,omitempty"`
}
