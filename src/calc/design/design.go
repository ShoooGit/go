package design

import (
	. "goa.design/goa/v3/dsl"
)

// API定義
var _ = API("calc", func() {
	// APIの説明（タイトルと説明）
	Title("Caluculator Service")
	Description("Service for adding numbes, a Goa teaser")

	// サーバ定義
	// Server はクライアントのリクエストを受け付ける単一のプロセスを記述
	// DSLは、サーバーがホストする一連のサービスとホストの詳細を定義
	Server("calc", func() {
		Description("calc hosts the Calculator Service.")

		// このサーバーによってホストされているサービスを列挙します
		Services("calc")

		// ホストとそのトランスポート URL を列挙します。
		Host("development", func() {
			Description("Development hosts.")
			// トランスポート固有の URL, サポートされているスキーマは以下です：
			// 'http', 'https', 'grpc' and 'grpcs' とそれぞれのデフォルト
			// ポート： 80, 443, 8080, 8443
			URI("http://localhost:8000/calc")
			URI("grpc://localhost:8080")
		})

		Host("production", func() {
			Description("Production hosts.")
			// URL は {param} 記法を用いてパラメータ化できます。
			URI("https://{version}.goa.design/calc")
			URI("grpcs://{version}.goa.design")

			// Variable は URI の変数を記述します
			Variable("version", String, "API version", func() {
				// URL パラメータにはデフォルト値か Enum バリデーション（あるいはその両方）が必要です
				Default("v1")
			})
		})
	})
})

var _ = Service("calc", func() {
	// 説明
	Description("The calc service performs operations on numbers.")

	// メソッド（HTTPでいうところのエンドポイントに相当）
	// サービスで提供されるメソッドを複数定義可能
	Method("add", func() {
		// Payload はメソッドのペイロードを記述
		// ここでは、ペイロードは2つのフィールドからなるオブジェクト
		Payload(func() {
			// Field はフィールドインデックス、フィールド名、タイプ、および説明が与えられたオブジェクトを記述
			Attribute("a", Int, "Left operand", func() {
				Meta("rpc:tag", "1")
				// バリデーションの実施
				Minimum(1)
				Maximum(10)
			})
			Field(2, "b", Int, "Right operand")
			// Field DSL は gRPC のフィールドを記述するためのエイリアスで、下記と同等
			// -------------------------------------------
			// Attribute("b", Int, "Right operand", func(){
			// 	Meta("rpc:tag", "2")
			// })
			// -------------------------------------------
			// Required は必須となるフィールドの名前を列挙
			// RequiredAttributeを使用すると、Reguiredを使用しなくていい
			Required("a", "b")
		})

		// Result はメソッドの返り値を記述
		// ここでは、結果は単純に1つの整数値(Int)
		Result(Int)

		// HTTP は HTTP トランスポートへの対応を記述
		HTTP(func() {
			// サービスへのリクエストは HTTP GET リクエストで構成されてる
			// ペイロードはパスパラメータとしてエンコードされる
			GET("/add/{a}/{b}")
			// レスポンスは HTTP ステータス "200 OK" を使用する
			// 結果はレスポンスボディにエンコードする（デフォルト）
			Response(StatusOK)
		})

		// GRPC は gRPC トランスポートへの対応を記述
		GRPC(func() {
			// レスポンスは gRPC コード "OK" を使用
			// 結果はレスポンスメッセージにエンコードされる（デフォルト）
			Response(CodeOK)
		})
	})
	// サービスメソッドは Method 関数を用いて記述
	// STEP1
	// Payload (Methodへの入力)
	// Result (Methodの正常系の出力)
	// Error (Methodが返す可能性があるエラー)
	// STEP2
	// HTTP へのマッピング

	Method("divide", func() {
		// 入力
		Payload(func() {
			Attribute("a", Int, "Left operand", func() {
				Meta("rpc:tag", "1")
			})
			Attribute("b", Int, "Right operand", func() {
				Meta("rpc:tag", "2")
			})
			Required("a", "b")
		})
		// 出力
		Result(Int)

		// エラー
		Error("DivByZero")

		// HTTP トランスポートへの割り当て
		HTTP(func() {
			GET("/divide/{a}/{b}")                  // --> 入力
			Response(StatusOK)                      // --> 正常系
			Response("DivByZero", StatusBadRequest) // --> エラー
		})
		// GRPC トランスポートへの割当
		GRPC(func() {
			Response(CodeOK)                           // --> 正常系
			Response("DivByZero", CodeInvalidArgument) // --> エラー (INVALID_ARGUMENT=3)
		})
	})
})
