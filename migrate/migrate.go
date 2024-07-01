package main

import (
	"github.com/anojaryal/json-crud-api-gin-gorm/initializers"
	"github.com/anojaryal/json-crud-api-gin-gorm/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}


func main() {
	// Migrate the schema
	initializers.DB.AutoMigrate(&models.Post{})
}