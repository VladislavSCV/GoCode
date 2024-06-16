package main

import (
	"log"

	"github.com/VladislavSCV/GoCode/migrations"
)

func main() {
	log.Println("Starting migrations...")
	migrations.ConnectToDB()
	log.Print(migrations.GetAllTasks())
}
