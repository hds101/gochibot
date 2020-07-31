package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	botToken := os.Getenv("GACHIBOT_TOKEN")
	discord, err := discordgo.New("Bot " + botToken)
	if err != nil {
		fmt.Println("Error creating discord session,", err)
		return
	}

	usr, err := discord.User("@me")
	if err != nil {
		fmt.Println("Error obtaining account details,", err)
		return
	}

	fmt.Println(usr)
}
