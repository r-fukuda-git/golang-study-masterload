# Go言語における安全なパスワードハッシュ化 (bcrypt)

## 1\. なぜ bcrypt を使うのか？

従来のハッシュ関数（SHA-1やMD5など）は、パスワードの保存には適していません。これらは計算が**高速すぎる**ため、データベースが漏洩した場合、攻撃者による総当たり攻撃（ブルートフォースアタック）やレインボーテーブル攻撃を容易にしてしまいます。

`bcrypt` は、パスワードのハッシュ化専用に設計されたアルゴリズムであり、以下の特徴によって高い安全性を持ちます。

  * **意図的に低速（高コスト）**:
    計算に時間がかかるように設計されており、総当たり攻撃の効率を大幅に低下させます。
  * **自動ソルト（Salt）**:
    ハッシュ化の際に「ソルト」と呼ばれるランダムなデータを自動で付与します。これにより、同じパスワードでも毎回異なるハッシュ値が生成され、レインボーテーブル攻撃を無効化します。

## 2\. 依存関係のインポート

`bcrypt` を使用するために、Go言語の拡張暗号化パッケージをインポートします。

```go
import "golang.org/x/crypto/bcrypt"
```

-----

## 3\. 実装方法

`bcrypt` の実装は、「ハッシュ化（保存）」と「比較（認証）」の2つの関数が中心となります。

### ① パスワードのハッシュ化 (新規登録時)

ユーザーが入力した生のパスワードをハッシュ化し、データベースに保存します。

  * **使用する関数**: `bcrypt.GenerateFromPassword(password []byte, cost int) ([]byte, error)`

**実装例:**
`cost`（計算コスト）は、`bcrypt.DefaultCost`（デフォルト値、通常は10）を指定するのが一般的です。

```go
import (
    "golang.org/x/crypto/bcrypt"
    "log"
)

// ユーザー登録時などに呼び出す関数
func HashPassword(password string) (string, error) {
    // 生のパスワードをバイト配列に変換
    passwordBytes := []byte(password)

    // ハッシュ化を実行
    // DefaultCost を使用すると、適切な計算コストが自動で設定される
    hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
    if err != nil {
        log.Println(err)
        return "", err
    }

    // ハッシュ化された文字列を返す
    // (例: $2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy)
    return string(hashedPassword), nil
}
```

### ② パスワードの比較 (ログイン時)

ユーザーが入力したパスワードが、データベースに保存されているハッシュ値と一致するかを検証します。

  * **使用する関数**: `bcrypt.CompareHashAndPassword(hashedPassword []byte, password []byte) error`

**重要**:
認証時は、入力されたパスワードを**再度ハッシュ化して文字列比較をしてはいけません**（`bcrypt` は毎回ソルトが変わるため）。必ず `CompareHashAndPassword` を使用してください。

**実装例:**

```go
import (
    "golang.org/x/crypto/bcrypt"
    "log"
)

// ログイン認証時などに呼び出す関数
// dbHash: データベースに保存されているハッシュ化済みパスワード
// plainPassword: ユーザーがログインフォームに入力した生のパスワード
func CheckPasswordHash(dbHash string, plainPassword string) bool {
    // データベースのハッシュと、入力されたパスワードをバイト配列に変換
    dbHashBytes := []byte(dbHash)
    plainPasswordBytes := []byte(plainPassword)

    // ハッシュとパスワードを比較
    err := bcrypt.CompareHashAndPassword(dbHashBytes, plainPasswordBytes)

    // err が nil の場合は認証成功
    if err != nil {
        log.Println("Password mismatch:", err) // bcrypt.ErrMismatchedHashAndPassword が返る
        return false
    }

    return true
}
```

## 4\. まとめ

  * パスワード保存には **SHA-1** ではなく **bcrypt** を使用する。
  * ハッシュ化には `bcrypt.GenerateFromPassword` を使う。
  * 認証（比較）には `bcrypt.CompareHashAndPassword` を使う。
  * `bcrypt` を使えば、ソルトの生成や管理を意識する必要がなく、安全なパスワード管理を簡単に実装できます。