package commands

import (
	"log"

	"github.com/jchambless/cacao/framework"
)

func HelpCommand(ctx framework.Context) {
	log.Println("Help command was called.")

	var message string

	// Todo: Build the string output before sending a reply
	message = ctx.Conf.BotName + " running v " + ctx.Conf.Version + "\n"
	message += "**Command List**\n"

	cmds := ctx.CmdHandler.GetCmds()
	for key, element := range cmds {
		if len(key) > 1 {
			message += "- " + key + ": " + element.GetHelp() + "\n"
		}
	}

	ctx.Reply(message)
}
