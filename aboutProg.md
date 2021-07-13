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

heroku公式ドキュメントに記述方法の説明があるのでそれを参考に作成する。([README](/README.md)末尾<b>Links</b>参照)
<li>README.md：</li>
　　アプリケーションの概要・環境構築・実行方法等を記載した文書
<li>aboutProg.md：</li>
　　システムの全体像・プログラムのメソッド等を説明した文書


<br>
<br>

***


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

***

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

***

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

***


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

<br>

<li>受信したメッセージが位置情報の場合</li>

```
case *linebot.LocationMessage:
				sendRestoInfo(bot, event)
```
位置情報が送信された場合は、```bot```, ```event```を引数として```sendRestoInfo```関数を実行する。

```sendRestoInfo```に関しては次項で詳しく記述する。

<br>

***

## sendRestoInfo関数

### 位置情報の取得


```
msg := e.Message.(*linebot.LocationMessage)

lat := strconv.FormatFloat(msg.Latitude, 'f', 2, 64)
lng := strconv.FormatFloat(msg.Longitude, 'f', 2, 64)
```
BOTが受信した位置情報を```msg```変数に代入し、その緯度と軽度がFloat型で取得されるので```strconv.FormatFloat```で文字列に置換する。

置換した文字列を、変数```lat```, ```lng```にそれぞれ代入する。

<br>

### 返信メッセージの生成
```
replyMsg := getRestoInfo(lat, lng)

	res := linebot.NewTemplateMessage(
		"レストラン一覧",
		linebot.NewCarouselTemplate(replyMsg...).WithImageOptions("rectangle", "cover"),
	)
```
ユーザに返信するメッセージの内容を```replyMsg```で定義する。

HotpepperAPIと接続されている```getResyoInfo```を用いてメッセージを返信する。

メッセージを<b>カルーセルタイプ</b>で表示するため、```NewCarouselTemplate```を用いる。
カルーセルタイプのメッセージとは以下のようなもの。

<p>
<img src="https://developers.line.biz/assets/img/carousel.d89a53f5.png" width="50%" ,align="center">
</p>

<br>

***

## 構造体の定義

```
// response APIのレスポンス
type response struct {
	Results results `json:"results"`
}

// respinse APIのレスポンス内容
type results struct {
	Shop []shop `json:"shop"`
}

// shop(レストラン一覧)
type shop struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Photo   photo  `json:"ohoto"`
	URLS    urls   `json:"urls"`
}
```
以降```photo```, ```mobile```などのデータ型定義についても同様。

<b>構造体</b>とは、型の異なるデータを集めたデータ型で、```type (構造体の名前) struct {}``` で定義する。

また、Golangでは構造体とJSONの構造を紐付けて記述するのが一般的である。

その際、
<li>フィールド (ここでいう <b>Results</b>や<b>Shop</b>など) の先頭は大文字で始める必要がある</li>
<li>フィールドの右側に<b>`json:"xxx"`</b>と記述する必要がある</li>
以上2点に注意する。

<br>

***

## getRestoInfo関数

```
func getRestoInfo(lat string, lng string) []*linebot.CarouselColumn {
	apikey := "(ここにAPIKEYを追加)"
	url := fmt.Sprintf(
		"https://webservice.recruit.co.jp/hotpepper/gourmet/v1/?format=json&key=%s&lat=%s&lng=%s",
		apikey, lat, lng)
```

getRestoInfo関数は返り値があるので、```lat string```, ```lng string```のように引数の後に返り値の型を定義する。

また、変数```url```の部分で実行するAPIのURLを生成する。

<br>

### APIを実行してレスポンスを取得
前項で生成したAPIのURLにアクセスし実行する。
```
resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
```
```
var data response
	if err := json.Unmarshal(body, &data);
```
レスポンスで取得したJSONデータをGolangで扱える形にして、```response```型の```data```に格納する。

ここまでではHotpepperに掲載されている全ての情報が格納されている状態。

<br>

### 店名と住所のみを抽出してデータとして格納する
```
var ccs []*linebot.CarouselColumn
	for _, shop := range data.Results.Shop {
		addr := shop.Address
		if 60 < utf8.RuneCountInString(addr) {
			addr = string([]rune(addr)[:60])
		}
```
返信メッセージをカルーセルタイプにするため、1行目で```*linebot.CarouselColumn```型のスライスを宣言し、生成する。

前項で```data```に格納したJSONデータのうち、<b>店名</b>と<b>住所</b>を抽出し、```shop```, ```addr```にそれぞれ格納する。

同サイズのカルーセルカードに揃えるためには住所を60字以内に収める必要があるので、```if文```で61字以上ある場合に超える部分をカットする処理を追加する。



```
		cc := linebot.NewCarouselColumn(
			shop.Photo.Mobile.L,
			shop.Name,
			addr,
			linebot.NewURIAction("ホットペッパーで開く", shop.URLS.PC),
		).WithImageOptions("#ffffff")
		ccs = append(ccs, cc)
	}
	return ccs
```
```shop.Photo.Nobile.L```, ```shop.Name```などはそれぞれBOTga返信するメッセージの各項目に対応している。

```ccs = append(ccs, cc)```で2店舗目以降のスライスを追加する。

<br>

***

# Links
<li>ソースコード</li>

[rikako1021/Go-LineBot](https://github.com/rikako1021/Go-LineBot) 

<li>バグ報告、issue追加</li>

[New Issue · rikako1021/Go-LineBot](https://github.com/rikako1021/Go-LineBot/issues/new)

<li>Golangチュートリアル</li>

[A Tour of Go](https://tour.golang.org/welcome/1)

<li>LINEmessagingAPI 公式SDK</li>

[LINE Messaging API SDK for Go](https://github.com/line/line-bot-sdk-go)

<li>プロジェクト構成

[Goのプロジェクト構成の基本](https://zenn.dev/nobonobo/articles/4fb018a24f9ee9)

<li>Golang基本文法</li>

[【Go】基本文法総まとめ - Qiita](https://qiita.com/k-penguin-sato/items/deaeab18aa416496e273)

