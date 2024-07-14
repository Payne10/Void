package web

import (
    "github.com/gin-gonic/gin"
    "github.com/payne10/void/internal/storage"
    "net/http"
)

// SetupRoutes sets up the Gin routes
func SetupRoutes(router *gin.Engine) {
    router.LoadHTMLGlob("internal/web/templates/*")
    router.Static("/static", "./internal/web/static")

    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "title": "Void Dashboard",
        })
    })

    router.GET("/packets", func(c *gin.Context) {
        packets := storage.GetPackets()
        c.JSON(http.StatusOK, packets)
    })
}

