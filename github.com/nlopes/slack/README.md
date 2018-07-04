# Slack API (by golang)

`.envrc`
```
export BOT_API_TOKEN="xxx"
export INCOMING_URL="xxx"
```

----

* [Slack Platform: Community | Slack](https://api.slack.com/community)
* [nlopes/slack: Slack API in Go](https://github.com/nlopes/slack)
	* Bot Users
	* Web API
	* Real Time Messaging API

ローカル環境でslackを管理するサービスを動かすならばReal Time Messaging APIの一択かな?!
[SlackのBotを書いてみた | とさいぬの隠し部屋](https://blog.myon.info/blog/2016-01-24/slack-bot/)

## my slack page
http://my.slack.com/apps/manage

[Slack APIを使ってプッシュ通知を受け取ろう | 株式会社バニーホップ](https://archive.bunnyhop.jp/lab-20141020/)

## Incoming Web Hook
Slackが発行するURLへPOSTするとSlackにメッセージが投稿される （チャンネル・ユーザ名・アイコン画像の指定可能）

## Outgoing WebHooks
Slackへ指定したメッセージが投稿されると、指定したURLに投稿内容がPOSTされる （POSTされるサーバ・システムを実装する必要あり）また、決められた形式（json）のレスポンスを返すことで同時にSlack上へ投稿することも可能
グローバルサーバが必要(heroku (with hubot) or amazon (with hubot), google apps script...)

## Slack API
メッセージの投稿に加え、ファイルのアップロード・チーム情報やチャンネル情報の取得等が利用可能

```
# github.com/nlopes/slack
slack.MessageEvent.
	Channel : メッセージが投稿されたチャンネルの固有ID
	Username : メッセージを投稿したユーザの固有ID
	Text : メッセージの内容
	SubType : 'message' or 'bot_message' or ...

```

