package main

import (
    "log"
    "os/exec"
    "regexp"

    "github.com/gin-gonic/gin"
    "github.com/payne10/Void.git/internal/capture"
    "github.com/payne10/Void.git/internal/storage"
    "github.com/payne10/Void.git/internal/web"
)

type Device struct {
    IP     string `json:"ip"`
    MAC    string `json:"mac"`
    Vendor string `json:"vendor"`
}

func runArpScan() ([]Device, error) {
    cmd := exec.Command("arp-scan", "--localnet")
    output, err := cmd.Output()
    if err != nil {
        return nil, err
    }

    re := regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+)\s+([\dA-Fa-f:]+)\s+(.+)`)
    matches := re.FindAllStringSubmatch(string(output), -1)

    var devices []Device
    for _, match := range matches {
        devices = append(devices, Device{
            IP:     match[1],
            MAC:    match[2],
            Vendor: match[3],
        })
    }
    return devices, nil
}

func main() {

    // Initialize Storage
    storage.InitializeDB()

    // Start packet capture on interface eth0 (replace with your interface)
    go capture.StartCapture("any")

    // Initialize the Gin router
    r := gin.Default()

    // Setup routes
    web.SetupRoutes(r)

    // Add the ARP scan endpoint
    r.GET("/scan", func(c *gin.Context) {
        devices, err := runArpScan()
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, devices)
    })

    // Start the server
    log.Fatal(r.Run(":8080")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

