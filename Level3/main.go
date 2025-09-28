package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Counts struct {
	lines int
	words int
	bytes int
}

func count(s *bufio.Scanner) Counts {
	c := Counts{} // lines, words, bytes すべてが0で初期化される、カウンターをゼロ値に明示的に表示

	for s.Scan() {
		line := s.Text()                     //読み込みに全ての行を取得
		c.lines++                            //行数を増やす
		c.words += len(strings.Fields(line)) //単語数を増やす＝その行から単語のリスト（配列）を作る
		c.bytes += len(line) + 1             // バイト数を数えて足す
	}
	return c
}

/* ここではscanの関数を用いて中で表示させてしまったがゆえに失敗した
func scan(s *bufio.Scanner, pattern string, lFlag bool, wFlag bool, cFlag bool) {
	c := Counts{} // lines, words, bytes すべてが0で初期化される、カウンターをゼロ値に明示的に表示

	for s.Scan() {
		line := s.Text()                     //読み込みに全ての行を取得
		c.lines++                            //行数を増やす
		c.words += len(strings.Fields(line)) //単語数を増やす＝その行から単語のリスト（配列）を作る
		c.bytes += len(line) + 1             // バイト数を数えて足す
	}

	fmt.Printf("  %d  %d  %d\n", c.lines, c.words, c.bytes)
}
*/

func printCounts(counts Counts, filename string, lFlag, wFlag, cFlag bool) { //counts Countsは変数countsに対して構造体であるCountsを定義している
	if !lFlag && !wFlag && !cFlag {
		lFlag, wFlag, cFlag = true, true, true
	}
	if lFlag {
		fmt.Printf("%8d", counts.lines)
	}
	if wFlag {
		fmt.Printf("%8d", counts.words)
	}
	if cFlag {
		fmt.Printf("%8d", counts.bytes)
	}
	fmt.Println(" " + filename)
	//"%8d"
	//%: 「ここに引数の値を埋め込んでください」という合図
	//d: decimal（10進数の整数）の略です。「ここに整数を表示してください」という意味
	//8: 「全体で8文字分の幅を確保してください」という意味です。もし表示する数値が8文字より少なければ、左側にスペースを詰めて右揃えにしてくれます。
}

func main() {
	lFlag := flag.Bool("l", false, "行数を表示します")
	wFlag := flag.Bool("w", false, "単語数を表示します")
	cFlag := flag.Bool("c", false, "バイト数を表示します")

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 { // 引数を指定していない場合panicになるため条件分岐の設定
		s := bufio.NewScanner(os.Stdin)
		counts := count(s)
		printCounts(counts, "", *lFlag, *wFlag, *cFlag) // ファイル名はないので空文字列
	} else {
		tc := Counts{}
		for _, file := range args {
			f, err := os.Open(file)
			if err != nil {
				fmt.Println(err)
				return
			}
			//fc := Counts{} // このファイル用のカウンターを用意
			s := bufio.NewScanner(f)
			fc := count(s)
			defer f.Close() //deferを使うと、「この関数が終了する直前に、この処理を実行してください」という予約ができる
			printCounts(fc, file, *lFlag, *wFlag, *cFlag)

			tc.lines += fc.lines
			tc.words += fc.words
			tc.bytes += fc.bytes
		}

		if len(args) > 1 {
			printCounts(tc, "total", *lFlag, *wFlag, *cFlag) //ファイル名はトータルと文字列で表示
		}
	}
}
