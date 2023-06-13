package main

import (
	"test/controller"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "test/docs"

	"github.com/gin-gonic/gin"
)

// @title		Swagger Documentation for API
// @version	1.0
func main() {
	//Gin Gonic
	r := gin.Default()
	//Swagger
	docs.SwaggerInfo.BasePath = "cgh"
	//Link For Swagger test
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//Routes
	setupRoutes(r)
}

// Routes
func setupRoutes(r *gin.Engine) {
	// Instantiate controllers
	article := controller.ArticleController{}

	// Set application context in URL - Do not edit this
	articles := r.Group("/cgh")
	{
		articles.GET("/articles", article.GetArticles)
		articles.GET("/article/:id", article.GetArticle)
	}
}
