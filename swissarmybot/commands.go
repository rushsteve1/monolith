package swissarmybot

import (
	"github.com/bwmarrin/discordgo"
)

var commands = []*discordgo.ApplicationCommand{
	{
		Name: "Add Quote",
		Type: discordgo.MessageApplicationCommand,
	},
	{
		Name:        "quote",
		Description: "",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "add",
				Description: "Add a quote to the database",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "who",
						Description: "Who said this?",
						Type:        discordgo.ApplicationCommandOptionUser,
						Required:    true,
					},
					{
						Name:        "text",
						Description: "What did they say?",
						Type:        discordgo.ApplicationCommandOptionString,
						Required:    true,
					},
				},
			},
			{
				Name:        "remove",
				Description: "Remove a quote from the database",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "id",
						Description: "What ID number should be removed?",
						Type:        discordgo.ApplicationCommandOptionInteger,
						Required:    true,
					},
				},
			},
			{
				Name:        "get",
				Description: "Get a quote from the database",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "id",
						Description: "What ID number should I get?",
						Type:        discordgo.ApplicationCommandOptionInteger,
						Required:    true,
					},
				},
			},
			{
				Name:        "list",
				Description: "List all the quotes by a user",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "who",
						Description: "What user should I list?",
						Type:        discordgo.ApplicationCommandOptionMentionable,
						Required:    false},
				},
			},
		},
	},
}
