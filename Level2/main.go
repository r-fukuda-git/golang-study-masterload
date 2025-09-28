package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func search(s *bufio.Scanner, pattern string, nFlag bool) {
	lineNumber := 1
	for s.Scan() {
		line := s.Text()
		if strings.Contains(line, pattern) {
			if nFlag {
				fmt.Printf("%d:%s\n", lineNumber, line)
			} else {
				fmt.Println(line)
			}
		}
		lineNumber++
	}
}

func main() {
	nFlag := flag.Bool("n", false, "行番号を表示します") //-n が指定された時に *nFlag の中身が true になる
	flag.Parse()
	/*
		for _, args := range flag.Args() {
			fmt.Println(args)
		}
	*/

	args := flag.Args()
	if len(args) == 0 { // 引数を指定していない場合panicになるため条件分岐の設定
		fmt.Fprintln(os.Stderr, "エラー: 検索パターンを指定してください。")
		os.Exit(1)
	}
	pattern := args[0]
	files := args[1:]

	if len(files) == 0 {
		scan := bufio.NewScanner(os.Stdin)
		search(scan, pattern, *nFlag) //search関数によって重複処理をまとめ
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			defer f.Close()

			scan := bufio.NewScanner(f)
			search(scan, pattern, *nFlag)
		}
	}
	//fmt.Println(pattern) ここの出力は不要
}
