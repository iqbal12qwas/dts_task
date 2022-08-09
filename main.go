package main

import (
    "gin-api/models" 
    "gin-api/routes"
)

func main() {

    db := models.SetupDB()
    // db.AutoMigrate(&models.Task{})
	// db.AutoMigrate(&models.Profile{})

    r := routes.SetupRoutes(db)
    r.Run()
}