package commands

import (
	"log"

	"github.com/jchambless/cacao/framework"
)

func HelpCommand(ctx framework.Context) {
	log.Println("Help command was called.")

	ctx.Reply(ctx.Conf.BotName + " running v " + ctx.Conf.Version)
}
