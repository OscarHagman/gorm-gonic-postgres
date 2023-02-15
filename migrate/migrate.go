package main

import (
	"basic-web-api/initalizers"
	"basic-web-api/models"
)

func init() {
	initalizers.LoadEnvVars()
	initalizers.ConnectToDB()
}

func main() {
	initalizers.DB.AutoMigrate(&models.Post{})
}
