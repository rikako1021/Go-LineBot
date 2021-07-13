# 0.　アプリケーションの概要

### 背景

今すぐ食事を取りたい時、またはある土地に詳しくない時に、インターネット上で検索すると情報量が多すぎる

登録や複雑な手続きなく気軽に使える飲食店検索ツールが欲しい
         
<br>

### 機能

LINEのトーク上で送信された位置情報をもとに、BOTアカウントが周辺の飲食店の情報をHotpepperグルメから引用し表示する。

<br>

### 使用したリソース

<li>Golang </li>
<li>LINEmessagingAPI,HotpepperAPI</li>
<li>github - バージョン管理</li>
<li>delve(Golangパッケージ) - デバッグ</li>
<li>Heroku - デプロイ</li>

<br>

******

# 1.　環境設定

## Goのインストール

<br>

<b>・Mac</b>

コマンドラインで以下のコマンドを実行

```
$ brew install go
```

<br>

<b>・Windows</b>

<br>

[公式HP](https://golang.org/doc/install)より、Windows用インストーラーを選択してダウンロード＋実行

インストール先フォルダは特段の理由がなければデフォルトのまま変更しなくて良いです。

<br>

## インストールできているか確認
コマンドラインで以下のコマンドを実行し、Goのバージョンが表示されれば問題なくインストールできています。
```
$ go version
go version go1.16.15 darwin/arm64
```
<br>

## 環境変数（パス）の設定
PATHに```%GOPATH%\bin```を登録します。

この登録が無いと```go get```でインストールしたコマンドを実行できません。

<br>

<b>・Mac</b>

以下のコマンドを実行して設定します。

```
$ echo "export GOPATH=$(go env GOPATH)" >> ~/.bash_profile
$ echo "export PATH=$PATH:$(go env PATH)/bin" >> ~/.bash_profile
$ source ~/.bash_profile
```

使っているシェルがzshの場合は```.bash_profile```の部分を```.zshrc```に書き換えてください。

*環境変数の追加・編集などの方法は[こちら](https://gabekore.org/mac-path-environmental-variable)が参考になります。

<br>

<b>・Windows</b>


以下のページを参考にしてください。

 [WindowsにGo言語をインストールする方法まとめ - Qiita](https://qiita.com/yoskeoka/items/0dcc62a07bf5eb48dc4b#gopathbin-%E3%82%92path%E3%81%AB%E7%99%BB%E9%8C%B2%E3%81%99%E3%82%8B)

<br>

## VSCodeのインストール・設定

[こちら](https://code.visualstudio.com/download)からOSを選択してインストールします。

<br>

### 拡張機能のインストール

VSCodeを起動して、```Ctrl+Shift+X```で拡張機能を開いて```Go```で検索します。

検索結果に表示される以下の拡張機能をインストールします。

![GO extention](https://user-images.githubusercontent.com/68047214/125348202-e3b7b880-e396-11eb-9cb5-36a51f2c78b9.png)

<br>

### 主な機能

この拡張機能では、
<li>Lint＆Format</li>
<li>デバッグ</li>
<li>コード補完</li>
を自動的にやってくれます。

<br>

******

# 2. コードをクローンして実行する

作業したいディレクトリで以下のコマンドを実行します。

```
$ git clone https://github.com/rikako1021/Go-LineBot.git
```

<br>

作業ディレクトリ内に<b>main.go</b>があることを確認し、以下のコマンドで実行します。

```
$ go run main.go
```

これでサーバが立ちます。

<br>

# 3. 各APIの設定と連携

## LINE APIの設定

### LINE Developersに登録

MessagingAPIを利用するにはLINEDevelopersへの登録が必要です。

以下の公式サイトを参考に、LINEDevelopersアカウントとチャネルを作成します。

[LINE Developersコンソールでチャネルを作成する](https://developers.line.biz/ja/docs/messaging-api/getting-started/)

[LINE Developersコンソールでボットを作成する](https://developers.line.biz/ja/docs/messaging-api/building-bot/)

<br>

### LINE Developersでシークレットを取得

LINE Developersの各チャネルページの<b>Basic settings</b>の下の方に<b>Channel secret</b>があります。

![secret1](https://user-images.githubusercontent.com/68047214/125357230-511d1680-e3a2-11eb-9524-1090ed91db01.png)

![secret2](https://user-images.githubusercontent.com/68047214/125357210-46fb1800-e3a2-11eb-8930-76b07dc5f7ea.png)

<br>

### LINE Developersでアクセストークンを取得

LINE Developersの各チャネルページの<b>Messaging API</b>の下の方に<b>Channel access token</b>があります。

![token](https://user-images.githubusercontent.com/68047214/125357183-3c408300-e3a2-11eb-940d-4f9a1bba8caa.png)

<br>

### 取得した認証情報をコードに追加

lineHandler関数の(secret),(access token)の部分にそれぞれ、上記で取得したシークレット・アクセストークンをそれぞれ追加します。

```main.go

func lineHandler(w http.ResponseWriter, r *http.Request) {
	//Bot初期化
	bot, err := linebot.New(
		"(secret)",
		"(access token)",
	)
	if err != nil {
		log.Fatal(err)
	}
```

<br>

## Herokuの設定

### アカウント登録

まずはHerokuのアカウントを取得します。
以下のページより必要事項を記入の上、「無料アカウント作成」をクリックしてください。

https://signup.heroku.com/jp

入力したメールアドレスに認証メールが届きますので、メッセージ内のリンクをクリックしてアカウントを有効化します。

最後にパスワードの設定を求められるので、任意のパスワードを設定してください。



![heroku](https://user-images.githubusercontent.com/68047214/125358517-f2589c80-e3a3-11eb-8250-9b0f5519132d.png)


### Heroku CLIのインストール

次に、Herokuをコマンドライン上から操作するためのツールをインストールします。

デプロイにはこちらのツールが必要になります。

<br>

<b>・Mac</b>

ターミナル上で下記コマンドを実行します。

```
brew tap heroku/brew && brew install heroku
```


<b>・Windows</b>

以下のページでWindows用のインストーラが提供されているので、ダウンロードし起動します。

[The Heroku CLI | Heroku Dev Center](https://devcenter.heroku.com/articles/heroku-cli)

インストール時の設定は全てデフォルトの状態で進み、```install```ボタンをクリックするとインストールされます。

<br>

インストールが完了したら、CLIを起動してみます。
Macであればターミナルで、WindowsであればコマンドプロンプトやGitBashで以下のコマンドを入力します。
```
heroku --version
```

これでバージョン情報が表示されれば、Heroku CLIのインストールは完了です。

<br>

### Herokuにログインする

先ほど作成したアカウントでHerokuにログインします。

```heroku login -i```を入力するとコマンドライン上でログインを要求されるので、メールアドレスとパスワードを入力します。

（```-i```抜きで実行するとブラウザ上にHerokuのログイン画面が立ち上がり、ブラウザ上でログイン認証を行います。）


<br>

# 4. デプロイ

ここからは、Herokuを用いて実際にデプロイしていきます。

### アプリケーション作成

まずは、クローンしてきたアプリケーションのディレクトリへ移動します。

```
$ cd Go-LineBot
```

次に、```heroku create```コマンドを入力すると、デプロイ用のアプリケーションが作成されます。

```
$ heroku create
Creating app... done, XXXX-ZZZZ-123456
https://XXXX-ZZZZ-123456.herokuapp.com/ | https://git.heroku.com/XXXX-ZZZZ-123456.git
```

ここで示される<b>XXXX-ZZZZ-123456</b>（例）が、作成されたアプリケーション名を表しています。

指定がない場合、<b>XXXX-ZZZZ-123456</b>のようなランダムな名称が割り振られます。

　<br>
 
### デプロイ

```git push heroku master```を入力すると、Heroku上へのデプロイが行われます。

```
$ git push heroku master
Enumerating objects: 400, done.
Counting objects: 100% (400/400), done.
Delta compression using up to 4 threads.
Compressing objects: 100% (190/190), done.
Writing objects: 100% (400/400), 178.59 KiB | 25.51 MiB/s, done.
Total 400 (delta 152), reused 400 (delta 152)
remote: Compressing source files... done.
・
・
・
remote:        https://XXXX-ZZZZ-123456.herokuapp.com/ deployed to Heroku
remote:
remote: Verifying deploy... done.
To https://git.heroku.com/XXXX-ZZZZ-123456.git
 * [new branch]      master -> master
```

これでデプロイ作業は完了です。

プログラムに変更を加え、Herokuに反映させたい場合は通常Githubにプッシュする時と同様に行います。

```
$ git push heroku master
```

<br>

### アクセスしてみる

デプロイした時に表示されたURL（例の場合は```https://XXXX-ZZZZ-123456.herokuapp.com/```）にアクセスし、HelloWorldなどの画面が表示されていれば正常にデプロイが完了しています。

<br>

### LINEBOTのWebhookに追加する

デプロイ先ページのURL+```/webhook```をLINE DevelopersのMessagingAPIの部分にあるWebhookURLとして設定します。

![webhook](https://user-images.githubusercontent.com/68047214/125362340-b0caf000-e3a9-11eb-9930-b2638f0bc1ff.png)

URLを入力後、「verify」ボタンをクリックし、```success```と表示されれば正常に接続できています。

<br>

# 5. 動作確認

LINE Developersの各チャネルページの<b>Messaging API</b>の上の方に<b>アカウントのQRコード</b>があります。

QRコードを読み取ってアカウントを友達追加し、実際にメッセージを送信してみます。

トーク画面下部の「＋」ボタンから位置情報を送信し、このように返信が来たらOKです。

<img src="https://user-images.githubusercontent.com/68047214/125363511-b3c6e000-e3ab-11eb-806e-cf36f61623b7.jpg" width="30%">
