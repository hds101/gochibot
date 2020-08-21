package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	botToken := os.Getenv("GACHIBOT_TOKEN")

	discord, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}

	// https://discord.com/developers/docs/topics/gateway#gateway-intents
	// discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates)

	err = discord.Open()
	if err != nil {
		log.Fatal("Error opening Discord session: ", err)
	}

	// reg cmds
	discord.AddHandler(messageCreate)

	log.Println("Gochibot is now running, CTRL-C to stop")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func messageCreate(s *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(message.Content, "!test") {
		channel, err := s.State.Channel(message.ChannelID)
		if err != nil {
			log.Println("Couldnt find channel")
			return
		}

		s.ChannelMessageSend(channel.ID, "yo2")
	}
}
