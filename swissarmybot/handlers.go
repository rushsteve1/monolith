package swissarmybot

import (
	"github.com/bwmarrin/discordgo"
)

var handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}
