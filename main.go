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

	usr, err := discord.User("@me")
	if err != nil {
		log.Fatal("Error obtaining account details,", err)
	}

	botId = usr.ID
	discord.AddHandler(commandHandler)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		discord.UpdateStatus(0, conf.DefaultStatus)
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
	channel, err := discord.State.Channel(message.ChannelID)
	if err != nil {
		log.Println("Error getting channel,", err)
		return
	}
	guild, err := discord.State.Guild(channel.GuildID)
	if err != nil {
		log.Println("Error getting guild,", err)
		return
	}
	ctx := framework.NewContext(discord, guild, channel, user, message, conf, CmdHandler)
	ctx.Args = args[1:]
	c := *command
	c(*ctx)
}

func registerCommands() {
	CmdHandler.Register("help", commands.HelpCommand, "Get help about this bots commands.")
	CmdHandler.Register("banplayer", commands.BanPlayerCommand, "Ban a player from the Minecraft server.")
	CmdHandler.Register("banip", commands.BanIPCommand, "Ban a IP address from server.")
	CmdHandler.Register("banlist", commands.BanListCommand, "Displays the banlist.")
	CmdHandler.Register("kick", commands.KickCommand, "Forcibly disconnects playername from the server, displaying an optional reason to them.")
	CmdHandler.Register("deop", commands.DeopCommand, "Revokes a player's operator status.")
	CmdHandler.Register("op", commands.OpCommand, "Grants playername operator status on the server.")
	CmdHandler.Register("stop", commands.ShutdownCommand, "Gracefully shuts down the server.")
	CmdHandler.Register("list", commands.PlayerListCommand, "Shows the names of all currently-connected players.")
}
