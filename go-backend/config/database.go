package config

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
    // XAMPP default: username=root, no password
    dsn := "root:@tcp(127.0.0.1:3306)/bikeshop?parseTime=true&loc=Local"
    
    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Connection pool settings
    DB.SetMaxOpenConns(25)
    DB.SetMaxIdleConns(5)
    DB.SetConnMaxLifetime(5 * time.Minute)

    // Test the connection
    err = DB.Ping()
    if err != nil {
        log.Fatal("Database ping failed:", err)
    }

    fmt.Println("Successfully connected to bikeshop database!")
}