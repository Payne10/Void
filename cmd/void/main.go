package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/payne10/Void.git/internal/capture"
    "github.com/payne10/Void.git/internal/storage"
    "github.com/payne10/Void.git/internal/web"
)

func main() {

    //Initialize Storage
    storage.InitializeDB()
    
    // Start packet capture on interface eth0 (replace with your interface)
    go capture.StartCapture("any")

    // Initialize the Gin router
    r := gin.Default()

    web.SetupRoutes(r)


    // Start the server
    log.Fatal(r.Run(":8080")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

