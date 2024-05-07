package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

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
	// Ignore messages from the bot by checking if
	// the message author ID is the same as the bot's ID
	if message.Author.ID == discord.State.User.ID {
		return
	}

	prefix := "."
	// Handle messages
	switch {
	case message.Content == prefix+"help":
		discord.ChannelMessageSend(message.ChannelID, "Hello World")

	case message.Content == prefix+"deb":
		em := &discordgo.MessageEmbed{
			Title:       "Debora?",
			Description: "Também conhecida com Pathynx, Debora é o amor da vida de Vinícius no qual este a ama profundamente e possui uma imensa saudade que não pode ser descrita com palavras e somente aqueles que já sentiram tal sentimento podem compreender a falta que essa mulher faz na vida deste homem.\n\nTodavia, periodicamente, ambos se encontram em passeios diversos e, por assim dizer, _matam_ a saudade e podem viver em paz novamente. Entretanto, no fim do dia devem se despedir e essa acaba por se tornar a parte mais difícil de vosso dia, ocasionando a lamúria de ambas as partes e os deixando novamente em profunda depressão.\n\nDetermina-se então que essa mulher é e sempre será a peça vital do homem que lamenta sua ausência dia após dia e que tanto ele quanto ela devem possuir o gozo de se verem ao menos 1 (uma) vez por dia.",
			Image:       &discordgo.MessageEmbedImage{URL: "https://wallpapers.com/images/hd/jinx-arcane-opening-door-light-0gi6oz7gtv9690rh.jpg"},
			Timestamp:   time.Now().Format(time.RFC3339),
			Color:       0x2475ee,
		}
		discord.ChannelMessageSendEmbed(message.ChannelID, em)

	case message.Content == prefix+"pin":
		ch, err := discord.Channel(message.ChannelID)
		if err != nil {
			log.Fatal(err)
		}

		messages, _  := discord.ChannelMessages(ch.ID, 2, "", "", "")
		var msgIDToPin string
		for _, msg := range messages {
			if msg.ID == message.ID {
				continue
			}
			msgIDToPin = msg.ID
			break
		}

		if discord.ChannelMessagePin(ch.ID, msgIDToPin) == nil {
			discord.ChannelMessageSend(message.ChannelID, "Mensagem pinada com sucesso! :)")

			return
		}

		discord.ChannelMessageSend(message.ChannelID, "Falha ao pinar esta mensagem! :(")
	}
}
