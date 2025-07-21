package Handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func GetUserByID(context *gin.Context) {
    id := context.Param("id")
    context.JSON(http.StatusOK, gin.H{
        "id":   id,
        "name": "Test User",
    })
}
