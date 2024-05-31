package controllers

import (
	"articles-api/initializers"
	"articles-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create an article
func CreateArticle(c *gin.Context) {
	var body struct {
		Title         string
		SubTitle      string
		Image         string
		Body          string
		MinutesToRead int64
		Comments      []models.Comment
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Get the user's id to add the article with
	user, _ := c.Get("user")
	article := models.Article{
		Title:         body.Title,
		SubTitle:      body.SubTitle,
		Image:         body.Image,
		Body:          body.Body,
		MinutesToRead: body.MinutesToRead,
		UserID:        user.(models.User).ID,
		// Comments:      body.Comments,
	}

	result := initializers.DB.Create(&article)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"New Article": article,
	})
}

// Get all articles
func GetAllArticles(c *gin.Context) {
	var articles []models.Article

	result := initializers.DB.Preload("Comments").Find(&articles)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, articles)
}

// Get one article by ID
func GetArticleById(c *gin.Context) {
	// Get the id of the article
	id := c.Param("id")
	var article models.Article

	result := initializers.DB.Preload("Comments").First(&article, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, article)
}

// Edit/Update an article
func UpdateAnArticle(c *gin.Context) {
	// Get the id of the post
	id := c.Param("id")

	// Get the data of the req body
	var body struct {
		Title         string
		SubTitle      string
		Image         string
		Body          string
		MinutesToRead int64
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Find the article we are updating
	var article models.Article

	initializers.DB.First(&article, id)

	user, _ := c.Get("user")

	if article.UserID != user.(models.User).ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden to edit other people's articles..."})
		return
	}
	// Update attributes with `struct`, will only update non-zero fields
	initializers.DB.Model(&article).Updates(models.Article{
		Title:         body.Title,
		SubTitle:      body.SubTitle,
		Image:         body.Image,
		Body:          body.Body,
		MinutesToRead: body.MinutesToRead,
		UserID:        user.(models.User).ID,
	})

	c.JSON(http.StatusOK, article)
}

// Delete an article
func DeleteArticle(c *gin.Context) {
	// Get id of article
	id := c.Param("id")

	// Get the article itself
	var article models.Article

	result := initializers.DB.First(&article, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// Get the owner of article
	user, _ := c.Get("user")

	// Delete the post only if it belongs to the owner
	if article.UserID != user.(models.User).ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden to delete other peoples articles"})
		return
	}
	initializers.DB.Delete(&models.Article{}, id)

	// Respond
	c.Status(http.StatusNoContent)
	c.JSON(200, gin.H{
		"post": "An article was deleted...",
	})
}

