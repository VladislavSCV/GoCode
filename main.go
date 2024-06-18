package main

import (
	"log"

	"github.com/VladislavSCV/GoCode/internal/db"
)

func main() {
	log.Println("Starting migrations...")
	db.ConnectToDB()
	log.Print(db.GetAllTasks())
}
