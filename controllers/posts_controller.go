package controllers

import (
	"basic-web-api/initalizers"
	"basic-web-api/models"
	"log"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	var req struct {
		Title string
		Body  string
	}

	c.Bind(&req)

	post := models.Post{Title: req.Title, Body: req.Body}
	result := initalizers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"msg": "internal server error",
		})

		log.Fatal(result.Error)
	} else {
		c.JSON(200, gin.H{
			"msg":  "Post has been created",
			"post": post,
		})
	}
}
