package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"vue-go-backend/config"
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

	var input struct {
		UserID        int64   `json:"user_id"`
		CategoryID    int     `json:"category_id"`
		Title         string  `json:"title"`
		Description   string  `json:"description"`
		Price         float64 `json:"price"`
		ConditionType string  `json:"condition_type"`
		Location      string  `json:"location"`
		ImageURL      string  `json:"image_url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	// Validate required fields
	if input.UserID == 0 || input.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "User ID and Title are required"})
		return
	}

	// Insert listing
	result, err := config.DB.Exec(`
        INSERT INTO listings (user_id, category_id, title, description, price, condition_type, location, status)
        VALUES (?, ?, ?, ?, ?, ?, ?, 'pending')
    `, input.UserID, input.CategoryID, input.Title, input.Description, input.Price, input.ConditionType, input.Location)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create listing: " + err.Error()})
		return
	}

	id, _ := result.LastInsertId()

	// Save image to listing_images table if provided
	if input.ImageURL != "" {
		_, err = config.DB.Exec("INSERT INTO listing_images (listing_id, image_url) VALUES (?, ?)", id, input.ImageURL)
		if err != nil {
			// Log but don't fail the request
			log.Printf("Failed to save image: %v", err)
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Listing created successfully",
		"id":      id,
	})
}

// UpdateListing updates an existing listing
func (h *ListingHandler) UpdateListing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid listing ID"})
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
		ImageURL      string  `json:"image_url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	_, err = config.DB.Exec(`
        UPDATE listings 
        SET title = ?, description = ?, price = ?, category_id = ?, 
            condition_type = ?, location = ?, status = ?
        WHERE id = ?
    `, input.Title, input.Description, input.Price, input.CategoryID,
		input.ConditionType, input.Location, input.Status, id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update listing"})
		return
	}

	// Update image if provided
	if input.ImageURL != "" {
		var imageID int64
		err = config.DB.QueryRow("SELECT id FROM listing_images WHERE listing_id = ? LIMIT 1", id).Scan(&imageID)
		if err == nil {
			config.DB.Exec("UPDATE listing_images SET image_url = ? WHERE id = ?", input.ImageURL, imageID)
		} else {
			config.DB.Exec("INSERT INTO listing_images (listing_id, image_url) VALUES (?, ?)", id, input.ImageURL)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Listing updated successfully"})
}

// DeleteListing deletes a listing
func (h *ListingHandler) DeleteListing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid listing ID"})
		return
	}

	// Delete images first (CASCADE should handle this, but just in case)
	config.DB.Exec("DELETE FROM listing_images WHERE listing_id = ?", id)

	// Delete the listing
	_, err = config.DB.Exec("DELETE FROM listings WHERE id = ?", id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete listing"})
		return
	}

	w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid listing ID"})
		return
	}

	var input struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil || input.Status == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Status is required"})
		return
	}

	validStatuses := map[string]bool{
		"pending": true, "approved": true, "rejected": true, "sold": true,
	}
	if !validStatuses[input.Status] {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid status value. Must be pending, approved, rejected, or sold"})
		return
	}

	_, err = config.DB.Exec("UPDATE listings SET status = ? WHERE id = ?", input.Status, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update listing status"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Listing status updated successfully",
		"id":      id,
		"status":  input.Status,
	})
}
