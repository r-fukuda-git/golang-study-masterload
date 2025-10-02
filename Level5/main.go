package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func checkStatus(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("BROKEN (%s) %s", err, url)
		return
	}
	defer resp.Body.Close() //deferで閉じないといけない

	ch <- fmt.Sprintf("OK     (%s) %s", resp.Status, url)
}

func main() {
	// 1. 引数が1つだけかチェックする
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "エラー：URLを1つ選択してください")
		os.Exit(1)
	}
	startUrl := os.Args[1]

	// 2. 抽出したリンクを保存するための、空のスライスを用意する
	links := []string{}

	// 起点となるURLをパース
	// 起点URLのパースは、ループの前に一度だけ行う
	baseUrl, err := url.Parse(startUrl)
	if err != nil {
		fmt.Println(err)
	}

	// 3. go-queryを使ってページを取得し、リンクを抽出する
	doc, err := goquery.NewDocument(startUrl)
	if err != nil {
		fmt.Println(err)
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href") // href (string): href属性の値そのもの
		if exists {                    //exists (bool): href属性が存在したかどうか
			// ★ここに抽出したリンクをlinksスライスに追加する処理を書く
			//links = append(links, href)

			// 抽出したhrefをパース
			hrefUrl, err := url.Parse(href)
			if err != nil {
				fmt.Println(err)
			}

			// ResolveReferenceで相対URLを絶対URLに変換
			absoluteUrl := baseUrl.ResolveReference(hrefUrl).String()

			// HTTP/HTTPSのリンクだけをスライスに追加する
			if strings.HasPrefix(absoluteUrl, "http") {
				links = append(links, absoluteUrl)
			}
		}
	})

	resultsChannel := make(chan string) //string型のデータを送受信できるチャネル

	// この時点で、linksスライスに全リンクが入っているはず
	fmt.Println("抽出したリンク:")

	for _, link := range links {
		go checkStatus(link, resultsChannel)
	}

	for i := 0; i < len(links); i++ { //1個以上リンクがあればそれを出力する
		fmt.Println(<-resultsChannel)
	}
}
