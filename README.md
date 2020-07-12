# Init-discord-bot

Initial package for creating a discord bot.

## Pre-configs
To test this project locally, you need to 
* [create a discord server](https://www.howtogeek.com/318890/how-to-set-up-your-own-discord-chat-server/) 
* [enable developer mode on your discord account](https://discordia.me/en/developer-mode)
* [get a bot token](https://github.com/reactiflux/discord-irc/wiki/Creating-a-discord-bot-&-getting-a-token)
* Get the channel ID where you want your bot to send messages by right clicking on the channel and choosing "Copy ID"
* Get the server ID of your discord server by right clicking on the server icon and choosing "Copy ID"

**Make note of the bot token, channel ID, and server ID.**

## Run 
```go
git clone git@github.com:Floor-Gang/init-discord-bot.git
cd init-discord-bot
go get .
go run main.go
```
This will show an error but it'll create a `config.yml` file in `init-discord-bot` directory. Open it and put the information from Pre-configs section like below:
```yml
token: Paste your bot token here
prefix: (does not work for this bot) the prefix that you want your bot to recognize. You can leave it empty for now
channel: Channel ID that you copied in pre-configs
guild: Server ID that you copied in pre-configs
```
Then run 
```go
go run main.go
```

Type ping in your discord server's general channel. You should see a "Pong!" response in the designated channel by the bot.

