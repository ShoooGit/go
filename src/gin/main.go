package main

import (
	"gin/article"
	"gin/handler"
	dblib "gin/lib"
	"gin/todo"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	dblib.Init()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	///////////////////////////////////////////////////
	article := article.New()
	// json
	router.GET("/article", handler.ArticlesGet(article))
	router.POST("/article", handler.ArticlesGet(article))

	todo := todo.New()
	router.GET("/", handler.TodosGet(todo))
	///////////////////////////////////////////////////

	// //Index
	// router.GET("/", func(ctx *gin.Context) {
	// 	todos := dblib.GetAll()
	// 	ctx.HTML(200, "index.html", gin.H{
	// 		"todos": todos,
	// 	})
	// })

	//Create
	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dblib.Insert(text, status)
		ctx.Redirect(302, "/")
	})

	//Detail
	router.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := dblib.GetOne(id)
		ctx.HTML(200, "detail.html", gin.H{"todo": todo})
	})

	//Update
	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dblib.Update(id, text, status)
		ctx.Redirect(302, "/")
	})

	//削除確認
	router.GET("/delete_check/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		todo := dblib.GetOne(id)
		ctx.HTML(200, "delete.html", gin.H{"todo": todo})
	})

	//Delete
	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		dblib.Delete(id)
		ctx.Redirect(302, "/")

	})
	router.Run()
}
