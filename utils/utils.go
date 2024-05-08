package utils

import "github.com/bwmarrin/discordgo"

type Option struct {
	Option discordgo.ApplicationCommandOption
}

type NewChoice struct {
	Name  string
	Value interface{}
}

type ChannelType struct {
	Type discordgo.ChannelType
}

func CreateCommand(name, description string, options ...*discordgo.ApplicationCommandOption) *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        name,
		Description: description,
		Options:     options,
	}
}

func StringOption(name, description string, min, max int, required bool, choices ...[]*discordgo.ApplicationCommandOptionChoice) *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        name,
		Description: description,
		MinLength:   &min,
		MaxLength:   max,
		Required:    required,
	}
}

func IntegerOption(name, description string, min, max int, required bool, choices ...[]*discordgo.ApplicationCommandOptionChoice) *discordgo.ApplicationCommandOption {
	// Need to convert min and max to float64
	minFloat := float64(min)
	maxFloat := float64(max)
	return &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        name,
		Description: description,
		MinValue:    &minFloat,
		MaxValue:    maxFloat,
		Required:    required,
	}
}

func NumberOption(name, description string, min, max float64, required bool, choices ...[]*discordgo.ApplicationCommandOptionChoice) *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionNumber,
		Name:        name,
		Description: description,
		MinValue:    &min,
		MaxValue:    max,
		Required:    required,
	}
}

func BoolOption(name, description string, required bool, choices ...[]*discordgo.ApplicationCommandOptionChoice) *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionBoolean,
		Name:        name,
		Description: description,
		Required:    required,
	}
}

func UserOption(name, description string, required bool, choices ...[]*discordgo.ApplicationCommandOptionChoice) *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionUser,
		Name:        name,
		Description: description,
		Required:    required,
	}
}

func RoleOption(name, description string, required bool, choices ...[]*discordgo.ApplicationCommandOptionChoice) *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionRole,
		Name:        name,
		Description: description,
		Required:    required,
	}
}

func ChannelOption(name, description string, required bool, choices ...[]*discordgo.ChannelType) *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionChannel,
		Name:        name,
		Description: description,
		Required:    required,
		ChannelTypes: []discordgo.ChannelType{
			discordgo.ChannelTypeGuildText,
		},
	}
}

func Channels(channels ...ChannelType) []*discordgo.ChannelType {
	result := make([]*discordgo.ChannelType, len(channels))
	for i, channel := range channels {
		result[i] = &channel.Type
	}

	return result
}

func Choices(choices ...NewChoice) []*discordgo.ApplicationCommandOptionChoice {
	result := make([]*discordgo.ApplicationCommandOptionChoice, len(choices))
	for i, choice := range choices {
		result[i] = &discordgo.ApplicationCommandOptionChoice{
			Name:  choice.Name,
			Value: choice.Value,
		}
	}
	return result
}
