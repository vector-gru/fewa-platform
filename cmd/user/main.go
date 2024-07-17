package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/lesi/tutor_booking_system/handlers"
    "github.com/lesi/tutor_booking_system/models"
    "github.com/lesi/tutor_booking_system/pkg/database"
    "github.com/lesi/tutor_booking_system/services"
)

func main() {
    // Initialize database
    db, err := database.InitDB()
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }

    // Auto migrate User model
    err = db.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    userService := services.NewUserService(db)
    userHandler := handlers.NewUserHandler(userService)

    http.HandleFunc("/register", userHandler.RegisterUser)

    // Serve index.html at /index.html
    http.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "../../static/index.html")
    })

    // Serve static files
    fs := http.FileServer(http.Dir("../../static"))
    http.Handle("/", http.StripPrefix("/", fs))

    srv := &http.Server{
        Addr:         ":8080",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Printf("ListenAndServe(): %v", err)
        }
    }()

    log.Println("Starting user service on port 8080...")

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
    <-stop

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        log.Printf("Server Shutdown Failed: %v", err)
    }
    log.Println("User service stopped")
}
