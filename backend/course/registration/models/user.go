package models

import (
    "gorm.io/gorm"
    "time"
)

type User struct {
    ID                   uint           `gorm:"primaryKey;autoIncrement"`
    FirstName            string         `json:"first_name" gorm:"size:50;not null"`
    LastName             string         `json:"last_name" gorm:"size:50;not null"`
    Email                string         `json:"email" gorm:"unique;size:100;not null"`
    Password             string         `json:"password" gorm:"size:255;not null"` // Changed json tag from "password_hash" to "password"
    PhoneNumber          string         `json:"phone_number" gorm:"size:20"`
    DateOfBirth          time.Time      `json:"date_of_birth"`
    Gender               string         `json:"gender" gorm:"size:10"`
    ProfilePicture       string         `json:"profile_picture"`
    StreetAddress        string         `json:"street_address" gorm:"size:255"`
    City                 string         `json:"city" gorm:"size:100"`
    State                string         `json:"state" gorm:"size:100"`
    PostalCode           string         `json:"postal_code" gorm:"size:20"`
    Country              string         `json:"country" gorm:"size:100"`
    PreferredLanguage    string         `json:"preferred_language" gorm:"size:50"`
    TimeZone             string         `json:"time_zone" gorm:"size:50"`
    AgreeToTerms         bool           `json:"agree_to_terms"`
    SubscribeToNewsletter bool          `json:"subscribe_to_newsletter"`
    Role                 string         `json:"role" gorm:"size:50;check:role IN ('tutor', 'admin', 'student')"`
    CreatedAt            time.Time      `json:"created_at" gorm:"default:current_timestamp"`
    UpdatedAt            time.Time      `json:"updated_at" gorm:"default:current_timestamp"`
}

// Migrate function to create or update the schema in the database
func Migrate(db *gorm.DB) error {
    return db.AutoMigrate(&User{})
}
