package main

import (
	"fmt"
	"net/http"
	"os"
)

func checkStatus(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%s: FAILD(%s)", url, err)
		return
	}
	defer resp.Body.Close()

	ch <- fmt.Sprintf("%s: %s", url, resp.Status)
}

func main() {
	resultsChannel := make(chan string)
	for _, url := range os.Args[1:] {
		go checkStatus(url, resultsChannel)
	}

	//処理結果
	for i := 0; i < len(os.Args[1:]); i++ {
		result := <-resultsChannel
		fmt.Println(result)
	}
}
