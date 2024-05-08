package bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"basic-command": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Hey there! Congratulations, you just executed your first slash command",
			},
		})
	},
	"deb": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Também conhecida com Pathynx, Debora é o amor da vida de Vinícius no qual este a ama profundamente e possui uma imensa saudade que não pode ser descrita com palavras e somente aqueles que já sentiram tal sentimento podem compreender a falta que essa mulher faz na vida deste homem.\n\nTodavia, periodicamente, ambos se encontram em passeios diversos e, por assim dizer, _matam_ a saudade e podem viver em paz novamente. Entretanto, no fim do dia devem se despedir e essa acaba por se tornar a parte mais difícil de vosso dia, ocasionando a lamúria de ambas as partes e os deixando novamente em profunda depressão.\n\nDetermina-se então que essa mulher é e sempre será a peça vital do homem que lamenta sua ausência dia após dia e que tanto ele quanto ela devem possuir o gozo de se verem ao menos 1 (uma) vez por dia.",
			},
		})
	},
	"basic-command-with-files": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Hey there! Congratulations, you just executed your first slash command with a file in the response",
				Files: []*discordgo.File{
					{
						ContentType: "text/plain",
						Name:        "test.txt",
						Reader:      strings.NewReader("Hello Discord!!"),
					},
				},
			},
		})
	},
	"options": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Access options in the order provided by the user.
		options := i.ApplicationCommandData().Options

		// Or convert the slice into a map
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		// This example stores the provided arguments in an []interface{}
		// which will be used to format the bot's response
		margs := make([]interface{}, 0, len(options))
		msgformat := "You learned how to use command options! " +
			"Take a look at the value(s) you entered:\n"

		// Get the value from the option map.
		// When the option exists, ok = true
		if option, ok := optionMap["string-option"]; ok {
			// Option values must be type asserted from interface{}.
			// Discordgo provides utility functions to make this simple.
			margs = append(margs, option.StringValue())
			msgformat += "> string-option: %s\n"
		}

		if opt, ok := optionMap["integer-option"]; ok {
			margs = append(margs, opt.IntValue())
			msgformat += "> integer-option: %d\n"
		}

		if opt, ok := optionMap["number-option"]; ok {
			margs = append(margs, opt.FloatValue())
			msgformat += "> number-option: %f\n"
		}

		if opt, ok := optionMap["bool-option"]; ok {
			margs = append(margs, opt.BoolValue())
			msgformat += "> bool-option: %v\n"
		}

		if opt, ok := optionMap["channel-option"]; ok {
			margs = append(margs, opt.ChannelValue(nil).ID)
			msgformat += "> channel-option: <#%s>\n"
		}

		if opt, ok := optionMap["user-option"]; ok {
			margs = append(margs, opt.UserValue(nil).ID)
			msgformat += "> user-option: <@%s>\n"
		}

		if opt, ok := optionMap["role-option"]; ok {
			margs = append(margs, opt.RoleValue(nil, "").ID)
			msgformat += "> role-option: <@&%s>\n"
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			// Ignore type for now, they will be discussed in "responses"
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf(
					msgformat,
					margs...,
				),
			},
		})
	},
	"subcommands": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		content := ""

		// As you can see, names of subcommands (nested, top-level)
		// and subcommand groups are provided through the arguments.
		switch options[0].Name {
		case "subcommand":
			content = "The top-level subcommand is executed. Now try to execute the nested one."
		case "subcommand-group":
			options = options[0].Options
			switch options[0].Name {
			case "nested-subcommand":
				content = "Nice, now you know how to execute nested commands too"
			default:
				content = "Oops, something went wrong.\n" +
					"Hol' up, you aren't supposed to see this message."
			}
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: content,
			},
		})
	},
	"responses": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Responses to a command are very important.
		// First of all, because you need to react to the interaction
		// by sending the response in 3 seconds after receiving, otherwise
		// interaction will be considered invalid and you can no longer
		// use the interaction token and ID for responding to the user's request

		content := ""
		// As you can see, the response type names used here are pretty self-explanatory,
		// but for those who want more information see the official documentation
		switch i.ApplicationCommandData().Options[0].IntValue() {
		case int64(discordgo.InteractionResponseChannelMessageWithSource):
			content =
				"You just responded to an interaction, sent a message and showed the original one. " +
					"Congratulations!"
			content +=
				"\nAlso... you can edit your response, wait 5 seconds and this message will be changed"
		default:
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseType(i.ApplicationCommandData().Options[0].IntValue()),
			})
			if err != nil {
				s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Content: "Something went wrong",
				})
			}
			return
		}

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseType(i.ApplicationCommandData().Options[0].IntValue()),
			Data: &discordgo.InteractionResponseData{
				Content: content,
			},
		})
		if err != nil {
			s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
				Content: "Something went wrong",
			})
			return
		}
		time.AfterFunc(time.Second*5, func() {
			content := content + "\n\nWell, now you know how to create and edit responses. " +
				"But you still don't know how to delete them... so... wait 10 seconds and this " +
				"message will be deleted."
			_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Content: &content,
			})
			if err != nil {
				s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Content: "Something went wrong",
				})
				return
			}
			time.Sleep(time.Second * 10)
			s.InteractionResponseDelete(i.Interaction)
		})
	},
	"followups": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Followup messages are basically regular messages (you can create as many of them as you wish)
		// but work as they are created by webhooks and their functionality
		// is for handling additional messages after sending a response.

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				// Note: this isn't documented, but you can use that if you want to.
				// This flag just allows you to create messages visible only for the caller of the command
				// (user who triggered the command)
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "Surprise!",
			},
		})
		msg, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: "Followup message has been created, after 5 seconds it will be edited",
		})
		if err != nil {
			s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
				Content: "Something went wrong",
			})
			return
		}
		time.Sleep(time.Second * 5)

		content := "Now the original message is gone and after 10 seconds this message will ~~self-destruct~~ be deleted."
		s.FollowupMessageEdit(i.Interaction, msg.ID, &discordgo.WebhookEdit{
			Content: &content,
		})

		time.Sleep(time.Second * 10)

		s.FollowupMessageDelete(i.Interaction, msg.ID)

		s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: "For those, who didn't skip anything and followed tutorial along fairly, " +
				"take a unicorn :unicorn: as reward!\n" +
				"Also, as bonus... look at the original interaction response :D",
		})
	},
}
