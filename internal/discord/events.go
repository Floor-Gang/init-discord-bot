package discord

import (
	"log"

	dg "github.com/bwmarrin/discordgo"
)

func (b *Bot) onMessage(s *dg.Session, m *dg.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.ChannelID != b.config.ChannelID {
		if m.Content == "Ping" {
			s.ChannelMessageSend(b.config.ChannelID, "Pong!")
		}
	}
}

func (b *Bot) onReady(s *dg.Session, _ *dg.Ready) {
	log.Printf("Ready as %s (version %s)\n", s.State.User.Username, b.version)
}
