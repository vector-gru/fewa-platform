package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "github.com/lesi/tutor_booking_system/models"
    "github.com/lesi/tutor_booking_system/pkg/database"
    "github.com/lesi/tutor_booking_system/pkg/logging"
    "github.com/lesi/tutor_booking_system/handlers"
)

func main() {
    // Load the .env file
    err := godotenv.Load("../../.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    logger := logging.NewLogger()
    database.InitDB()
    database.DB.AutoMigrate(&models.User{})

    r := mux.NewRouter()
    handlers.RegisterUserHandlers(r)

    srv := &http.Server{
        Addr:         ":8080",
        Handler:      r,
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            logger.Error("ListenAndServe():", err)
        }
    }()

    log.Println("Starting user service on port 8080...")

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
    <-stop

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        logger.Error("Server Shutdown Failed:", err)
    }
    log.Println("User service stopped")
}
