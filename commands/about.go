package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/jchambless/cacao/framework"
	"github.com/jchambless/cacao/util"
)

func AboutCommand(ctx framework.Context) {
	log.Printf("About command was called with args %q", ctx.Args)

	embed := &discordgo.MessageEmbed{
		Title:       "About " + ctx.Conf.BotName + " " + ctx.Conf.Version,
		Description: "A simple Discord bot to manage a Minecraft server from Discord instead of using the console. It doesn't support all commands yet, but it will in the future. It has recently been updated to support Discords new API/security features.",
		Color:       0xff0000,
		Fields:      []*discordgo.MessageEmbedField{},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Bot created by: " + util.GetAuthor() + " | " + util.GetAuthorUrl(),
		},
	}

	_, err := ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
	if err != nil {
		log.Println("Error sending help message:", err)
	}
}
