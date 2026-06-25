package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "vue-go-backend/config"

    "github.com/gorilla/mux"
)

type DashboardHandler struct{}

func NewDashboardHandler() *DashboardHandler {
    return &DashboardHandler{}
}

// GetUserDashboard returns user dashboard stats
func (h *DashboardHandler) GetUserDashboard(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    params := mux.Vars(r)
    userID, _ := strconv.ParseInt(params["userId"], 10, 64)
    
    dashboard := make(map[string]interface{})
    
    // Active listings count
    var activeListings int
    config.DB.QueryRow(
        "SELECT COUNT(*) FROM listings WHERE user_id = ? AND status IN ('active', 'approved')",
        userID,
    ).Scan(&activeListings)
    dashboard["active_listings"] = activeListings
    
    // Total sales
    var totalSales float64
    config.DB.QueryRow(
        "SELECT COALESCE(SUM(price), 0) FROM listings WHERE user_id = ? AND status = 'sold'",
        userID,
    ).Scan(&totalSales)
    dashboard["total_sales"] = totalSales
    
    // Total listings
    var totalListings int
    config.DB.QueryRow(
        "SELECT COUNT(*) FROM listings WHERE user_id = ?", userID,
    ).Scan(&totalListings)
    dashboard["total_listings"] = totalListings
    
    // Unread messages count
    var unreadMessages int
    config.DB.QueryRow(`
        SELECT COUNT(*) FROM messages m
        JOIN conversations c ON m.conversation_id = c.id
        WHERE (c.buyer_id = ? OR c.seller_id = ?) 
        AND m.sender_id != ? 
        AND m.is_read = 0
    `, userID, userID, userID).Scan(&unreadMessages)
    dashboard["unread_messages"] = unreadMessages
    
    json.NewEncoder(w).Encode(dashboard)
}

    func (h *DashboardHandler) GetUserListings(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        
        params := mux.Vars(r)
        userID, _ := strconv.ParseInt(params["userId"], 10, 64)
        
        // Use subquery for image to avoid GROUP BY issues
        rows, err := config.DB.Query(`
            SELECT l.id, l.title, COALESCE(l.description, '') as description, 
                l.price, l.status, l.condition_type, 
                COALESCE(l.location, '') as location, COALESCE(l.views, 0) as views, 
                l.created_at, l.updated_at,
                COALESCE(c.name, '') as category_name,
                COALESCE((SELECT li.image_url FROM listing_images li WHERE li.listing_id = l.id LIMIT 1), '') as image_url
            FROM listings l
            LEFT JOIN categories c ON l.category_id = c.id
            WHERE l.user_id = ?
            ORDER BY l.created_at DESC
        `, userID)
        if err != nil {
            http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
            return
        }
        defer rows.Close()
        
        var listings []map[string]interface{}
        for rows.Next() {
            var id int64
            var title, description, status, conditionType, location, categoryName, imageUrl string
            var price float64
            var views int
            var createdAt, updatedAt string
            
            err := rows.Scan(&id, &title, &description, &price, &status,
                &conditionType, &location, &views, &createdAt, &updatedAt,
                &categoryName, &imageUrl)
            if err != nil {
                continue
            }
            
            listings = append(listings, map[string]interface{}{
                "id":             id,
                "title":          title,
                "description":    description,
                "price":          price,
                "status":         status,
                "condition_type": conditionType,
                "location":       location,
                "views":          views,
                "category":       categoryName,
                "image_url":      imageUrl,
                "created_at":     createdAt,
                "updated_at":     updatedAt,
            })
        }
        
        if listings == nil {
            listings = []map[string]interface{}{}
        }
        
        json.NewEncoder(w).Encode(listings)
    }

// GetUserMessages returns user's messages
func (h *DashboardHandler) GetUserMessages(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    params := mux.Vars(r)
    userID, _ := strconv.ParseInt(params["userId"], 10, 64)
    
    rows, err := config.DB.Query(`
        SELECT m.id, m.message, m.sent_at, m.is_read,
               u.first_name, u.last_name, COALESCE(u.profile_image, '') as profile_image
        FROM messages m
        JOIN conversations c ON m.conversation_id = c.id
        JOIN users u ON m.sender_id = u.id
        WHERE (c.buyer_id = ? OR c.seller_id = ?)
        AND m.sender_id != ?
        ORDER BY m.sent_at DESC
        LIMIT 10
    `, userID, userID, userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()
    
    var messages []map[string]interface{}
    for rows.Next() {
        var id int64
        var message, sentAt, firstName, lastName, profileImage string
        var isRead bool
        rows.Scan(&id, &message, &sentAt, &isRead, &firstName, &lastName, &profileImage)
        
        messages = append(messages, map[string]interface{}{
            "id":            id,
            "message":       message,
            "sent_at":       sentAt,
            "is_read":       isRead,
            "sender_name":   firstName + " " + lastName,
            "sender_image":  profileImage,
        })
    }
    
    if messages == nil {
        messages = []map[string]interface{}{}
    }
    
    json.NewEncoder(w).Encode(messages)
}