package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	//ハンドラ
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/callback", lineHandler)

	fmt.Println("http://localhost:8080 で起動中")
	//サーバきどう
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World ^_^; ^_^"
	fmt.Fprintf(w, msg)
}

func lineHandler(w http.ResponseWriter, r *http.Request) {
	//Bot初期化 (o^-^o)
	bot, err := linebot.New(
		"9ddc19f09c29661624873ec584880c3a",
		"8n0qXxvo/SRyXMPCErf8blmeM+N2Nh91UWSEf3zYC0JlCKXdrdDUajMFgL+3L0dW1xfgqj6CLApFrHzHHdSaymRaJgZPhK/8Ne2FDww6GTc7BO2QmGPjl3Sh0DOGnkCNG1n6cmGKOLOb5W3ayzV2bwdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		log.Fatal(err)
	}

	//Botイベント取得 (‥^▽^‥)
	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		//メッセージ受信のイベント
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			//テキスト形式のメッセージ
			case *linebot.TextMessage:
				replyMessage := message.Text
				_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				if err != nil {
					log.Print(err)
				}
			case *linebot.LocationMessage:
				sendResstroInfo(bot, event)
			}

		}
	}
}
func sendResstroInfo(bot *linebot.Client, e *linebot.Event) {
	msg := e.Message.(*linebot.LocationMessage)

	lat := strconv.FormatFloat(msg.Latitude, 'f', 2, 64)
	lng := strconv.FormatFloat(msg.Longitude, 'f', 2, 64)

	replyMsg := fmt.Sprintf("緯度：%s\n軽度：%s", lat, lng)

	_, err := bot.ReplyMessage(e.ReplyToken, linebot.NewTextMessage(replyMsg)).Do()
	if err != nil {
		log.Print(err)
	}
}
