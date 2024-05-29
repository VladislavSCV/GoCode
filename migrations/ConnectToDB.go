package main

import (
	"database/sql"
	"fmt"
	"log"
    "encoding/json"
    "io/ioutil"

	_ "github.com/lib/pq"
)

type Config_db struct {
    Host     string `json:"host"`
    Port     int    `json:"port"`
    Username string `json:"username"`
    Password string `json:"password"`
    Dbname   string `json:"dbname"`
}

var db *sql.DB

var config Config_db

func init() {
    jsonFile, err := ioutil.ReadFile("../config/secrets/db.json")
    if err != nil {
        log.Fatal(err)
    }

    err = json.Unmarshal(jsonFile, &config)
    if err != nil {
        log.Fatal(err)
    }
}



func main() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        config.Host, config.Port, config.Username, config.Password, config.Dbname)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Successfully connected!")
}
