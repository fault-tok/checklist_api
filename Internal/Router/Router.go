package Router

import (
    "github.com/gin-gonic/gin"
    "CHECKLIST_API/Internal/Handler"
)

func SetupRoutes(router *gin.Engine) {
    user := router.Group("/users")
    {
        user.GET("/:id", Handler.GetUserByID)
    }

    test := router.Group("/test")
    {
        test.GET("/", Handler.TestGet)
        test.POST("/", Handler.TestPost)
    }
}