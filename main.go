package main

import (
	"fmt"

	"github.com/wanton-idol/DiscordBot/bot"
	"github.com/wanton-idol/DiscordBot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
