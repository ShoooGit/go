package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gihyodocker/todoapi"
)

func main() {

	// ① 必要な環境変数を格納した構造体を作成
	env, err := todoapi.CreateEnv()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	// ② MySQL Masterへ接続するための構造体を作成
	masterDB, err := todoapi.CreateDbMap(env.MasterURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s is invalid database", env.SlaveURL)
		return
	}

	mux := http.NewServeMux()

	// ④ ヘルスチェック用APIのハンドラを作成
	hc := func(w http.ResponseWriter, r *http.Request) {
		log.Println("[GET] /hc")
		w.Write([]byte("OK"))
	}

	// ⑤ TODO操作APIのハンドラを作成
	todoHandler := todoapi.NewTodoHandler(masterDB, slaveDB)

	// ⑥ ハンドラをAPIエンドポイントとして登録
	mux.Handle("/todo", todoHandler)
	mux.HandleFunc("/hc", hc)

	// ⑦ サーバのポートやハンドラを設定し、Listenを開始
	s := http.Server{
		Addr:    env.Bind,
		Handler: mux,
	}
	log.Printf("Listen HTTP Server")
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
