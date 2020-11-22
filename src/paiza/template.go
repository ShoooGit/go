package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 標準入力を受け取る準備
	stdin := bufio.NewScanner(os.Stdin)

	// 標準入力の受け取り
	stdin.Scan()

	// 入力値を半角スペースで区切る
	texts := strings.Fields(stdin.Text())

	// 入力値を数値に変換
	sample, _ := strconv.Atoi(texts[0])
	fmt.Println(sample)
	
	// 何かしらのループ処理
	for stdin.Scan() {
		fmt.Println(stdin.Text())
	}
}
