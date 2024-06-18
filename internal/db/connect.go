package db

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    _ "github.com/lib/pq"
)

type Config struct {
    Data_db Config_db `json:"data_db"`
}

type Config_db struct {
    Host     string `json:"host"`
    Port     int    `json:"port"`
    Username string `json:"username"`
    Password string `json:"password"`
    Dbname   string `json:"dbname"`
}

var (
    config Config
    DB *sql.DB
)

func init() {
    jsonFile, err := ioutil.ReadFile(`C:\Users\VladislavSCV\OneDrive\Desktop\VSC\GO\GoCode\internal\config\secrets\db.json`)
    if err != nil {
        log.Fatal(err)
    }

    err = json.Unmarshal(jsonFile, &config)
    if err != nil {
        log.Fatal(err)
    }
}

func ConnectToDB() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    config.Data_db.Host, config.Data_db.Port, config.Data_db.Username, config.Data_db.Password, config.Data_db.Dbname)

    var err error
    DB, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal(err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Successfully connected!")
}
