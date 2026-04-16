# go-hands-on

プログラミング初心者向けのGoハンズオン教材です。

## 構成

| ディレクトリ | 内容 |
|---|---|
| `01_hello_world` | Hello, World! を出力する |
| `02_loop_numbers` | ループで1〜10を表示する |
| `03_loop_function` | ループを関数にする |
| `04_even_odd` | 偶奇を判定する関数 |
| `05_even_numbers` | 1〜100の偶数を表示する |
| `06_fizzbuzz` | FizzBuzz |
| `07_api_hello_world` | Hello, World! を返すREST API |

## 実行方法

```bash
# 各ディレクトリで go run main.go を実行する
go run 01_hello_world/main.go

# APIサーバーの場合
go run 07_api_hello_world/main.go
# → curl http://localhost:8080/hello
```
