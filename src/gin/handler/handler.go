package handler

import (
	"gin/article"
	"gin/todo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ArticlesGet(articles *article.Articles) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result := articles.GetAll()
		ctx.JSON(http.StatusOK, result)
	}
}

type ArticlePostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description`
}

func ArticlePost(post *article.Articles) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestBody := ArticlePostRequest{}
		ctx.Bind(&requestBody)

		item := article.Item{
			Title:       requestBody.Title,
			Description: requestBody.Description,
		}
		post.Add(item)

		ctx.Status(http.StatusNoContent)
	}
}

func TodosGet(todos *todo.Todos) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result := todos.GetAll()
		ctx.JSON(http.StatusOK, result)
	}
}
