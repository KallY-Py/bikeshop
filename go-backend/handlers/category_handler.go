package handlers

import (
    "encoding/json"
    "net/http"

    "vue-go-backend/repository"
)

type CategoryHandler struct {
    repo *repository.CategoryRepository
}

func NewCategoryHandler() *CategoryHandler {
    return &CategoryHandler{
        repo: repository.NewCategoryRepository(),
    }
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    categoryType := r.URL.Query().Get("type")
    
    var categories interface{}
    var err error
    
    if categoryType != "" {
        categories, err = h.repo.GetByType(categoryType)
    } else {
        categories, err = h.repo.GetAll()
    }
    
    if err != nil {
        http.Error(w, `{"error":"Failed to fetch categories"}`, http.StatusInternalServerError)
        return
    }
    
    json.NewEncoder(w).Encode(categories)
}