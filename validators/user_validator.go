package validators

import (
    "errors"
    "regexp"
    "github.com/lesi/tutor_booking_system/models"
)

func ValidateUser(user models.User) error {
    if user.FirstName == "" {
        return errors.New("first name is required")
    }
    if user.LastName == "" {
        return errors.New("last name is required")
    }
    if user.Email == "" {
        return errors.New("email is required")
    }
    if !isValidEmail(user.Email) {
        return errors.New("invalid email format")
    }
    if user.Password == "" {
        return errors.New("password is required")
    }
    if !isValidPassword(user.Password) {
        return errors.New("password must contain at least one special character")
    }
    if user.PhoneNumber == "" {
        return errors.New("phone number is required")
    }
    if user.DateOfBirth.IsZero() {
        return errors.New("date of birth is required")
    }
    if user.Gender == "" {
        return errors.New("gender is required")
    }
    if user.StreetAddress == "" {
        return errors.New("street address is required")
    }
    if user.City == "" {
        return errors.New("city is required")
    }
    if user.State == "" {
        return errors.New("state is required")
    }
    if user.PostalCode == "" {
        return errors.New("postal code is required")
    }
    if user.Country == "" {
        return errors.New("country is required")
    }
    if user.PreferredLanguage == "" {
        return errors.New("preferred language is required")
    }
    if user.TimeZone == "" {
        return errors.New("time zone is required")
    }
    if !user.AgreeToTerms {
        return errors.New("you must agree to the terms")
    }
    return nil
}

func isValidEmail(email string) bool {
    re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
    return re.MatchString(email)
}

func isValidPassword(password string) bool {
    re := regexp.MustCompile(`[!@#~$%^&*(),.?":{}|<>]`)
    return re.MatchString(password)
}
