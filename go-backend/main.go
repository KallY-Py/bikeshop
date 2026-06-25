package main

import (
	"fmt"
	"log"
	"net/http"

	"vue-go-backend/config"
	"vue-go-backend/handlers"
	"vue-go-backend/middleware"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Initialize database connection
	config.ConnectDB()
	defer config.DB.Close()

	// Create router
	r := mux.NewRouter()

	// Initialize handlers
	userHandler := handlers.NewUserHandler()
	listingHandler := handlers.NewListingHandler()
	categoryHandler := handlers.NewCategoryHandler()
	authHandler := handlers.NewAuthHandler()
	dashboardHandler := handlers.NewDashboardHandler()

	// API Routes
	api := r.PathPrefix("/api").Subrouter()

	// Auth routes (public)
	api.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	api.HandleFunc("/auth/login", authHandler.Login).Methods("POST")

	// User routes
	api.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	api.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	api.HandleFunc("/users/{id}", userHandler.UpdateUserByID).Methods("PUT")

	// Dashboard routes
	api.HandleFunc("/dashboard/{userId}", dashboardHandler.GetUserDashboard).Methods("GET")
	api.HandleFunc("/dashboard/{userId}/listings", dashboardHandler.GetUserListings).Methods("GET")
	api.HandleFunc("/dashboard/{userId}/messages", dashboardHandler.GetUserMessages).Methods("GET")

	// Protected routes example
	protected := api.PathPrefix("/protected").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/profile", userHandler.GetProfile).Methods("GET")

	// Listing routes
	api.HandleFunc("/listings", listingHandler.GetListings).Methods("GET")
	api.HandleFunc("/listings/{id}", listingHandler.GetListing).Methods("GET")
	api.HandleFunc("/listings", listingHandler.CreateListing).Methods("POST")
	api.HandleFunc("/listings/{id}", listingHandler.UpdateListing).Methods("PUT")    
	api.HandleFunc("/listings/{id}", listingHandler.DeleteListing).Methods("DELETE")

	// Category routes
	api.HandleFunc("/categories", categoryHandler.GetCategories).Methods("GET")

	// Health check
	api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok","message":"Bikeshop API is running"}`))
	}).Methods("GET")

	// CORS configuration - More permissive for development
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		Debug:            true,
	})

	handler := c.Handler(r)

	// Start server
	port := ":3000"
	fmt.Printf("🚀 Bikeshop API server running on http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, handler))
}
