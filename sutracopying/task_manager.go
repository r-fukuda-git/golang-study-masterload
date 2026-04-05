package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Stringer interface {
	ToString() string
}

type Task struct {
	ID        int
	Title     string
	Completed bool
	Duration  int
}

func (t *Task) ToString() string {
	status := "未完了"
	if t.Completed {
		status = "完了"
	}
	return fmt.Sprintf("[%d] %s (%s)", t.ID, t.Title, status)
}

type TaskList struct {
	tasks []Task
	mu    sync.Mutex
}

func (l *TaskList) AddTask(title string, completed bool, duration int) error {
	if title == "" {
		return errors.New("タイトルが空なのよ")
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	id := len(l.tasks) + 1
	l.tasks = append(l.tasks, Task{ID: id, Title: title, Completed: completed, Duration: duration})
	return nil
}

func saveToCloud(t Task, ch chan<- string) {
	time.Sleep(time.Duration(t.Duration) * time.Second)
	ch <- fmt.Sprintf("Task %d %s (所要時間: %d秒)のタスクをクラウドに同期しました。", t.ID, t.Title, t.Duration)
}

func main() {
	myTasks := TaskList{
		tasks: []Task{
			{ID: 1, Title: "Goの基礎を学ぶの設計図を書くよ", Completed: true, Duration: 3},
			{ID: 2, Title: "テストを実施します", Completed: false, Duration: 2},
			{ID: 3, Title: "テストをしているよ", Completed: true, Duration: 1},
			{ID: 4, Title: "テスト完了したらデプロイするよ", Completed: true, Duration: 6},
			{ID: 5, Title: "デプロイ中・・", Completed: true, Duration: 10},
		},
	}

	if err := myTasks.AddTask("写経を完成する", false, 4); err != nil {
		fmt.Println("エラー:", err)
	}

	if err := myTasks.AddTask("今日はこのコードを改造させる", true, 10); err != nil {
		fmt.Println("エラー:", err)
	}

	for _, t := range myTasks.tasks {
		fmt.Println(t.ToString())
	}

	stats := map[string]int{
		"total": len(myTasks.tasks),
	}
	fmt.Printf("\n統計: 合計 %d 件 \n", stats["total"])

	fmt.Println("バックグラウンド処理を開始")
	startTime := time.Now()
	msgChan := make(chan string)
	completeCount := 0

	for _, t := range myTasks.tasks {
		if t.Completed {
			completeCount++
			go saveToCloud(t, msgChan)
		}
	}
	fmt.Printf("\n%d 件の同期処理を実施する\n", completeCount)
	fmt.Printf("\n%d件の完了済みタスクを同期中\n", completeCount)

	for i := 0; i < completeCount; i++ {
		msg := <-msgChan
		fmt.Println("通知:", msg)
	}

	finishTime := time.Since(startTime)
	fmt.Printf("全て完了（合計時間:%v）\n", finishTime)

}
