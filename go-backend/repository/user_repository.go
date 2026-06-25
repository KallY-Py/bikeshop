package repository

import (
    "database/sql"
    "vue-go-backend/config"
    "vue-go-backend/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
    return &UserRepository{}
}

func (r *UserRepository) GetAll() ([]models.User, error) {
    query := `SELECT id, first_name, last_name, username, email, 
              phone_number, profile_image, role, status, created_at, updated_at 
              FROM users ORDER BY created_at DESC`
    
    rows, err := config.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        err := rows.Scan(
            &user.ID, &user.FirstName, &user.LastName, &user.Username,
            &user.Email, &user.PhoneNumber, &user.ProfileImage,
            &user.Role, &user.Status, &user.CreatedAt, &user.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

func (r *UserRepository) GetByID(id int64) (*models.User, error) {
    user := &models.User{}
    query := `SELECT id, first_name, last_name, username, email, 
              phone_number, profile_image, role, status, created_at, updated_at 
              FROM users WHERE id = ?`
    
    err := config.DB.QueryRow(query, id).Scan(
        &user.ID, &user.FirstName, &user.LastName, &user.Username,
        &user.Email, &user.PhoneNumber, &user.ProfileImage,
        &user.Role, &user.Status, &user.CreatedAt, &user.UpdatedAt,
    )
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return user, nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
    user := &models.User{}
    query := `SELECT id, first_name, last_name, username, email, 
              password_hash, phone_number, profile_image, role, status, created_at, updated_at 
              FROM users WHERE email = ?`
    
    err := config.DB.QueryRow(query, email).Scan(
        &user.ID, &user.FirstName, &user.LastName, &user.Username,
        &user.Email, &user.PasswordHash, &user.PhoneNumber, &user.ProfileImage,
        &user.Role, &user.Status, &user.CreatedAt, &user.UpdatedAt,
    )
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return user, nil
}

func (r *UserRepository) Create(user *models.User) (int64, error) {
    query := `INSERT INTO users (first_name, last_name, username, email, password_hash, phone_number) 
              VALUES (?, ?, ?, ?, ?, ?)`
    
    result, err := config.DB.Exec(query,
        user.FirstName, user.LastName, user.Username,
        user.Email, user.PasswordHash, user.PhoneNumber,
    )
    if err != nil {
        return 0, err
    }
    return result.LastInsertId()
}

func (r *UserRepository) Update(user *models.User) error {
    query := `UPDATE users SET 
        first_name = ?, 
        last_name = ?, 
        username = ?, 
        email = ?, 
        phone_number = ?, 
        profile_image = ?,
        bio = ?,
        location = ?
        WHERE id = ?`
    
    _, err := config.DB.Exec(query,
        user.FirstName, 
        user.LastName, 
        user.Username,
        user.Email, 
        user.PhoneNumber, 
        user.ProfileImage,
        user.Bio,
        user.Location,
        user.ID,
    )
    return err
}

func (r *UserRepository) Delete(id int64) error {
    _, err := config.DB.Exec("DELETE FROM users WHERE id = ?", id)
    return err
}