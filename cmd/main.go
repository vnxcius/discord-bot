package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/vnxcius/discord-bot/bot"
)

func init() {
	// load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}
}

func main() {
	bot.BotToken = os.Getenv("BOT_TOKEN")
	bot.Run()
}
