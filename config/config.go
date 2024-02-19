package config

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
)

func Config(key string) string {
    err := godotenv.Load(".env")
    if err != nil {
        fmt.Print("error loading environment variable")
    }
    return os.Getenv(key)
}