package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/lesi/tutor_booking_system/pkg/database"
    "github.com/lesi/tutor_booking_system/pkg/logging"
    "github.com/lesi/tutor_booking_system/services/user"
    "github.com/lesi/tutor_booking_system/models" // Ensure this import is correct
)

func main() {
    logger := logging.NewLogger()

    // Initialize the database connection
    database.InitDB()

    // Perform database migrations
    database.DB.AutoMigrate(&models.User{})
    database.DB.AutoMigrate(&models.Course{})

    userService := user.NewService()

    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            users, err := userService.GetAllUsers(context.Background())
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(users)
        } else {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    // Serve the static HTML files
    fs := http.FileServer(http.Dir("../../static"))
    http.Handle("/", http.StripPrefix("/", fs))

    srv := &http.Server{
        Addr:         ":8080",
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
