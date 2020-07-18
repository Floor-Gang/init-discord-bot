package discord

import (
	"github.com/Floor-Gang/init-discord-bot/internal"
	dg "github.com/bwmarrin/discordgo"
)

// Util struct methods
func (b *Bot) checkChannel(commandMessage *dg.Message) bool {
	return commandMessage.ChannelID == b.config.ChannelID
}

func (b *Bot) checkRoles(member *dg.Member) bool {
	return internal.StringInSlice(b.config.LeadDevID, member.Roles)
}

func checkAction(action string) bool {
	if action == "remove" || action == "filter" {
		return true
	}
	return false
}
