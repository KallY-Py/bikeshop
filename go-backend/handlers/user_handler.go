package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "vue-go-backend/models"
    "vue-go-backend/repository"

    "github.com/gorilla/mux"
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
        http.Error(w, `{"error":"Failed to fetch users"}`, http.StatusInternalServerError)
        return
    }
    
    json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    params := mux.Vars(r)
    id, err := strconv.ParseInt(params["id"], 10, 64)
    if err != nil {
        http.Error(w, `{"error":"Invalid user ID"}`, http.StatusBadRequest)
        return
    }
    
    user, err := h.repo.GetByID(id)
    if err != nil {
        http.Error(w, `{"error":"Database error"}`, http.StatusInternalServerError)
        return
    }
    if user == nil {
        http.Error(w, `{"error":"User not found"}`, http.StatusNotFound)
        return
    }
    
    json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    var input models.UserRegister
    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil {
        http.Error(w, `{"error":"Invalid request body"}`, http.StatusBadRequest)
        return
    }
    
    // Basic validation
    if input.Email == "" || input.Password == "" || input.Username == "" {
        http.Error(w, `{"error":"Email, password, and username are required"}`, http.StatusBadRequest)
        return
    }
    
    // In production, you should hash the password!
    user := &models.User{
        FirstName:    input.FirstName,
        LastName:     input.LastName,
        Username:     input.Username,
        Email:        input.Email,
        PasswordHash: input.Password, // WARNING: Hash this in production!
    }
    
    if input.PhoneNumber != "" {
        user.PhoneNumber = &input.PhoneNumber
    }
    
    id, err := h.repo.Create(user)
    if err != nil {
        http.Error(w, `{"error":"Failed to create user: `+err.Error()+`"}`, http.StatusInternalServerError)
        return
    }
    
    user.ID = id
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "User created successfully",
        "user_id": id,
    })
}