package Handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type TestPostRequest struct {
    Name string `json:"name"`
}

func TestGet(context *gin.Context) {
    testValue := context.Query("test")
    context.JSON(http.StatusOK, gin.H{
        "test":   testValue,
    })
}

func TestPost(context *gin.Context) {
    var req TestPostRequest

    // bind JSON body เข้ากับ struct
    if err := context.ShouldBindJSON(&req); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // response
    context.JSON(http.StatusOK, gin.H{
        "message": "JSON received",
        "name":    req.Name,
    })
}
