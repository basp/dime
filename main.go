package main

import (
    _ "code.google.com/p/odbc"
    "database/sql"
    "log"
    "os"
)

func main() {
    wd, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("%s", wd)
    connStr := "Driver={SQL Server Native Client 11.0};Server=localhost\\localdb;Database=sandbox;Trusted_Connection=yes"
    db, err := sql.Open("odbc", connStr)
    if err != nil {
        log.Fatal(err)
    }
    rows, err := db.Query("SELECT name FROM sys.objects")
    if err != nil {
        log.Fatal(err)
    }
    for rows.Next() {
        var name string
        if err := rows.Scan(&name); err != nil {
            log.Fatal(err)
        }
        log.Printf("%s", name)
    }
    db.Close()
    log.Printf("ok.")
}