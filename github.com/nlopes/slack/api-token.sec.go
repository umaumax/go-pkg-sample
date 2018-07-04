package main

import (
	"os"
)

func init() {
	BotAPIToken = os.Getenv("BOT_API_TOKEN")
}
