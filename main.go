package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/jchambless/cacao/framework"

	"github.com/jchambless/cacao/commands"

	"github.com/bwmarrin/discordgo"
)

var (
	PREFIX     string
	botId      string
	conf       *framework.Config
	CmdHandler *framework.CommandHandler
)

func init() {
	conf = framework.LoadConfig()
	PREFIX = conf.Prefix
}

func main() {
	CmdHandler = framework.NewCommandHandler()
	registerCommands()

	discord, err := discordgo.New("Bot " + conf.BotToken)
	if err != nil {
		log.Fatal("Error creating Discord session", err)
	}

	discord.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages | discordgo.IntentsMessageContent

	discord.State.TrackChannels = true

	usr, err := discord.User("@me")
	if err != nil {
		log.Fatal("Error obtaining account details,", err)
	}

	botId = usr.ID
	discord.AddHandler(commandHandler)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		discord.UpdateCustomStatus(conf.DefaultStatus)
		guilds := discord.State.Guilds
		log.Println("Ready with", len(guilds), "guilds.")
	})

	err = discord.Open()
	if err != nil {
		log.Fatal("Error opening connection", err)
	}
	defer discord.Close()

	fmt.Println("Bot is now running. Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.ID == botId || user.Bot {
		return
	}

	content := message.Content
	if len(content) <= len(PREFIX) {
		return
	}

	if content[:len(PREFIX)] != PREFIX {
		return
	}

	content = content[len(PREFIX):]
	if len(content) < 1 {
		return
	}

	args := strings.Fields(content)
	name := strings.ToLower(args[0])
	command, found := CmdHandler.Get(name)
	if !found {
		log.Println("Command not found,", name)
		return
	}

	var channel *discordgo.Channel
	channel, err := discord.State.Channel(message.ChannelID)
	if err != nil {
		// Fallback to API call if not found in state
		log.Println("Channel not in state cache, fetching from API")
		channel, err = discord.Channel(message.ChannelID)
		if err != nil {
			log.Println("Error getting channel,", err)
			return
		}
	}

	var guild *discordgo.Guild
	guild, err = discord.State.Guild(channel.GuildID)
	if err != nil {
		// Fall back to API call if state doesn't have the guild
		log.Println("Guild not in state cache, fetching from API")
		guild, err = discord.Guild(channel.GuildID)
		if err != nil {
			log.Println("Error getting guild:", err)
			return
		}
	}

	ctx := framework.NewContext(discord, guild, channel, user, message, conf, CmdHandler)
	ctx.Args = args[1:]
	c := *command
	c(*ctx)
}

func registerCommands() {
	CmdHandler.Register("help", commands.HelpCommand, "!mc help", "Get help about this bots commands.")
	CmdHandler.Register("ban", commands.BanPlayerCommand, "!mc ban <player>", "Ban a player from the Minecraft server.")
	CmdHandler.Register("banip", commands.BanIPCommand, "!mc banip <ip>", "Ban a IP address from server.")
	//CmdHandler.Register("banlist", commands.BanListCommand, "!mc banlist", "Displays the banlist.")
	CmdHandler.Register("kick", commands.KickCommand, "!mc kick <player>", "Forcibly disconnects playername from the server, displaying an optional reason to them.")
	CmdHandler.Register("deop", commands.DeOperatorCommand, "!mc deop <player>", "Revokes a player's operator status.")
	CmdHandler.Register("op", commands.OperatorCommand, "!mc op <player>", "Grants playername operator status on the server.")
	CmdHandler.Register("stop", commands.ShutdownCommand, "!mc stop", "Gracefully shuts down the server.")
	CmdHandler.Register("list", commands.PlayerListCommand, "!mc list", "Shows the names of all currently-connected players.")
	CmdHandler.Register("pardon", commands.PardonPlayerCommand, "!mc pardon <player>", "Removes player from the blacklist, allowing them to connect again.")
	CmdHandler.Register("ping", commands.PingCommand, "!mc ping", "Checks if server is up and returns information regarding it.")
	CmdHandler.Register("about", commands.AboutCommand, "!mc about", "Get information about the bot.")
}
