package discord

import (
	util "github.com/Floor-Gang/utilpkg"
	bututil "github.com/Floor-Gang/utilpkg/botutil"
	dg "github.com/bwmarrin/discordgo"
)

// Example command
func (bot *Bot) ping(msg *dg.MessageCreate) {
	member := msg.Member
	if member == nil {
		_, _ = bututil.Reply(bot.session, msg.Message, "Please use this command in a guild.")
		return
	}
	if !bot.checkChannel(msg.Message) {
		_, _ = bututil.Reply(bot.session, msg.Message,
			"Please use this command in the designated channel.",
		)
		return
	}

	if _, err := bututil.Reply(bot.session, msg.Message, "Pong"); err != nil {
		util.Report("Bot failed to reply to ping command.", err)
	}
}
