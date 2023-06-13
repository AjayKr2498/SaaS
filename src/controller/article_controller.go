package controller

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"test/service"
	"test/service/dto"

	"github.com/gin-gonic/gin"
)

// ArticleController
type ArticleController struct{}

// Instance of Service
var articleService = service.ArticleService{}

// GetArticles - Get All Articles
//
//	@Description	List all the articles from database
//	@BasePath		/cgh
//	@Tags			common
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.Articles
//	@Failure		500	{object}	dto.ErrorDTO
//	@Router			/articles [get]
func (article *ArticleController) GetArticles(c *gin.Context) {
	//Pass to service
	res, err := articleService.GetArticles()
	//check error
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorDTO{ErrorCode: "err-500", ErrorMessage: "Error in accessing articles"})
	}
	//return
	c.JSON(http.StatusOK, res)
}

// GetArticle - Get  Article By its id
//
//	@Description	Get Article  from database by its id
//	@BasePath		/cgh
//	@Tags			common
//	@Accept			json
//	@Produce		json
//	@param			id	path		string	true	"id"
//	@Success		200	{object}	dto.Article
//	@Success		404	{object}	dto.SuccessDTO
//	@Failure		500	{object}	dto.ErrorDTO
//	@Router			/article/{id} [get]
func (article *ArticleController) GetArticle(c *gin.Context) {
	//Get param value
	param := c.Param("id")
	//Convert string into int64
	id, err := strconv.ParseInt(param, 10, 64)
	//check error
	if err != nil {
		fmt.Printf("Error while converting string to int64 - %s", err)
		c.JSON(http.StatusBadRequest, dto.SuccessDTO{
			Message: "Error while converting param value string to int64",
		})
	}
	//Pass to service
	res, err := articleService.GetArticle(id)
	//check error
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorDTO{ErrorCode: "err-500", ErrorMessage: "Error in accessing articles"})
	}
	//check response, If it is empty
	if reflect.DeepEqual(dto.Article{}, res) {
		c.JSON(http.StatusNotFound, dto.SuccessDTO{
			Message: "Record not found",
		})
	}
	//return
	c.JSON(http.StatusOK, res)
}
