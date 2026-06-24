package repository

import (
    "database/sql"
    "vue-go-backend/config"
    "vue-go-backend/models"
)

type ListingRepository struct{}

func NewListingRepository() *ListingRepository {
    return &ListingRepository{}
}

func (r *ListingRepository) GetAll() ([]models.Listing, error) {
    query := `SELECT id, user_id, category_id, title, description, price, 
              condition_type, location, status, views, created_at, updated_at 
              FROM listings ORDER BY created_at DESC`
    
    rows, err := config.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var listings []models.Listing
    for rows.Next() {
        var listing models.Listing
        err := rows.Scan(
            &listing.ID, &listing.UserID, &listing.CategoryID,
            &listing.Title, &listing.Description, &listing.Price,
            &listing.ConditionType, &listing.Location, &listing.Status,
            &listing.Views, &listing.CreatedAt, &listing.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        listings = append(listings, listing)
    }
    return listings, nil
}

func (r *ListingRepository) GetByID(id int64) (*models.ListingWithDetails, error) {
    listing := &models.ListingWithDetails{}
    query := `SELECT l.id, l.user_id, l.category_id, l.title, l.description, l.price, 
              l.condition_type, l.location, l.status, l.views, l.created_at, l.updated_at,
              c.name, c.type, CONCAT(u.first_name, ' ', u.last_name) as seller_name
              FROM listings l
              JOIN categories c ON l.category_id = c.id
              JOIN users u ON l.user_id = u.id
              WHERE l.id = ?`
    
    err := config.DB.QueryRow(query, id).Scan(
        &listing.ID, &listing.UserID, &listing.CategoryID,
        &listing.Title, &listing.Description, &listing.Price,
        &listing.ConditionType, &listing.Location, &listing.Status,
        &listing.Views, &listing.CreatedAt, &listing.UpdatedAt,
        &listing.CategoryName, &listing.CategoryType, &listing.SellerName,
    )
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }

    // Get images
    images, _ := r.GetListingImages(id)
    listing.Images = images

    return listing, nil
}

func (r *ListingRepository) Create(listing *models.Listing) (int64, error) {
    query := `INSERT INTO listings (user_id, category_id, title, description, price, 
              condition_type, location) VALUES (?, ?, ?, ?, ?, ?, ?)`
    
    result, err := config.DB.Exec(query,
        listing.UserID, listing.CategoryID, listing.Title,
        listing.Description, listing.Price, listing.ConditionType,
        listing.Location,
    )
    if err != nil {
        return 0, err
    }
    return result.LastInsertId()
}

func (r *ListingRepository) GetListingImages(listingID int64) ([]string, error) {
    rows, err := config.DB.Query(
        "SELECT image_url FROM listing_images WHERE listing_id = ?", listingID,
    )
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var images []string
    for rows.Next() {
        var url string
        if err := rows.Scan(&url); err != nil {
            return nil, err
        }
        images = append(images, url)
    }
    return images, nil
}