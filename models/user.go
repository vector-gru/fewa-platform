/* package models

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
} */

package models

import "time"

type User struct {
    ID                   uint      `gorm:"primaryKey;autoIncrement"`
    FirstName            string    `json:"first_name"`
    LastName             string    `json:"last_name"`
    Email                string    `json:"email" gorm:"uniqueIndex"`
    Password             string    `json:"password"`
    PhoneNumber          string    `json:"phone_number"`
    DateOfBirth          time.Time `json:"date_of_birth"`
    Gender               string    `json:"gender"`
    ProfilePicture       string    `json:"profile_picture"`
    StreetAddress        string    `json:"street_address"`
    City                 string    `json:"city"`
    State                string    `json:"state"`
    PostalCode           string    `json:"postal_code"`
    Country              string    `json:"country"`
    PreferredLanguage    string    `json:"preferred_language"`
    TimeZone             string    `json:"time_zone"`
    AgreeToTerms         string    `json:"agree_to_terms"`  // Changed to string
    SubscribeToNewsletter bool      `json:"subscribe_to_newsletter"`
    CreatedAt            time.Time `json:"created_at"`
    UpdatedAt            time.Time `json:"updated_at"`
}
