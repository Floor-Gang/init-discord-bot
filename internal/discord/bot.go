package discord

import (
	"github.com/Floor-Gang/init-discord-bot/internal"
	util "github.com/Floor-Gang/utilpkg"
	dg "github.com/bwmarrin/discordgo"
)

// Bot structure
type Bot struct {
	version string
	session *dg.Session
	config  *internal.Config
}

// Start starts discord client, configuration and database
func Start(config internal.Config) error {
	var err error
	botConfig := internal.GetConfig(config.Location)

	client, err := dg.New("Bot " + botConfig.Token)

	if err != nil {
		panic(err)
	}

	bot := Bot{
		session: client,
		config:  &botConfig,
	}

	client.AddHandler(bot.onReady)
	client.AddHandler(bot.onMessage)

	if err = client.Open(); err != nil {
		util.Report("Was an authentication token provided?", err)
	}

	return err
}
