package models

import (
    "gorm.io/gorm"
    "time"
)

type Course struct {
    gorm.Model
    Title       string    `json:"title"`
    Description string    `json:"description"`
    TutorID     uint      `json:"tutor_id"`
    Duration    int       `json:"duration"`    // Duration in hours
    StartDate   time.Time `json:"start_date"`
    EndDate     time.Time `json:"end_date"`
    Category    string    `json:"category"`
    Level       string    `json:"level"`
}
