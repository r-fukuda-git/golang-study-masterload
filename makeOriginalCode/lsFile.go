package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"syscall"
	"text/tabwriter"
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
	fmt.Println(dir) // ここでディレクトリパスを取得している

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0) // 最終的な出力するための変数で定義

	// list := flag.Bool("-l", false, "詳細に表示します") // 引数を確認するため
	list := flag.Bool("l", false, "詳細に表示") // ハイフンは含めないのが一般的
	// flag.Bool という関数の「戻り値の型」が、最初から *bool（ポインタ型）として設計されている
	flag.Parse()

	if *list {
		//このlistがなぜポインタ付きなのかが謎？？
		// ① 「値（あたい）」そのものを入れる変数
		// ② 「ポインタ（住所）」を入れる変数
		// bool 型: true か false かを判定できる。
		// *bool 型（ポインタ）: 0x14000... という 「数字（住所）」。
		// fileName, err := os.ReadDir(dir) // ←この状態だとカレントディレクトリの情報を取得してしまっているため、ループで一つずつのファイルを処理する必要がある
		fileName, err := os.ReadDir(dir)
		for _, fileNames := range fileName {
			// for ..rangeについて2つの値を返してきます。
			// 第1戻り値: 現在の要素の インデックス（添え字）。0, 1, 2... という数字。
			// 第2戻り値: 現在の要素の 中身（値）。今回で言えば、ファイル情報そのもの。
			// go wayのため、変数したルールは必ず使用しないといけないので、_,として変数をブランクにおいている
			nilProcess(err)
			// fmt.Println(fileNames)

			// ここで1つのファイル情報を取得しているため
			// 詳細を出力させる
			// fileSize, err := os.Stat(dir)
			// 上記の形だと現在のフォルダそのものを持ってきている状態なので、フォルダ+ファイルを組み合わせないといけない
			fullPath := filepath.Join(dir, fileNames.Name())
			fileInfo, err := os.Stat(fullPath)
			nilProcess(err)
			// fmt.Println(fileInfo)

			//	list := flag.Bool("-l", false, "詳細に表示します")
			//	flag.Parse()
			// これはループの中で呼びだす必要はない

			// 詳細に表示できたので、ここから変えていく
			// userId, err := user.LookupId(dir)
			// nilProcess(err)
			// fmt.Println(userId)
			// 上記ではUIDを取り出せずにエラーになってしまう
			stat := fileInfo.Sys().(*syscall.Stat_t)
			// fileInfo.Size(): 全OS共通（表のメニュー）であり、fileInfo.Sys(): 各OS独自の生データ（裏メニューへの入り口）になる
			// 基本的にGolangは他OSでも動くようにできているけど、UIDやGIDはUNIXのみの機能のため、Sys()という設定が必要になってくる
			// fileInfo.Sys()の中身はなんでも入るany型となっているので、Mac/Linuxの統計データ（*syscall.Stat_t）という中身を断定しているアサーション設定が必要になってく

			uid := fmt.Sprint(stat.Uid)
			// uid := fmt.Printf(stat.Uid)
			// fmt.Printfについてf画面に出力するという副作用がメインの関数です。戻り値として「文字列」を返してくれません。
			// fmt.Sprint渡されたデータを文字列に変えて 戻り値 として返すための関数なので、今回戻り値が必要なSprintを使用する

			u, err := user.LookupId(uid)
			// fmt.Println(u)

			modTime := fileInfo.ModTime().Format("Jan _2 15:04") // タイムスタンプ用の設定
			fmt.Fprintln(w, fileInfo.Mode(), fileInfo.Size(), u.Username, modTime, fileNames.Name())
		}
		w.Flush()

		// nilProcess(err) //上記にループ文に加える
		// fmt.Println(fileName) //上記にループ文に加える

		// fileSize, err := os.Stat(dir) //上記にループ文に加える
		// nilProcess(err) //上記にループ文に加える
		// fmt.Println(fileSize) //上記にループ文に加える
	} else {
		fileName, err := os.ReadDir(dir)
		for _, fileNames := range fileName {
			fmt.Println(fileNames)
			nilProcess(err)
		}

	}

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
// filepath.Joinは個別のファイルを作成するパッケージ
// syscallパッケージ
// tabwriter.NewWriter
/// 引数の位置,パラメータ名,設定値,役割（インフラエンジニア的解釈）
/// 第1,output,os.Stdout,出力先。画面に出すなら標準出力。
/// 第2,minwidth,0,最小カラム幅。通常は0でOK（内容に合わせるため）。
/// 第3,tabwidth,8,タブの幅。タブ文字1つをスペース何個分と計算するか。
/// 第4,padding,2,カラム間の余白。列と列の間に最低何個スペースを入れるか。
/// 第5,padchar,' ',穴埋め文字。隙間を何で埋めるか（通常は半角スペース）。
/// 第6,flags,0,特殊設定。デバッグ用などのフラグ（通常は0でOK）。
