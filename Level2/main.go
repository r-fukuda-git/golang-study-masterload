package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

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
		lineNumber := 1
		for scan.Scan() {
			line := scan.Text()
			if strings.Contains(line, pattern) {
				if *nFlag {
					fmt.Printf("%d:%s\n", lineNumber, line)
				} else {
					fmt.Println(line)
				}
				lineNumber++
			}
		}
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			defer f.Close()

			s := bufio.NewScanner(f)
			lineNumber := 1
			for s.Scan() {
				//fmt.Println(s.Text()) ここが全て表示となっていたため条件設定
				line := s.Text()
				if strings.Contains(line, pattern) { // パターンが含まれていたら、この{}の中が実行される
					if *nFlag {
						fmt.Printf("%d:%s\n", lineNumber, line)
					} else {
						fmt.Println(line)
					}
					lineNumber++
				}
			}
		}
	}
	//fmt.Println(pattern)
}
