package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "strings"
    "time"

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

    // Serve uploaded files statically
    r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

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

    // File upload endpoint
    api.HandleFunc("/upload", uploadHandler).Methods("POST")

    // Category routes
    api.HandleFunc("/categories", categoryHandler.GetCategories).Methods("GET")

    // Health check
    api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"status":"ok","message":"Bikeshop API is running"}`))
    }).Methods("GET")

    // CORS configuration
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
        AllowedHeaders:   []string{"*"},
        AllowCredentials: false,
        Debug:            false,
    })

    handler := c.Handler(r)

    // Start server
    port := ":3000"
    fmt.Printf("🚀 Bikeshop API server running on http://localhost%s\n", port)
    fmt.Printf("📁 Uploads served from http://localhost%s/uploads/\n", port)

    log.Fatal(http.ListenAndServe(port, handler))
}

// Upload handler function
func uploadHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Limit upload size to 10MB
    if err := r.ParseMultipartForm(10 << 20); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "File too large. Max 10MB."})
        return
    }

    file, handler, err := r.FormFile("image")
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "No file uploaded. Please select a JPEG or PNG image."})
        return
    }
    defer file.Close()

    // Validate file type - JPEG or PNG only
    ext := strings.ToLower(filepath.Ext(handler.Filename))
    if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid file type. Only JPEG and PNG images are allowed."})
        return
    }

    // Validate file size (max 5MB)
    if handler.Size > 5*1024*1024 {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "File too large. Max 5MB."})
        return
    }

    // Create uploads directory if not exists
    if err := os.MkdirAll("uploads", 0755); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": "Server error. Please try again."})
        return
    }

    // Generate unique filename
    filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

    // Create destination file
    dst, err := os.Create(filepath.Join("uploads", filename))
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": "Failed to save file."})
        return
    }
    defer dst.Close()

    // Copy file
    if _, err := io.Copy(dst, file); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": "Failed to save file."})
        return
    }

    // Return the URL
    imageURL := fmt.Sprintf("http://localhost:3000/uploads/%s", filename)
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "image_url": imageURL,
        "message":   "File uploaded successfully",
    })
}