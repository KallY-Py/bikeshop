package models

import "time"

type User struct {
    ID           int64     `json:"id"`
    FirstName    string    `json:"first_name"`
    LastName     string    `json:"last_name"`
    Username     string    `json:"username"`
    Email        string    `json:"email"`
    PasswordHash string    `json:"-"` // "-" means this won't be included in JSON
    PhoneNumber  *string   `json:"phone_number,omitempty"`
    ProfileImage *string   `json:"profile_image,omitempty"`
    Role         string    `json:"role"`
    Status       string    `json:"status"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}

// For registration (without exposing password hash)
type UserRegister struct {
    FirstName   string `json:"first_name"`
    LastName    string `json:"last_name"`
    Username    string `json:"username"`
    Email       string `json:"email"`
    Password    string `json:"password"`
    PhoneNumber string `json:"phone_number,omitempty"`
}

// For login
type UserLogin struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

// For public profile (no sensitive data)
type UserProfile struct {
    ID           int64     `json:"id"`
    FirstName    string    `json:"first_name"`
    LastName     string    `json:"last_name"`
    Username     string    `json:"username"`
    PhoneNumber  *string   `json:"phone_number,omitempty"`
    ProfileImage *string   `json:"profile_image,omitempty"`
    CreatedAt    time.Time `json:"created_at"`
}