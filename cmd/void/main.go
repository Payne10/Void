package main

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "githb.com/payne10/Void/internal/capture"
)

func main() {
    // Start packet capture on interface eth0 (replace with your interface)
    go capture.StartCapture("eth0")

    // Initialize the Gin router
    r := gin.Default()

    // Set up routes
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })

    r.GET("/packets", func(c *gin.Context) {
        packets := capture.GetPackets()
        // Simplified representation of packets for JSON response
        var packetData []string
        for _, packet := range packets {
            packetData = append(packetData, packet.String())
        }
        c.JSON(http.StatusOK, packetData)
    })

    // Start the server
    log.Fatal(r.Run(":8080")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

