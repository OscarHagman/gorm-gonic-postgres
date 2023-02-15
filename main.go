package main

import (
	"basic-web-api/controllers"
	"basic-web-api/initalizers"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initalizers.LoadEnvVars()
	initalizers.ConnectToDB()
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "pong",
		})
	})

	router.POST("/post", controllers.PostsCreate)

	runAt := "0.0.0.0:" + os.Getenv("PORT")
	router.Run(runAt)
}
