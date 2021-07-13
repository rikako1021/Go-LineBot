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

プログラムに必要なパッケージをインポートする。

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

```http.HandleFunc("/", helloHandler)```ではルートドメインにアクセスすると```helloHandler```が実行されるように登録されている。

<br>

### HTTPサーバの起動
```http.ListenAndServe```でHTTPサーバを起動している

通常はlocalhost:8000などのローカルサーバで良いが、今回はHerokuでデプロイするステップがあり、Herokuはデプロイ時勝手にポート番号を決定するのでこちらで指定することができない。

(localhostでポート番号を指定したまま実行するとデプロイエラーが発生する。)

そのため、決定されたポート番号を```Getenv("PORT")```で取得できるようにしている。

<br>

### 立てたサーバのURLをコンソールに出力
```
fmt.Println("https://fathomless-depths-28419.herokuapp.com/ で起動中")
```
今回はHerokuでデプロイするため、Herokuでのアプリケーション作成時に生成されるURL([README](/README.md)4.参照)を出力し、アクセスできるようにしている。

プログラムを実行してサーバを立てると、```Println```の中身がコンソールに出力される。

<br>

## helloHandler関数
```
func helloHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	fmt.Fprint(w, msg)
}
```
### 関数宣言
Golangでは、関数は　
```
func 関数名([引数]) [戻り値の型] {
    処理
}
```
と定義する。helloHandlerでは戻り値がない。

### 変数宣言
```msg := "Hello World"```の部分。


変数は、
```変数名 := 初期値```で宣言できる。

型推論ができ、右辺から判断して自動で型付けされる。

<br>

## lineHandler関数
### botの初期化
```
bot, err := linebot.New(
		"(secret)",
		"(access token)",
	)
	if err != nil {
		log.Fatal(err)
	}
```
Golangでは複数の変数(この場合はbot,err)を宣言できる。

```secret```, ```access token```にそれぞれLINEmessagingAPIの認証情報を追加する。

(APIの登録・接続の詳細は[README](/README.md)3.参照)

<br>

### イベントの取得
```
events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
```
この部分で、LINEトーク上でBOTアカウントが受信したメッセージの内容をチェックする。

<br>

### イベントの種類によって個別に処理

<li>受信したメッセージがテキスト形式の場合

```
case *linebot.TextMessage:
				replyMessage := message.Text
				_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				if err != nil {
					log.Print(err)
				}
```
受信したテキストの内容を、変数```replyMessage```に入れ、```NewTestMessage```メソッドの引数として実行され、メッセージとして送信される。

(ここではテキストメッセージを受信すると王蟲返しする。)

<li>受信したメッセージが位置情報の場合</li>

