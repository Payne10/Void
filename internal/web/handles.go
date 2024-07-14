package web

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func SetupRoutes(router *gin.Engine) {
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "title": "Void Dashboard",
        })
    })
    // Add more routes as needed
}

