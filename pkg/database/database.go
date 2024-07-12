package database

type Database interface {
    // Define database methods
}

type database struct {
    // DB connection fields
}

func NewDatabase() Database {
    return &database{}
}
