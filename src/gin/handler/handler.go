package handler

import (
	"gin/article"
	"gin/dblib"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		todos := dblib.GetAll()
		c.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	}
}

func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		text := c.PostForm("text")
		status := c.PostForm("status")
		dblib.Insert(text, status)
		c.Redirect(302, "/")
	}
}

func Detail() gin.HandlerFunc {
	return func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := dblib.GetOne(id)
		c.HTML(200, "detail.html", gin.H{"todo": todo})
		// c.JSON(http.StatusOK, todo)
	}
}

func Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		text := c.PostForm("text")
		status := c.PostForm("status")
		dblib.Update(id, text, status)
		c.Redirect(302, "/")
	}
}

func CheckDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		todo := dblib.GetOne(id)
		c.HTML(200, "delete.html", gin.H{"todo": todo})
	}
}

func Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		dblib.Delete(id)
		c.Redirect(302, "/")
	}
}

/////////////////////////////////////////////////////////////////////////////////
func ArticlesGet(articles *article.Articles) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := articles.GetAll()
		c.JSON(http.StatusOK, result)
	}
}

type ArticlePostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description`
}

func ArticlePost(post *article.Articles) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := ArticlePostRequest{}
		c.Bind(&requestBody)

		item := article.Item{
			Title:       requestBody.Title,
			Description: requestBody.Description,
		}
		post.Add(item)

		c.Status(http.StatusNoContent)
	}
}

/////////////////////////////////////////////////////////////////////////////////
