package repository

import (
    "vue-go-backend/config"
    "vue-go-backend/models"
)

type CategoryRepository struct{}

func NewCategoryRepository() *CategoryRepository {
    return &CategoryRepository{}
}

func (r *CategoryRepository) GetAll() ([]models.Category, error) {
    rows, err := config.DB.Query("SELECT id, name, type, created_at FROM categories ORDER BY name")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var categories []models.Category
    for rows.Next() {
        var cat models.Category
        err := rows.Scan(&cat.ID, &cat.Name, &cat.Type, &cat.CreatedAt)
        if err != nil {
            return nil, err
        }
        categories = append(categories, cat)
    }
    return categories, nil
}

func (r *CategoryRepository) GetByType(categoryType string) ([]models.Category, error) {
    rows, err := config.DB.Query("SELECT id, name, type, created_at FROM categories WHERE type = ? ORDER BY name", categoryType)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var categories []models.Category
    for rows.Next() {
        var cat models.Category
        err := rows.Scan(&cat.ID, &cat.Name, &cat.Type, &cat.CreatedAt)
        if err != nil {
            return nil, err
        }
        categories = append(categories, cat)
    }
    return categories, nil
}