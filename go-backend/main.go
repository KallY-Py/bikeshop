package main

import (
	"fmt"
	"log"
	"net/http"

	"vue-go-backend/config"
	"vue-go-backend/handlers"
	"vue-go-backend/middleware"
	"vue-go-backend/repository"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	config.ConnectDB()
	defer config.DB.Close()

	r := mux.NewRouter()

	userHandler := handlers.NewUserHandler()
	listingHandler := handlers.NewListingHandler()
	categoryHandler := handlers.NewCategoryHandler()
	authHandler := handlers.NewAuthHandler()

	// Keep both dashboard and report handlers
	dashboardHandler := handlers.NewDashboardHandler()
	reportHandler := handlers.NewReportHandler(repository.NewReportRepository(config.DB))

	api := r.PathPrefix("/api").Subrouter()

	// Auth routes (public)
	api.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	api.HandleFunc("/auth/login", authHandler.Login).Methods("POST")

	// Public user routes
	api.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	api.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	api.HandleFunc("/users/{id}", userHandler.UpdateUserByID).Methods("PUT")

	// Dashboard routes
	api.HandleFunc("/dashboard/{userId}", dashboardHandler.GetUserDashboard).Methods("GET")
	api.HandleFunc("/dashboard/{userId}/listings", dashboardHandler.GetUserListings).Methods("GET")
	api.HandleFunc("/dashboard/{userId}/messages", dashboardHandler.GetUserMessages).Methods("GET")

	// Protected routes
	protected := api.PathPrefix("/protected").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/profile", userHandler.GetProfile).Methods("GET")
	protected.HandleFunc("/profile", userHandler.UpdateProfile).Methods("PUT")
	protected.HandleFunc("/me", userHandler.GetCurrentUser).Methods("GET")

	// Admin-only routes
	admin := api.PathPrefix("/admin").Subrouter()
	admin.Use(middleware.AuthMiddleware)
	admin.Use(middleware.AdminMiddleware)
	admin.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	admin.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	admin.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	admin.HandleFunc("/users/{id}/status", userHandler.UpdateUserStatus).Methods("PATCH")
	admin.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	// Report management routes
	admin.HandleFunc("/reports", reportHandler.GetReports).Methods("GET")
	admin.HandleFunc("/reports/{id}", reportHandler.UpdateReportStatus).Methods("PATCH")

	// Listing status update route
	admin.HandleFunc("/listings/{id}/status", listingHandler.UpdateListingStatus).Methods("PATCH")

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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		Debug:            true,
	})

	handler := c.Handler(r)

	port := ":3000"
	fmt.Printf("🚀 Bikeshop API server running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
