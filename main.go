package main

import (
    _ "code.google.com/p/odbc"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
)

var wd string
var tx sql.Tx

func walkDatabases(path string, walkFunc func(dir string) error) error {
    entries, err := ioutil.ReadDir(path)
    if err != nil {
        log.Fatal(err)
    }
    for _, e := range entries {
        if e.IsDir() {
            dir, err := filepath.Abs(e.Name())
            if err != nil {
                return err
            }
            walkFunc(dir)
        }
    }
    return nil
}

func walkScripts(dir string) error {
    walkFunc := func(path string, info os.FileInfo, err error) error {
        if filepath.Ext(path) == ".sql" {
            log.Printf(path)                        
        }
        return err
    }
    log.Printf("%s", wd)
    filepath.Walk(wd, walkFunc)
    return nil
}

func main() {
    var err error
    wd, err = os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    dbdir := filepath.Join(wd, "databases")
    walkFunc := func(dir string) error {
        log.Printf(dir)
        return nil
    }
    walkDatabases(dbdir, walkFunc)
}