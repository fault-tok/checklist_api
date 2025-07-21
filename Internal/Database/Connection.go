package Database

import (
    "gorm.io/driver/sqlserver"
    "gorm.io/gorm"
    "log"
    "CHECKLIST_API/Configs"
)

var Database *gorm.DB

func ConnectDatabase() {
    dataSourceName, err := config.GetDataSourceName()
    if err != nil {
        log.Fatalf("error message: %v", err)
    }

    database, err := gorm.Open(sqlserver.Open(dataSourceName), &gorm.Config{})
    if err != nil {
        log.Fatalf("error_message: failed to connect database: %v", err)
    }

    Database = database
    log.Println("✅ connected to sql server database")
}