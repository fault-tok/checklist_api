package config

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
)

func LoadEnv() {
    files := []string{".env.development", ".env.production"}
    if err := godotenv.Load(files...); err != nil {
        fmt.Println("warning: no env file loaded:", err)
    }
}

func GetDataSourceName() (string, error) {
	LoadEnv()

    dbUser     := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbServer   := os.Getenv("DB_SERVER")
    dbPort     := os.Getenv("DB_PORT")
    dbName     := os.Getenv("DB_NAME")

    if dbUser == "" || dbPassword == "" || dbServer == "" || dbPort == "" || dbName == "" {
        return "", fmt.Errorf("missing env variable")
    }

    dataSourceName := fmt.Sprintf(
        "sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable&trustServerCertificate=true",
        dbUser, dbPassword, dbServer, dbPort, dbName,
    )
    return dataSourceName, nil
}