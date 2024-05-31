package main

import (
	"articles-api/controllers"
	"articles-api/initializers"
	"articles-api/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDb()
}

func main() {
	r := gin.Default()
	publicRoutes := r.Group("/public")
	{
		// Authentication routes
		publicRoutes.POST("/signup", controllers.SignUp)
		publicRoutes.POST("/login", controllers.Login)
	}

	protectedRoutes := r.Group("/protected")
	protectedRoutes.Use(middleware.RequireAuth)

	{
		// Articles routes
		protectedRoutes.GET("/validate", controllers.Validate)
		protectedRoutes.GET("/articles", controllers.GetAllArticles)
		protectedRoutes.GET("/articles/:id", controllers.GetArticleById)
		protectedRoutes.POST("/articles", controllers.CreateArticle)
		protectedRoutes.PUT("/articles/:id", controllers.UpdateAnArticle)
		protectedRoutes.DELETE("/articles/:id", controllers.DeleteArticle)
		// Comments routes
		protectedRoutes.GET("/articles/:id/comments", controllers.GetAllComments)
		protectedRoutes.GET("/comments/:id", controllers.GetCommentByID)
		protectedRoutes.POST("/articles/:id/comments", controllers.PostComment)
		protectedRoutes.PUT("/comments/:id", controllers.EditComment)
		protectedRoutes.DELETE("/comments/:id", controllers.DeleteComment)
	}

	r.Run()
}
