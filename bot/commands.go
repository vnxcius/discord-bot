package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/vnxcius/discord-bot/utils"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		utils.CreateCommand("basic-command", "Basic command"),
		utils.CreateCommand("permission-overview", "Command for demonstration of default command permissions"),
		utils.CreateCommand("basic-command-with-files", "Basic command with files"),
		utils.CreateCommand("deb", "Quem é a Debora?",
			utils.UserOption("user-option", "Marque a gatinha se quiser", false),
		),
		utils.CreateCommand("followup", "Followup message"),
		utils.CreateCommand("options", "Command for demonstrating options",
			utils.StringOption("string-option", "String option", 1, 20, true),
			utils.IntegerOption("integer-option", "Integer option", 1, 10, true),
			utils.NumberOption("number-option", "Number option", 1, 10, true),
			utils.BoolOption("bool-option", "Bool option", true),
			utils.ChannelOption("channel-option", "Channel option", false, utils.Channels(
				utils.ChannelType{Type: discordgo.ChannelTypeGuildText},
				utils.ChannelType{Type: discordgo.ChannelTypeGuildVoice},
			)),
			utils.UserOption("user-option", "User option", false),
			utils.RoleOption("role-option", "Role option", false),
		),
		utils.CreateCommand("responses", "Interaction responses testing initiative",
			utils.IntegerOption("resp-type", "Response type", 4, 5, true, utils.Choices(
				utils.NewChoice{Name: "Channel message with source", Value: 4},
				utils.NewChoice{Name: "Deferred channel message with source", Value: 5},
			)),
		),
	}
)
