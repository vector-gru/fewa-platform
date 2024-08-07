// Save this as main.go
package main

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

func main() {
    password := "password*A11"
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    fmt.Println("Hashed Password:", string(hashedPassword))

    err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
    if err != nil {
        fmt.Println("Password does not match")
    } else {
        fmt.Println("Password matches")
    }
}
