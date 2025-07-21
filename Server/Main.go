package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "CHECKLIST_API/Internal/Database"
    "CHECKLIST_API/Internal/Router"
)

func main() {
    Database.ConnectDatabase()

    router := gin.Default()
    Router.SetupRoutes(router) //เรียกฟังก์ชัน SetupRouter จาก internal

    // Run server
    if err := router.Run(":8080"); err != nil {
        log.Fatal("Failed to run server: ", err)
    }
}
