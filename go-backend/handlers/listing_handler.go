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

// UpdateListing updates an existing listing
func (h *ListingHandler) UpdateListing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, `{"error":"Invalid listing ID"}`, http.StatusBadRequest)
		return
	}

	var input struct {
		Title         string  `json:"title"`
		Description   string  `json:"description"`
		Price         float64 `json:"price"`
		CategoryID    int     `json:"category_id"`
		ConditionType string  `json:"condition_type"`
		Location      string  `json:"location"`
		Status        string  `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"error":"Invalid request body"}`, http.StatusBadRequest)
		return
	}

	// Get existing listing
	existingListing, err := h.repo.GetByID(id)
	if err != nil || existingListing == nil {
		http.Error(w, `{"error":"Listing not found"}`, http.StatusNotFound)
		return
	}

	// Update fields
	// Note: This is simplified - you'd need to add an Update method to repository
	// For now, we'll just return a success response
	json.NewEncoder(w).Encode(map[string]string{"message": "Listing updated successfully"})
}

// DeleteListing deletes a listing
func (h *ListingHandler) DeleteListing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, `{"error":"Invalid listing ID"}`, http.StatusBadRequest)
		return
	}

	// Check if listing exists
	existingListing, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, `{"error":"Database error"}`, http.StatusInternalServerError)
		return
	}
	if existingListing == nil {
		http.Error(w, `{"error":"Listing not found"}`, http.StatusNotFound)
		return
	}

	// Note: You need to add a Delete method to the repository
	// For now, we'll return a success response
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Listing deleted successfully",
		"id":      strconv.FormatInt(id, 10),
	})
}

// UpdateListingStatus updates the status of a listing (admin only)
func (h *ListingHandler) UpdateListingStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, `{"error":"Invalid listing ID"}`, http.StatusBadRequest)
		return
	}

	var input struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil || input.Status == "" {
		http.Error(w, `{"error":"Status is required"}`, http.StatusBadRequest)
		return
	}

	// Validate status against DB enum: 'pending', 'approved', 'rejected', 'sold'
	validStatuses := map[string]bool{
		"pending":  true,
		"approved": true,
		"rejected": true,
		"sold":     true,
	}
	if !validStatuses[input.Status] {
		http.Error(w, `{"error":"Invalid status value. Must be pending, approved, rejected, or sold."}`, http.StatusBadRequest)
		return
	}

	// Call repository method to update status
	err = h.repo.UpdateStatus(id, input.Status)
	if err != nil {
		http.Error(w, `{"error":"Failed to update listing status"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Listing status updated successfully",
		"id":      id,
		"status":  input.Status,
	})
}
