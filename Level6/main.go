package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

func init() {
	LoggingSettings(logfile)
}

// ログはグローバル変数で表記
var logfile = "error.log"

func LoggingSettings(logfilePath string) {

	file, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}

	multiLogFile := io.MultiWriter(os.Stdout, file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

var (
	//todos[1]で設定Todoの構造化したアイテムを取得できる
	//mapはキーとバリューのペアを格納するためのデータ構造
	todos  = make(map[int]Todo)
	nextID = 1
	//Mutexを使用することで競合状態を回避
	mu sync.Mutex
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func index(w http.ResponseWriter, r *http.Request) {
	//switch式でcaseごとにGETかPUTかでリクエストを分けていく方法を採用
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(todos); err != nil {
			//todosというデータをJSON形式に翻訳し、wという変数でレスポンスをするように条件式
			log.Printf("JSONエンコードエラー: %v", err)
			http.Error(w, "サーバエラー", http.StatusInternalServerError)
		}
	case http.MethodPost:
		var newTodo Todo
		if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
			//r.Bodyで届いたJSONデータをnewTodoに書き写す処理
			http.Error(w, "不正なリクエストボディです", http.StatusBadRequest)
			return
		}

		mu.Lock() //ロック中
		newTodo.ID = nextID
		todos[nextID] = newTodo
		nextID++
		mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTodo)

	default:
		http.Error(w, "サポートされていないメソッドです", http.StatusMethodNotAllowed)
	}
}

func StartMainServer() error {
	http.HandleFunc("/todos", index) //HandleFuncでパス指定し、サーバ起動
	fmt.Println("サーバーをポート8080で起動します...")
	return http.ListenAndServe(":8080", nil)
}

func main() {
	todos[1] = Todo{ID: 1, Title: "本日のタスク"}
	nextID = 2
	StartMainServer()
}
