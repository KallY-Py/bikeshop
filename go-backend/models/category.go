package models

import "time"

type Category struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Type      string    `json:"type"` // 'bike' or 'part'
    CreatedAt time.Time `json:"created_at"`
}