package main

import (
	"log"
	"os"
	"strings"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	client := twitch.NewClient(os.Getenv("TWITCH_USER"), os.Getenv("TWITCH_AUTH"))

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		if message.Tags["first-msg"] != "0" {
			client.Say(message.Channel, "@"+message.User.DisplayName+" have you seen this yet? https://i.nuuls.com/XHHZk.jpeg")
		}
	})

	channels := strings.Split(os.Getenv("TWITCH_CHANNELS"), ",")

	client.Join(channels...)
	// client.Say("mmattbtw", "TriHard hello from go")

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
