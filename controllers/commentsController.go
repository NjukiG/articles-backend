package controllers

import (
	"articles-api/initializers"
	"articles-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Write/post a comment
func PostComment(c *gin.Context) {
	// Get the id of the articles you'll post a comment to.
	articleId := c.Param("id")

	var article models.Article

	result := initializers.DB.Preload("Comments").First(&article, articleId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Article not found",
		})
		return
	}

	var body struct {
		Content string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
	}

	// Get the user's id to add the article with
	user, _ := c.Get("user")

	comment := models.Comment{
		Content:   body.Content,
		UserID:    user.(models.User).ID,
		ArticleID: article.ID,
	}

	newComment := initializers.DB.Create(&comment)

	if newComment.Error != nil {
		c.Status(400)
		return
	}

	// Append the new comment to the article's Comments slice
	article.Comments = append(article.Comments, comment)

	c.JSON(http.StatusCreated, gin.H{
		"New Comment": comment,
	})

}

// Fetch all comments
func GetAllComments(c *gin.Context) {
	// Get the article id to the comments you want
	articleID := c.Param("id")

	var comments []models.Comment

	result := initializers.DB.Where("article_id = ?", articleID).Find(&comments)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Comments not found",
		})
		return
	}
	c.JSON(http.StatusOK, comments)
}

// Fetch one comment by ID
func GetCommentByID(c *gin.Context) {
	commentID := c.Param("id")

	var comment models.Comment

	result := initializers.DB.First(&comment, commentID)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Comment not found",
		})
		return
	}
	c.JSON(http.StatusOK, comment)
}

// Edit a comment
func EditComment(c *gin.Context) {
	// Get id of comment to edit
	commentID := c.Param("id")

	// Get the data of the req body
	var body struct {
		Content string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Find the comment we are updating
	var comment models.Comment

	// If comment isnt available return
	if err := initializers.DB.First(&comment, commentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// Get the owner of the comment and return if the arent allowed to edit someone else's comment
	user, _ := c.Get("user")
	if comment.UserID != user.(models.User).ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden to edit other's comments..."})
		return
	}

	// Update attributes with `struct`, will only update non-zero fields
	initializers.DB.Model(&comment).Updates(models.Comment{
		Content:   body.Content,
		UserID:    user.(models.User).ID,
		ArticleID: comment.ArticleID,
	})

	c.JSON(http.StatusOK, comment)
}

// DELETE A COMMENT
func DeleteComment(c *gin.Context) {
	// Get id of comment
	commentID := c.Param("id")

	// Get the comment itself
	var comment models.Comment

	result := initializers.DB.First(&comment, commentID)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// Get the owner of comment
	user, _ := c.Get("user")

	// Delete the comment only if it belongs to the owner
	if comment.UserID != user.(models.User).ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden to delete other peoples comments"})
		return
	}

	if err := initializers.DB.Delete(&models.Comment{}, commentID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	// Respond
	c.Status(http.StatusNoContent)
	c.JSON(200, gin.H{
		"post": "A comment was deleted...",
	})
}
