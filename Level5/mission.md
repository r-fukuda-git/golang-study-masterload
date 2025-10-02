
## 課題5：並行リンクチェッカー付きウェブスクレイパー

指定したURLのウェブページを取得し、そのページ内に含まれるすべてのリンク（`<a>`タグの`href`属性）を抽出します。そして、抽出したリンクに対して、課題4で作成した**並行ステータスチェッカー**のロジックを使い、リンク切れがないかをチェックするツールを作成します。

### 要件

1.  **入力:**

      * コマンドライン引数で、スクレイピングの起点となるURLを**1つ**だけ受け取ります。

2.  **リンクの抽出:**

      * 指定されたURLのHTMLを取得します。
      * 取得したHTMLの中から、すべての`<a>`タグを探し出し、その`href`属性の値をリストとして抽出してください。

3.  **リンクの正規化:**

      * 抽出したリンクには、`/about`のような相対パスや、`#section1`のようなページ内リンクが含まれていることがあります。
      * 絶対URL（例: `https://example.com/about`）に変換する必要があるものは変換し、HTTP/HTTPS以外のリンク（例: `mailto:test@example.com`）やページ内リンクはチェック対象から除外してください。

4.  **並行チェックと出力:**

      * 抽出・正規化したURLのリストに対して、課題4で実装した「**ゴルーチンで並行チェックし、チャネルで結果を集約する**」ロジックを再利用してください。
      * すべてのリンクのチェックが完了したら、リンクとそのステータスを一覧で表示してください。
          * 例:
            ```
            Checking links on https://example.com ...

            OK     (200 OK) https://example.com/about
            OK     (200 OK) https://www.iana.org/domains/example
            BROKEN (404 Not Found) https://example.com/broken-link
            ...
            ```

### ヒント

  * **HTMLの解析:** Goの標準ライブラリにはHTMLを解析するための`golang.org/x/net/html`パッケージがあります。しかし、これは少し複雑なので、より簡単に使えるサードパーティ製のライブラリ\*\*`go-query`\*\*を使ってみるのがお勧めです。
      * インストール: `go get github.com/PuerkitoBio/go-query`
      * `go-query`を使うと、jQueryのようにCSSセレクタ（例: `"a"`）を使ってHTML要素を簡単に見つけられます。
    <!-- end list -->
    ```go
    // 簡単な使用例
    doc, err := goquery.NewDocument(url)
    if err != nil { /* ... */ }
    doc.Find("a").Each(func(i int, s *goquery.Selection) {
        href, exists := s.Attr("href")
        if exists {
            // hrefが見つかった時の処理
        }
    })
    ```
  * **URLの正規化:** `net/url`パッケージが非常に役立ちます。
      * `url.Parse(urlString)`でURL文字列を構造体にパースできます。
      * 起点となるURLの構造体に対して`base.ResolveReference(relative)`メソッドを使うと、相対URLを簡単に絶対URLに変換できます。

この課題は、\*\*「外部ライブラリの導入」「HTMLの解析」「URLの処理」\*\*という新しい要素が加わり、これまでの知識を組み合わせる応用問題です。これまでで最もチャレンジングですが、これができれば非常に実用的なツールを一人で作れるようになります。頑張ってください！