package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	post := models.Post{Title: "Hello", Body: "the body "}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error,
		})
		return
	}	

	c.JSON(200, gin.H{
		"post": post,
	})
}
