package main

import (
	"fmt"
	"math/rand"
)

// 一番初めはmain関数から始めないといけない
func main() {
	mainFirst()
	mainSecond()
}

func mainFirst() {
	// まずは出力できるか確認
	fmt.Println("my favorite number is", rand.Intn(100))

	// アンチパターンとして下記をしてしまった
	// number := fmt.Println("my favorite number is", rand.Intn(100))
	// これの何がいけないのかというと
	// ・fmt.Printlnは表示する内容を返す関数ではない
	// ・何を変数として置きたいのかを確認する必要がある。→ 例えば文字列を変数として置きたいのか？数字を変数として置きたいのか？

	// まずは文字列を変数としておく場合
	// string := "my favorite number is"
	// 上記のようにstringという型の名前を置いてしまうと、文字列型として今後使用できなくなるから異なる変数に置き換えるべき
	msg := "my favorite number is"
	fmt.Println(msg, rand.Intn(100))

	// 数値を変数としておく場合
	number := rand.Intn(100)
	fmt.Println(msg, number)

	// Go言語には戻り値を厳格に記載する必要がある
	// Printlnは成功したデータ＝バイト値、失敗した情報をエラーとして2つの値が返ってくる
	// Bashの概念ではあまり、ないこと基本標準出力で結果が画面に出てくるけど、戻り値に近い概念の終了ステータスと認識すればいいか
	// strings, err := "my favorite number is"
	// 上記のもミス。ただの変数だから戻り値が返ってくるはずがない

	strings, err := fmt.Println(msg, number)
	if err != nil {
		fmt.Println(err)
	}
	// 下記は合計のバイト数になるのはなんで？上記にてstringsを変数と置いているので、いつも通りお気に入りの数字はという形になるはずなのに
	// Go言語は画面への出力は行うのと別で書き込んだサイズも返すという明確な役割がある
	// fmt.Printlnは画面に表示することが役割
	// つまり、画面に表示させるのと同時に、最後の報告としてバイト数を表示させているだけ
	fmt.Println(strings)

	// 続いて最後の報告結果としてバイト数を表示させているが、これをif文で分岐させたい
	ok := "バイト数は25となります。"
	ng := "バイト数は25以外になりました"
	if strings == 25 {
		fmt.Println(ok)
		mainFirst() // 無限ループの恐れがあるけど、1桁の数字がでると24バイト数になるのでループから抜ける
	} else {
		fmt.Println(ng)
	}

	// 簡易付きif文について
	// 関数の実行結果をその場で判定し、変数の寿命を最小限にする形
	// なんで簡易付きif文が存在するのかについて、使い終わった変数を即座に破棄することでメモリの安全性を高めるため
	if ok1, err := fmt.Println(strings); err == nil && ok1 > 20 {
		fmt.Println(ok)
	} else if err != nil {
		// ネストが深くなるため、elseでエラー文を記述するのは可読性が低くなるため避けるべき記述
		// 上記のようなelse ifエラー文を記述するのは良い
		fmt.Println(err)
	}
}

func mainSecond() {
	fmt.Println("次の処理へ進める")

	// for文でループさせる
	for i := 1; i < 5; i++ {
		fmt.Printf("%d回目のループです\n", i)

		// 今回は変数をif文中に入れておき、iの数が3に来たら変数を出力させ、ループを止める設定にした
		if three := "今回ばかりは許さんぞおおお"; i == 3 {
			fmt.Println(three)
			break
		}
	}
}
