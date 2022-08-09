package main

import (
        // book adalah directory root project go yang kita buat
    "gin-api/models" // memanggil package models pada directory models
    "gin-api/routes"
)

func main() {

    db := models.SetupDB()
    // db.AutoMigrate(&models.Task{})
	// db.AutoMigrate(&models.Profile{})

    r := routes.SetupRoutes(db)
    r.Run()
}