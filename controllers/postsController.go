package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
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

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))
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

func PostsUpdate(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error,
		})
		return
	}

	post.Title = body.Title
	post.Body = body.Body
	result = initializers.DB.Save(&post)
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