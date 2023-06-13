package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"test/model"
	"test/service/dto"
)

// ArticleRepository
type ArticleRepository struct{}

// Read Data from JOSN file
func Database() ([]model.Article, error) {
	//Bind data according to Article Model Structure
	var article []model.Article
	//Read JSON file
	articles, err := ioutil.ReadFile("data/data.json")
	//Check for error
	if err != nil {
		fmt.Printf("Error while reading the file :: %s", err)
		return []model.Article{}, err
	}
	//Unmarshal the data
	err = json.Unmarshal([]byte(articles), &article)
	//Check for error
	if err != nil {
		fmt.Printf("Error while unmarshal JSON :: %s", err)
		return []model.Article{}, err
	}
	//return
	return article, nil
}

// GetArticles
func (repo ArticleRepository) GetArticles() (dto.Articles, error) {
	//To decalre article dto
	var articleDTO []dto.Article

	//Fecth data from database
	articles, _ := Database()

	//Loop over articles
	for _, article := range articles {
		//Set article to dto
		articleDTO = append(articleDTO, dto.Article{
			Id:     article.Id,
			Title:  article.Title,
			Body:   article.Body,
			Status: article.Status,
		})
	}

	//return
	return dto.Articles{
		Articles: articleDTO,
		Total:    len(articles),
	}, nil
}

// GetArticle
func (repo ArticleRepository) GetArticle(id int64) (dto.Article, error) {

	//Fecth data from database
	articles, _ := Database()

	//Loop over articles
	for _, article := range articles {
		if article.Id == id {
			return dto.Article{
				Id:     article.Id,
				Title:  article.Title,
				Body:   article.Body,
				Status: article.Status,
			}, nil
		}
	}

	//return
	return dto.Article{}, nil
}
