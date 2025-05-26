package commands

import (
	"log"

	"github.com/jchambless/cacao/framework"
	"github.com/jchambless/cacao/util"
)

func PlayerListCommand(ctx framework.Context) {
	log.Println("Player list command was called.")

	rconCommand := "/list"

	resp, err := util.RconExecutor(ctx.Conf.ServerIP, ctx.Conf.RconPort, ctx.Conf.RconPassword, rconCommand)
	if err != nil {
		return
	}

	log.Printf("Server response: %q", resp)

	ctx.Reply(resp)
}
