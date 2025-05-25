package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/jchambless/cacao/framework"
)

func HelpCommand(ctx framework.Context) {
	log.Println("Help command was called.")

	var cmds framework.CmdMap = ctx.CmdHandler.GetCmds()
	if cmds == nil {
		log.Println("No commands found.")
		return
	}

	var messageFields []*discordgo.MessageEmbedField

	for key, element := range cmds {
		if len(key) > 1 {
			messageFields = append(messageFields, &discordgo.MessageEmbedField{
				Name:   element.GetHelp(),
				Value:  element.GetDescription(),
				Inline: false,
			})
		}
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Bot Commands",
		Description: "Here are the commands you can use:",
		Color:       0x00ff00,
		Fields:      messageFields,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Use " + ctx.Conf.Prefix + "help <command> for more info",
		},
	}

	_, err := ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
	if err != nil {
		log.Println("Error sending help message:", err)
	}
}
