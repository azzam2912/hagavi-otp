package config

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {
    // load .env file
    err := godotenv.Load(".env")
    if err != nil {
        fmt.Print("error loading .env")
    }
    return os.Getenv(key)
}