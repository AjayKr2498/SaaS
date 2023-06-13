package service

import (
	"fmt"
	"test/repository"
	"test/service/dto"
)

// ArticleService
type ArticleService struct{}

// Repository Instance
var articleRepository repository.ArticleRepository = repository.ArticleRepository{}

// GetArticles
func (articleService *ArticleService) GetArticles() (dto.Articles, error) {
	//Pass to repository
	res, err := articleRepository.GetArticles()
	//check err
	if err != nil {
		fmt.Printf("gGetting error in - GetArticles - %s", err)
		return dto.Articles{}, err
	}
	//return
	return res, nil
}

// GetArticle
func (articleService *ArticleService) GetArticle(id int64) (dto.Article, error) {
	//Pass to repository
	res, err := articleRepository.GetArticle(id)
	//check err
	if err != nil {
		fmt.Printf("gGetting error in - GetArticle - %s", err)
		return dto.Article{}, err
	}
	//return
	return res, nil
}
