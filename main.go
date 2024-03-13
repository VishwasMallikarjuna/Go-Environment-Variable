package main

import (
    "fmt"
    "os"
)

func main() {
    // Get the value of the DATABASE_URL environment variable
    databaseURL := os.Getenv("DATABASE_URL")
    fmt.Println("DATABASE_URL:", databaseURL)
}
