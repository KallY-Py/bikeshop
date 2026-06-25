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

	// Basic validation
	if input.Email == "" || input.Password == "" || input.Username == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Email, password, and username are required")
		return
	}

	// Hash the password
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
		Role:         "user", // Default role
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
	utils.RespondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "User created successfully",
		"user_id": id,
		"user":    user,
	})
}

// GetProfile returns the current user's profile (protected route)
func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get user ID from context (set by auth middleware)
	userID, ok := r.Context().Value(utils.UserIDKey).(int)
	if !ok {
		utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// Fetch user from database
	user, err := h.repo.GetByID(int64(userID))
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch user profile")
		return
	}
	if user == nil {
		utils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	// Remove sensitive data before sending
	user.PasswordHash = "" // Clear password hash

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Profile retrieved successfully",
		"user":    user,
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
        FirstName   string `json:"first_name"`
        LastName    string `json:"last_name"`
        Username    string `json:"username"`
        Email       string `json:"email"`
        PhoneNumber string `json:"phone_number"`
        Bio         string `json:"bio"`
        Location    string `json:"location"`
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
    if input.Bio != "" {
        existingUser.Bio = &input.Bio
    }
    if input.Location != "" {
        existingUser.Location = &input.Location
    }
    if input.ProfileImage != "" {
        existingUser.ProfileImage = &input.ProfileImage
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
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Database error")
        return
    }
    if existingUser == nil {
        utils.RespondWithError(w, http.StatusNotFound, "User not found")
        return
    }

    // Update user fields
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
    if input.Bio != "" {
        existingUser.Bio = &input.Bio
    }
    if input.Location != "" {
        existingUser.Location = &input.Location
    }
    if input.ProfileImage != "" {
        existingUser.ProfileImage = &input.ProfileImage
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
        utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update user: "+err.Error())
        return
    }

    existingUser.PasswordHash = ""

    utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
        "message": "Profile updated successfully",
        "user":    existingUser,
    })
}

// GetCurrentUser returns the current authenticated user
func (h *UserHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get user ID from context
	userID, ok := r.Context().Value(utils.UserIDKey).(int)
	if !ok {
		utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	user, err := h.repo.GetByID(int64(userID))
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch user")
		return
	}
	if user == nil {
		utils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	// Remove sensitive data
	user.PasswordHash = ""

	utils.RespondWithJSON(w, http.StatusOK, user)
}

// DeleteUser deletes a user (admin only or self-deletion)
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Get current user role from context
	userRole, ok := r.Context().Value(utils.UserRoleKey).(string)
	if !ok {
		utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// Check if user is deleting themselves or is admin
	currentUserID, _ := r.Context().Value(utils.UserIDKey).(int)
	if userRole != "admin" && currentUserID != int(id) {
		utils.RespondWithError(w, http.StatusForbidden, "You can only delete your own account")
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
