package main

import (
	"encoding/json"
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"unicode/utf8"
)

func main() {
	//ハンドラ
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/callback", lineHandler)

	fmt.Println("http://localhost:5000 で起動中")

	//サーバきどう
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World ^_^; ^_^"
	fmt.Fprint(w, msg)
}

func lineHandler(w http.ResponseWriter, r *http.Request) {
	//Bot初期化
	bot, err := linebot.New(
		"9ddc19f09c29661624873ec584880c3a",
		"AObAf46/9JgGeuWvu0xATv8/jAL20ObglW8D2SmbuFSUVWR+XSwsNC7dvpeVaZvg1xfgqj6CLApFrHzHHdSaymRaJgZPhK/8Ne2FDww6GTe9L3GKUJpH8XBbYI8yTLohh9DlyMD0Xnj3PgieaVHZcAdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		log.Fatal(err)
	}

	//Botイベント取得
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
				sendRestoInfo(bot, event)
			}

		}
	}
}

func sendRestoInfo(bot *linebot.Client, e *linebot.Event) {
	msg := e.Message.(*linebot.LocationMessage)

	lat := strconv.FormatFloat(msg.Latitude, 'f', 2, 64)
	lng := strconv.FormatFloat(msg.Longitude, 'f', 2, 64)

	replyMsg := getRestoInfo(lat, lng)

	res := linebot.NewTemplateMessage(
		"レストラン一覧",
		linebot.NewCarouselTemplate(replyMsg...).WithImageOptions("rectangle", "cover"),
	)

	if _, err := bot.ReplyMessage(e.ReplyToken, res).Do(); err != nil {
		log.Print(err)
	}
}

// responseAPI
type response struct {
	Results results `json:"results"`
}

// respinseAPI レスポンス内容
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

// photo 写真URL一覧
type photo struct {
	Mobile mobile `json:"mobile"`
}

// mobile モバイル用
type mobile struct {
	L string `json:"l"`
}

// urls URL一覧
type urls struct {
	PC string `json:"pc"`
}

func getRestoInfo(lat string, lng string) []*linebot.CarouselColumn {
	apikey := "8d6f6ca4b5d9872e"
	url := fmt.Sprintf(
		"https://webservice.recruit.co.jp/hotpepper/gourmet/v1/?format=json&key=%s&lat=%s&lng=%s",
		apikey, lat, lng)

	// ボディ取得
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data response
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	var ccs []*linebot.CarouselColumn
	for _, shop := range data.Results.Shop {
		addr := shop.Address
		if 60 < utf8.RuneCountInString(addr) {
			addr = string([]rune(addr)[:60])
		}
		cc := linebot.NewCarouselColumn(
			shop.Photo.Mobile.L,
			shop.Name,
			addr,
			linebot.NewURIAction("ホットペッパーで開く", shop.URLS.PC),
		).WithImageOptions("#ffffff")
		ccs = append(ccs, cc)
	}
	return ccs
}
