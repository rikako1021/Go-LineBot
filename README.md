# 環境設定

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

<br>

インストール先フォルダは特段の理由がなければデフォルトのまま変更しなくて良いです。

<br>

## インストールできているか確認
コマンドラインで以下のコマンドを実行し、Goのバージョンが表示されれば問題なくインストールできている。
```
$ go version
go version go1.16.15 darwin/arm64
```
<br>

## 環境変数（パス）の設定
PATHに```%GOPATH%\bin```を登録する。

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


