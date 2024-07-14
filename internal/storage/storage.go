package storage

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// InitializeDB initializes the database and creates the table
func InitializeDB() {
    var err error
    db, err = sql.Open("sqlite3", "./packets.db")
    if err != nil {
        log.Fatal(err)
    }

    createTableSQL := `CREATE TABLE IF NOT EXISTS packets (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "timestamp" DATETIME DEFAULT CURRENT_TIMESTAMP,
        "data" TEXT
    );`

    _, err = db.Exec(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }
}

// InsertPacket inserts a captured packet into the database
func InsertPacket(data string) {
    insertPacketSQL := `INSERT INTO packets (data) VALUES (?)`
    statement, err := db.Prepare(insertPacketSQL)
    if err != nil {
        log.Fatal(err)
    }
    _, err = statement.Exec(data)
    if err != nil {
        log.Fatal(err)
    }
}

// GetPackets retrieves all packets from the database
func GetPackets() []string {
    rows, err := db.Query("SELECT data FROM packets")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var packets []string
    for rows.Next() {
        var data string
        err = rows.Scan(&data)
        if err != nil {
            log.Fatal(err)
        }
        packets = append(packets, data)
    }
    return packets
}

