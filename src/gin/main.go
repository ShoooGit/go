package main

import (
	"gin/article"
	"gin/controller"
	"gin/dblib"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// DBの初期化
	dblib.Init()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	// ルーティング
	router.GET("/", controller.Index)
	router.POST("/new", controller.Create)
	router.GET("/detail/:id", controller.Detail)
	router.POST("/update/:id", controller.Update)
	router.GET("/delete_check/:id", controller.CheckDelete)
	router.POST("/delete/:id", controller.Delete)

	///////////////////////////////////////////////////
	// jsonテスト
	article := article.New()
	router.GET("/article", controller.ArticlesGet(article))
	router.POST("/article", controller.ArticlesGet(article))
	///////////////////////////////////////////////////

	router.Run(":3000")

}
