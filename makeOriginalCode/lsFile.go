package main

import (
	"fmt"
	"os"
	"os/user"
)

// mainパッケージにて設定し、bashのls機能を再現しようとしている
// エラーのnil処理を関数として実行すれば簡単！と思ったけど違うんやな
// 変数errが見えていないことが原因でエラーがでている
// 引数がないため、lsFile関数では何を見ればいいかわからずにいる
// func nilProcess() {
// なぜ引数が大事なの？ → この関数がよぶときは必ずerror型のデータをもってくることというルールづけになるから
func nilProcess(err error) { // err errorと引数をつけることで外からデータを受け取れる準備ができる
	if err != nil {
		fmt.Println(err)
	}
}

// 本処理実行
func lsFile() {
	dir, err := os.Getwd()
	nilProcess(err)  // errという変数を持ってくることでこの関数はエラー型と認識する
	fmt.Println(dir) // これはフルパスを取得している

	// fileName, err := os.ReadDir(dir) // ←この状態だとカレントディレクトリの情報を取得してしまっているため、ループで一つずつのファイルを処理する必要がある
	fileName, err := os.ReadDir(dir)
	for _, fileNames := range fileName {
		nilProcess(err)
		fmt.Println(fileNames)

		// ここで1つのファイル情報を取得しているため
		// 詳細を出力させる
		fileSize, err := os.Stat(dir)
		nilProcess(err)
		fmt.Println(fileSize)

		// 詳細に表示できたので、ここから変えていく
		userId, err := user.LookupId(dir)
		nilProcess(err)
		fmt.Println(userId)
	}

	// nilProcess(err) //上記にループ文に加える
	// fmt.Println(fileName) //上記にループ文に加える

	// fileSize, err := os.Stat(dir) //上記にループ文に加える
	// nilProcess(err) //上記にループ文に加える
	// fmt.Println(fileSize) //上記にループ文に加える
}

//	ヒント
//	実装したい機能,使うパッケージ・関数,取得できる情報（戻り値）
//	ファイル名一覧,os.ReadDir,DirEntry.Name()
//	ファイルサイズ,os.Stat -> FileInfo,FileInfo.Size()
//	パーミッション,os.Stat -> FileInfo,FileInfo.Mode()
//	更新日時,os.Stat -> FileInfo,FileInfo.ModTime()
//	隠しファイルの判定,文字列操作,"strings.HasPrefix(name, ""."")"
// os.ReadDir(path): 指定したフォルダの中にある「ファイル名のリスト」を取得します。
// os.Stat(path): 特定のファイルについて、もっと深く、詳細な情報を取得します。
// user.LookupId(uid): 「UIDが 0 の人の名前を教えて！」とOSに問い合わせるパッケージです。
// Time.Format(layout): 日付データを、指定した見た目の文字列（string）に変換します。
// オプション機能を実装するために使用する標準パッケージは flag です。
