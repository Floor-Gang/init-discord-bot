package main

import (
	"github.com/Floor-Gang/init-discord-bot/internal"
	"github.com/Floor-Gang/init-discord-bot/internal/discord"
	util "github.com/Floor-Gang/utilpkg"
)

const (
	configPath = "./config.yml"
)

func main() {
	config := internal.GetConfig(configPath)

	discord.Start(config)

	util.KeepAlive()
}
