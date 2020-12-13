package main

import (
	"gin/article"
	"gin/dblib"
	"gin/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	dblib.Init()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	//Index
	router.GET("/", handler.Index())

	//Create
	router.POST("/new", handler.Create())

	//Detail
	router.POST("/detail/:id", handler.Detail())

	//Update
	router.POST("/update/:id", handler.Update())

	//削除確認
	router.GET("/delete_check/:id", handler.CheckDelete())

	//Delete
	router.POST("/delete/:id", handler.Delete())

	///////////////////////////////////////////////////
	article := article.New()
	// json
	router.GET("/article", handler.ArticlesGet(article))
	router.POST("/article", handler.ArticlesGet(article))
	///////////////////////////////////////////////////

	router.Run()

}
