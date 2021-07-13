<br>

※こちらはシステムのファイル、関数などについて説明した文書です。

※開発環境構築・プログラムの実行・デプロイ方法等については[README](/README.md)をご覧下さい。

<br>

# システムの全体構造

<li>main.go：</li>
　　プログラムを記述するメインファイル
<li>go.mod：</li>
　　Golangの外部パッケージ管理システムであるGoModulesを使うためのファイル
<li>go.sum<自動生成>：</li>
　　パッケージの依存関係を管理するためのファイル(npmでいうpackage.lock.jsonのようなもの)
<li>Procfile：</li>
　　Herokuでのデプロイに用いるWebdynoを管理するファイル
<li>README.md：</li>
　　アプリケーションの概要・環境構築・実行方法等を記載した文書
<li>aboutProg.md：</li>
　　システムの全体像・プログラムのメソッド等を説明した文書


<br>
<br>

# main.go

## import部分

プログラムに必要なパッケージをインポートします。

<li>"encoding/json"</li>
<li>"fmt"</li>
<li>"io/ioutil"</li>
<li>"log"</li>
<li>"net/http"</li>
<li>"os"</li>
<li>"strconv"</li>
<li>"unicode/utf8"</li>
<li>"github.com/line/line-bot-sdk-go/linebot"</li>

<br>

## main関数

```
func main() {
	//ハンドラ登録
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/callback", lineHandler)

	fmt.Println("https://fathomless-depths-28419.herokuapp.com/ で起動中")

	//サーバ起動
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
```
### ハンドラ登録
指定したURLにアクセスすると実行される関数を登録する

```http.HandleFunc("/", helloHandler)```ではルートドメインにアクセスすると```helloHandler```が実行されるように登録されている

### HTTPサーバの起動
