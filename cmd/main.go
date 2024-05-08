package main

import (
	"flag"
	"log"
	"log/slog"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/vnxcius/discord-bot/bot"
)

// Bot parameters
var (
	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

var s *discordgo.Session

func init() {
	var err error

	// Load environment variables
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		os.Exit(1)
	}

	// Load flags
	flag.Parse()

	// Create a new Discord session using the provided bot token.
	s, err = discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	// Register commands
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := bot.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func main() {
	// Create logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		logger.Info("Logged in as: " + s.State.User.Username + "#" + s.State.User.Discriminator)
	})
	err := s.Open()
	if err != nil {
		logger.Error("Cannot open the session", "error", err)
		os.Exit(1)
	}

	// Register slash commands
	logger.Info("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(bot.Commands))
	for i, v := range bot.Commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, *GuildID, v)
		if err != nil {
			logger.Error("Cannot create command", "command", v.Name, "error", err)
			panic(err)
		}
		registeredCommands[i] = cmd
	}
	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	if *RemoveCommands {
		logger.Info("Removing commands...")
		// We need to fetch the commands, since deleting requires the command ID.
		// We are doing this from the returned commands on line 375, because using
		// this will delete all the commands, which might not be desirable, so we
		// are deleting only the commands that we added.
		registeredCommands, err := s.ApplicationCommands(s.State.User.ID, *GuildID)
		if err != nil {
			logger.Error("Could not fetch registered commands", "error", err)
			os.Exit(1)
		}

		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, *GuildID, v.ID)
			if err != nil {
				logger.Error("Cannot delete command", "command", v.Name, "error", err)
				panic(err)
			}
		}
	}

	logger.Info("Gracefully shutting down.")
}
