package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/testgo/go-sammy/initializers"
	"github.com/testgo/go-sammy/models"
)

func PostsCreate(c *gin.Context) {
	// Get data off a req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//Get id off url
	id := c.Param("id")
	// Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get id of url
	id := c.Param("id")
	// Get the data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)
}

func PostsStartPage(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"post": "Healthy",
	})
}
