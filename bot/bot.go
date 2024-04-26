package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var BotToken string

func checkNilErr(err error) {
	if err != nil {
		log.Fatal("Error message")
	}
}

func Run() {
	// Create a session
	discord, err := discordgo.New("Bot " + BotToken)
	checkNilErr(err)

	// Add a event handler
	discord.AddHandler(newMessage)

	// Open session
	discord.Open()
	defer discord.Close() // Close session

	// Run until CTRL-C
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore messages from the bot
	if message.Author.ID == discord.State.User.ID {
		return
	}

	// Handle messages
	switch {
	case message.Content == "!help":
		discord.ChannelMessageSend(message.ChannelID, "Hello World")
	case message.Content == "!bye":
		discord.ChannelMessageSend(message.ChannelID, "Bye")
	}
}
