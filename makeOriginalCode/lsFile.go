package main

import (
	"fmt"
	"os"
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
	nilProcess(err) // errという変数を持ってくることでこの関数はエラー型と認識する
	fmt.Println(dir)

	fileName, err := os.ReadDir(dir)
	nilProcess(err)
	fmt.Println(fileName)

	fileSize, err := os.Stat(dir)
	nilProcess(err)
	fmt.Println(fileSize)

	// 今後やりたいこと
	// ディレクトリにあるファイルの取得はできているからそれに対して個別にするようにしたい
}

//	ヒント
//	実装したい機能,使うパッケージ・関数,取得できる情報（戻り値）
//	ファイル名一覧,os.ReadDir,DirEntry.Name()
//	ファイルサイズ,os.Stat -> FileInfo,FileInfo.Size()
//	パーミッション,os.Stat -> FileInfo,FileInfo.Mode()
//	更新日時,os.Stat -> FileInfo,FileInfo.ModTime()
//	隠しファイルの判定,文字列操作,"strings.HasPrefix(name, ""."")"
