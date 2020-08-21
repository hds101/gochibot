package main

import (
	"gochibot/src/commands"
	"gochibot/src/lib"

	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var CommandHandler *lib.CommandHandler

func main() {
	CommandHandler = lib.InitCommandHandler()
	CommandHandler.Register("!temp", commands.TempCommand, "help message")

	discord, err := discordgo.New("Bot " + os.Getenv("GACHIBOT_TOKEN"))
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}

	discord.AddHandler(handleMessage)

	// https://discord.com/developers/docs/topics/gateway#gateway-intents
	// discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates)

	err = discord.Open()
	if err != nil {
		log.Fatal("Error opening Discord session: ", err)
	}

	log.Println("Gochibot is now running, CTRL-C to stop")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func handleMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	author := message.Author
	if author.ID == session.State.User.ID {
		return
	}

	content := message.Content
	// TODO: check prefix

	args := strings.Fields(content)
	commandName := strings.ToLower(args[0])

	command, found := CommandHandler.Get(commandName)
	if !found {
		return
	}

	channel, err := session.State.Channel(message.ChannelID)
	if err != nil {
		log.Println("Couldnt find channel: ", err)
		return
	}

	guild, err := session.State.Guild(channel.GuildID)
	if err != nil {
		log.Println("Couldnt find guild: ", err)
		return
	}

	context := lib.InitContext(session, guild, channel, author, message, args[1:])

	cmd := *command
	cmd(*context)
}
