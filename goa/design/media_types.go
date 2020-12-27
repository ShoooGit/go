package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// レスポンスデータの定義
// MediaTypeに名前をつけます
var IntegerMedia = MediaType("application/vnd.integer+json", func()
{
	// 説明
	Description("example")
	// どのような値があるか（複数定義できる）
	Attributes(func() {
		// idはInteger型
		Attribute("id", Integer, "id", func(){
			// 返すレスポンスの例
			Example(1)
		})
		// レスポンスに必須な要素（基本は全て必須にしたほうが楽）
		Required("id")
	})
	View("default", func() {
		Attribute("id")
	})
})