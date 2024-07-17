package models

import "time"

type User struct {
    ID                   uint      `gorm:"primaryKey;autoIncrement"`
    FirstName            string    `json:"first_name"`
    LastName             string    `json:"last_name"`
    Email                string    `json:"email" gorm:"uniqueIndex"`
    Password             string    `json:"password"`
    AgreeToTerms         string    `json:"agree_to_terms"`  // Changed to string
    CreatedAt            time.Time `json:"created_at"`
    UpdatedAt            time.Time `json:"updated_at"`
}
