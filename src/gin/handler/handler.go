package handler

import (
	"gin/article"
	"gin/dblib"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		todos := dblib.GetAll()
		ctx.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	}
}

func Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dblib.Insert(text, status)
		ctx.Redirect(302, "/")
	}
}

func Detail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := dblib.GetOne(id)
		// ctx.HTML(200, "detail.html", gin.H{"todo": todo})
		ctx.JSON(http.StatusOK, todo)
	}
}

func Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dblib.Update(id, text, status)
		ctx.Redirect(302, "/")
	}
}

func CheckDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		todo := dblib.GetOne(id)
		ctx.HTML(200, "delete.html", gin.H{"todo": todo})
	}
}

func Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		dblib.Delete(id)
		ctx.Redirect(302, "/")
	}
}

/////////////////////////////////////////////////////////////////////////////////
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

/////////////////////////////////////////////////////////////////////////////////
