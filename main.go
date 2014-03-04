package main

import (
    _ "code.google.com/p/odbc"
    "database/sql"
    _ "flag"
    _ "io/ioutil"
    "log"
    "os"
    _ "path/filepath"
)

var (
    mssrv       = flag.String("mssrv", "server", "ms sql server name")
    msdb        = flag.String("msdb", "dbname", "ms sql server database name")
    msdriver    = flag.String("msdriver", "sql server", "ms sql odbc driver name")
)

var wd string
var tx sql.Tx

func connect() (db *sql.DB, err error) {
    params := map[string]string {
        "driver"                : *msdriver,
        "server"                : *mssrv,
        "database"              : *msdb,
        "trusted_connection"    : "true"
    }
    var c string
    for n, v := range params {
        c += n + "=" + v + ";"
    }
    db, err = sql.Open("odbc", c)
    if err != nil {
        return nil, err
    }
    return db, nil
}

func main() {
    var err error
    wd, err = os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    params := map[string]string {
        "driver": "sql server",
        "server": "localhost\\localdb",
        "database": "sandbox",
        "trusted_connection": "yes",
    }
    var c string
    for n, v := range params {
        c += n + "="+ v + ";"
    }
    var db *sql.DB
    db, err = sql.Open("odbc", c)
    if err != nil {
        log.Fatal(err)
    }

    err = db.Close()
    if err != nil {
        log.Fatal(err)
    }
}