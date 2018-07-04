package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/nlopes/slack"
)

var (
	BotAPIToken = string
)

func main() {
	api := slack.New(BotAPIToken)
	api.SetDebug(true)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				// Ignore hello

			case *slack.ConnectedEvent:
				fmt.Println("Infos:", ev.Info)
				fmt.Println("Connection counter:", ev.ConnectionCount)
				// Replace #general with your Channel ID
				//				rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "#general"))

			case *slack.MessageEvent:
				me := msg.Data.(*slack.MessageEvent)
				//	NOTE botのメッセージには反応しない
				if me.SubType == "bot_message" {
					break
				}

				fmt.Printf("Message: %v\n", ev)
				text := strings.ToLower(me.Text)
				if strings.HasPrefix(text, "hello") {
					slack.MessageEvent.SubType
					text := fmt.Sprintf("Hello: %s(%s)", me.Username, me.User)
					channel := me.Channel
					params := slack.PostMessageParameters{}
					channelID, timestamp, err := api.PostMessage(channel, text, params)
					if err != nil {
						log.Println("post message error:", err)
						break
					}
					log.Printf("post Channel(%s:ID(%s)) at %s : %s\n%v\n", channel, channelID, timestamp, text, params)
				}

			case *slack.PresenceChangeEvent:
				fmt.Printf("Presence Change: %v\n", ev)

			case *slack.LatencyReport:
				fmt.Printf("Current latency: %v\n", ev.Value)

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:

				// Ignore other events..
				// fmt.Printf("Unexpected: %v\n", msg.Data)
			}
		}
	}
}
