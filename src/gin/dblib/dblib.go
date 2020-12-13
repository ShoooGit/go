package dblib

import "github.com/jinzhu/gorm"

// モデルの定義
type Todo struct {
	gorm.Model
	Text   string
	Status string
}

type Item struct {
	Text   string `json:"text"`
	Status string `json:"status"`
}

// DBマイグレート
func Init() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースを開けません(dbInit)")
	}
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

// DB追加
func Insert(text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースを開けません(dbInsert)")
	}
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

// DB更新
func Update(id int, text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースを開けません(dbUpdate")
	}
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

// DB削除
func Delete(id int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースを開けません(dbDelete)")
	}
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

// DB全取得
func GetAll() []Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースを開けません(dbGetAll")
	}
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

// DB単独取得
func GetOne(id int) Item {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースを開けません(dbGetOne)")
	}
	var todo Todo
	var item Item
	db.First(&todo, id)
	db.Close()
	item.Text = todo.Text
	item.Status = todo.Status
	return item
}
