package controllers

import (
	models "go-api/Models"
	"go-api/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func GetPost(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func GetPostById(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, "id = ?", id)

	c.JSON(http.StatusOK, gin.H{
		"posts": post,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	var post models.Post
	r := initializers.DB.First(&post, "id = ?", id)

	if r.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Record not found",
		})
		return
	}

	post.Body = body.Body
	post.Title = body.Title
	result := initializers.DB.Save(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"affected": result.RowsAffected,
		"post":     post,
	})
}
