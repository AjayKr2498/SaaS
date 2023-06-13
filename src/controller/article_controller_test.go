package controller

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestArticleController_GetArticles(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name    string
		article *ArticleController
		args    args
	}{
		// TODO: Add test cases.
		{
			name:    "Get Articles",
			article: &ArticleController{},
			args:    args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			article := &ArticleController{}
			article.GetArticles(tt.args.c)
		})
	}
}

func TestArticleController_GetArticle(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name    string
		article *ArticleController
		args    args
	}{
		// TODO: Add test cases.
		{
			name:    "Get Artice",
			article: &ArticleController{},
			args: args{
				c: &gin.Context{
					Params: []gin.Param{
						gin.Param{
							Key:   "id",
							Value: "10",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			article := &ArticleController{}
			article.GetArticle(tt.args.c)
		})
	}
}
