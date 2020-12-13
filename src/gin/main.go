package main

import (
	"gin/article"
	"gin/dblib"
	"gin/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// DBの初期化
	dblib.Init()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	// ルーティング
	router.GET("/", handler.Index())
	router.POST("/new", handler.Create())
	router.GET("/detail/:id", handler.Detail())
	router.POST("/update/:id", handler.Update())
	router.GET("/delete_check/:id", handler.CheckDelete())
	router.POST("/delete/:id", handler.Delete())

	///////////////////////////////////////////////////
	// jsonテスト
	article := article.New()
	router.GET("/article", handler.ArticlesGet(article))
	router.POST("/article", handler.ArticlesGet(article))
	///////////////////////////////////////////////////

	router.Run(":3000")

}
