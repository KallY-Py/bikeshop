package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"vue-go-backend/models"
	"vue-go-backend/repository"
	"vue-go-backend/utils"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	repo *repository.UserRepository
}

// NewUserHandler creates a new instance of UserHandler
func NewUserHandler() *UserHandler {
	return &UserHandler{
		repo: repository.NewUserRepository(),
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := h.repo.GetAll()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch users")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, users)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	user, err := h.repo.GetByID(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database error")
		return
	}
	if user == nil {
		utils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}
	user.PasswordHash = "" // Don't send password hash
	utils.RespondWithJSON(w, http.StatusOK, user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input models.UserRegister
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if input.Email == "" || input.Password == "" || input.Username == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Email, password, and username are required")
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}
	user := &models.User{
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: string(hashedPassword),
		Role:         "user",
		Status:       "active",
	}
	if input.PhoneNumber != "" {
		user.PhoneNumber = &input.PhoneNumber
	}
	id, err := h.repo.Create(user)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create user: "+err.Error())
		return
	}
	user.ID = int(id)
	user.PasswordHash = ""
	utils.RespondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "User created successfully",
		"user_id": id,
		"user":    user,
	})
}

// UpdateUser updates a specific user by ID (admin only)
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	var input models.UserUpdate
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	existingUser, err := h.repo.GetByID(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database error")
		return
	}
	if existingUser == nil {
		utils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}
	if input.FirstName != "" {
		existingUser.FirstName = input.FirstName
	}
	if input.LastName != "" {
		existingUser.LastName = input.LastName
	}
	if input.Username != "" {
		existingUser.Username = input.Username
	}
	if input.Email != "" {
		existingUser.Email = input.Email
	}
	if input.PhoneNumber != "" {
		existingUser.PhoneNumber = &input.PhoneNumber
	}
	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to hash password")
			return
		}
		existingUser.PasswordHash = string(hashedPassword)
	}
	err = h.repo.Update(existingUser)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update user")
		return
	}
	existingUser.PasswordHash = ""
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "User updated successfully",
		"user":    existingUser,
	})
}

// UpdateUserByID updates a user by ID (for direct API calls)
func (h *UserHandler) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var input struct {
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Username     string `json:"username"`
		Email        string `json:"email"`
		PhoneNumber  string `json:"phone_number"`
		Bio          string `json:"bio"`
		Location     string `json:"location"`
		ProfileImage string `json:"profile_image"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	existingUser, err := h.repo.GetByID(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database error")
		return
	}
	if existingUser == nil {
		utils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	// Update fields if provided
	if input.FirstName != "" {
		existingUser.FirstName = input.FirstName
	}
	if input.LastName != "" {
		existingUser.LastName = input.LastName
	}
	if input.Username != "" {
		existingUser.Username = input.Username
	}
	if input.Email != "" {
		existingUser.Email = input.Email
	}
	if input.PhoneNumber != "" {
		existingUser.PhoneNumber = &input.PhoneNumber
	}

	err = h.repo.Update(existingUser)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update user: "+err.Error())
		return
	}

	existingUser.PasswordHash = ""

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "User updated successfully",
		"user":    existingUser,
	})
}

// UpdateUserStatus updates only the status field of a user (admin only)
func (h *UserHandler) UpdateUserStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	var input struct {
		Status string `json:"status"`
	}
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.Status == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Status is required")
		return
	}
	validStatuses := map[string]bool{"active": true, "suspended": true, "banned": true}
	if !validStatuses[input.Status] {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid status value")
		return
	}
	existingUser, err := h.repo.GetByID(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database error")
		return
	}
	if existingUser == nil {
		utils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}
	existingUser.Status = input.Status
	err = h.repo.Update(existingUser)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update user status")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "User status updated successfully",
		"user_id": id,
		"status":  input.Status,
	})
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	err = h.repo.Delete(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
	})
}

// Protected route handlers
func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID, ok := r.Context().Value(utils.UserIDKey).(int)
	if !ok {
		utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	user, err := h.repo.GetByID(int64(userID))
	if err != nil || user == nil {
		utils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}
	user.PasswordHash = ""
	utils.RespondWithJSON(w, http.StatusOK, user)
}

func (h *UserHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID, ok := r.Context().Value(utils.UserIDKey).(int)
	if !ok {
		utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	var input models.UserUpdate
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	existingUser, err := h.repo.GetByID(int64(userID))
	if err != nil || existingUser == nil {
		utils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}
	if input.FirstName != "" {
		existingUser.FirstName = input.FirstName
	}
	if input.LastName != "" {
		existingUser.LastName = input.LastName
	}
	if input.Username != "" {
		existingUser.Username = input.Username
	}
	if input.Email != "" {
		existingUser.Email = input.Email
	}
	if input.PhoneNumber != "" {
		existingUser.PhoneNumber = &input.PhoneNumber
	}
	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to hash password")
			return
		}
		existingUser.PasswordHash = string(hashedPassword)
	}
	err = h.repo.Update(existingUser)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update profile")
		return
	}
	existingUser.PasswordHash = ""
	utils.RespondWithJSON(w, http.StatusOK, existingUser)
}

func (h *UserHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	h.GetProfile(w, r) // Reuse GetProfile logic
}
