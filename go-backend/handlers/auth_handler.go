package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"vue-go-backend/config"
	"vue-go-backend/models"
	"vue-go-backend/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

type RegisterRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginRequest struct {
	Identifier string `json:"identifier"` // email or username
	Password   string `json:"password"`
}

type AuthResponse struct {
	Token   string      `json:"token"`
	User    models.User `json:"user"`
	Message string      `json:"message"`
}

// Register handles user registration
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate required fields
	if req.FirstName == "" || req.LastName == "" || req.Username == "" ||
		req.Email == "" || req.Password == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "All fields are required")
		return
	}

	// Check if username already exists
	var existingUser models.User
	err := config.DB.QueryRow(
		"SELECT id FROM users WHERE username = ? OR email = ?",
		req.Username, req.Email,
	).Scan(&existingUser.ID)

	if err == nil {
		utils.RespondWithError(w, http.StatusConflict, "Username or email already exists")
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	// Insert user
	var user models.User
	result, err := config.DB.Exec(
		`INSERT INTO users (first_name, last_name, username, email, phone_number, password_hash, role, created_at, updated_at)
         VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		req.FirstName, req.LastName, req.Username, req.Email, req.PhoneNumber,
		string(hashedPassword), "user", time.Now(), time.Now(),
	)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}
	userID, err := result.LastInsertId()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch created user")
		return
	}
	user.ID = int(userID)
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Username = req.Username
	user.Email = req.Email
	if req.PhoneNumber != "" {
		user.PhoneNumber = &req.PhoneNumber
	}
	user.Role = "user"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Username, user.Role)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	response := AuthResponse{
		Token:   token,
		User:    user,
		Message: "Registration successful",
	}

	utils.RespondWithJSON(w, http.StatusCreated, response)
}

// Login handles user login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Identifier == "" || req.Password == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Identifier and password are required")
		return
	}

	// Find user by username or email
	var user models.User
	var passwordHash string
	err := config.DB.QueryRow(
		`SELECT id, first_name, last_name, username, email, phone_number, 
                password_hash, role, created_at, updated_at
         FROM users 
         WHERE username = ? OR email = ?`,
		req.Identifier, req.Identifier,
	).Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Username,
		&user.Email, &user.PhoneNumber, &passwordHash,
		&user.Role, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.Password))
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Username, user.Role)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	response := AuthResponse{
		Token:   token,
		User:    user,
		Message: "Login successful",
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}
