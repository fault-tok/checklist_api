package Database

import (
    "gorm.io/driver/sqlserver"
    "gorm.io/gorm"
    "log"
)

var Database *gorm.DB

func ConnectDatabase() {
	//sqlserver://<user>:<password>@<host>:<port>?database=<db_name>
    DataSourceName := "sqlserver://checklist_user:checklist1234@localhost:1433?database=CHECK_LIST"

    database, err := gorm.Open(sqlserver.Open(DataSourceName), &gorm.Config{})
    if err != nil {
        log.Fatal("❌ Failed to connect to database:", err)
    }

    Database = database
    log.Println("✅ Connected to SQL Server database")
}