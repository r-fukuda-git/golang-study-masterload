package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := len(os.Args)
	fmt.Println(args)

	for _, args := range os.Args[1:] { //引数1からにすることでプログラム自身の名前を除外とする
		fmt.Printf("ファイル名: %s\n", args)

		// open file
		f, err := os.Open(args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		defer f.Close()

		// read file
		s := bufio.NewScanner(f)
		for s.Scan() {
			fmt.Println(s.Text())
		}

		// output
		//fmt.Println(f) //ポインタを表示させてしまったため、文字化けが出力された

	}
}
