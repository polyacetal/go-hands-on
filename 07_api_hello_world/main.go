package main

import (
	"fmt"
	"net/http"
)

// helloHandler は /hello にリクエストが来たときに呼ばれる関数です。
// w: クライアントへのレスポンスを書き込む
// r: クライアントから受け取ったリクエストの情報
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// レスポンスのContent-TypeをJSONに設定する（ボディを書く前に行う）
	w.Header().Set("Content-Type", "application/json")

	// JSONをレスポンスボディとして送信する
	// 動的なデータがある場合は encoding/json パッケージを使う
	fmt.Fprintf(w, `{"message": "Hello, World!"}`)
}

func main() {
	// /hello へのリクエストを helloHandler で処理するよう登録する
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("サーバーを起動しました: http://localhost:8080")
	fmt.Println("試してみよう: curl http://localhost:8080/hello")

	// ポート8080でHTTPサーバーを起動する（Ctrl+C で停止）
	http.ListenAndServe(":8080", nil)
}
