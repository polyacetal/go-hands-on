package main

import (
	"fmt"
	"net/http"
)

// corsMiddleware はすべてのレスポンスにCORSヘッダーを付与するミドルウェアです。
// ローカルのHTMLファイル（file://）からAPIを呼び出せるようにするために必要です。
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}

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
	// corsMiddleware でラップすることで、ブラウザからの直接アクセスが可能になります
	http.HandleFunc("/hello", corsMiddleware(helloHandler))

	fmt.Println("サーバーを起動しました: http://localhost:8888")
	fmt.Println("試してみよう: curl http://localhost:8888/hello")
	fmt.Println("フロントエンド: 07_api_hello_world/index.html をブラウザで開いてください")

	// ポート8080でHTTPサーバーを起動する（Ctrl+C で停止）
	http.ListenAndServe(":8888", nil)
}
