package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "vue-go-backend/models"
    "vue-go-backend/repository"

    "github.com/gorilla/mux"
)

type ListingHandler struct {
    repo *repository.ListingRepository
}

func NewListingHandler() *ListingHandler {
    return &ListingHandler{
        repo: repository.NewListingRepository(),
    }
}

func (h *ListingHandler) GetListings(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    listings, err := h.repo.GetAll()
    if err != nil {
        http.Error(w, `{"error":"Failed to fetch listings"}`, http.StatusInternalServerError)
        return
    }
    
    json.NewEncoder(w).Encode(listings)
}

func (h *ListingHandler) GetListing(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    params := mux.Vars(r)
    id, err := strconv.ParseInt(params["id"], 10, 64)
    if err != nil {
        http.Error(w, `{"error":"Invalid listing ID"}`, http.StatusBadRequest)
        return
    }
    
    listing, err := h.repo.GetByID(id)
    if err != nil {
        http.Error(w, `{"error":"Database error"}`, http.StatusInternalServerError)
        return
    }
    if listing == nil {
        http.Error(w, `{"error":"Listing not found"}`, http.StatusNotFound)
        return
    }
    
    json.NewEncoder(w).Encode(listing)
}

func (h *ListingHandler) CreateListing(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    var listing models.Listing
    err := json.NewDecoder(r.Body).Decode(&listing)
    if err != nil {
        http.Error(w, `{"error":"Invalid request body"}`, http.StatusBadRequest)
        return
    }
    
    id, err := h.repo.Create(&listing)
    if err != nil {
        http.Error(w, `{"error":"Failed to create listing"}`, http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message":    "Listing created successfully",
        "listing_id": id,
    })
}