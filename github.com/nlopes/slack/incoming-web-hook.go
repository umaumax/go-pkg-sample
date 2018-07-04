package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var (
	IncomingUrl string = os.Getenv("INCOMING_URL")
)

type Slack struct {
	Text      string `json:"text"`       //投稿内容
	Username  string `json:"username"`   //投稿者名 or Bot名（存在しなくてOK）
	IconEmoji string `json:"icon_emoji"` //アイコン絵文字
	IconUrl   string `json:"icon_url"`   //アイコンURL（icon_emojiが存在する場合は、適応されない）
	Channel   string `json:"channel"`    //#部屋名
}

func main() {
	params, _ := json.Marshal(Slack{
		Text:      "Hello World",
		Username:  "gopher",
		IconEmoji: "",
		//Icon_url:		"http://www.ensky.co.jp/item/images/save/07151842_53c4f7839cd4e.jpg",
		IconUrl: "https://qiita-image-store.s3.amazonaws.com/0/14952/5c851c55-98c9-0b5b-5e21-39c136a05844.png",
		Channel: "#bot"})
	//	Channel:"#random"})

	resp, _ := http.PostForm(
		IncomingUrl,
		url.Values{"payload": {string(params)}},
	)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(string(body))
}
