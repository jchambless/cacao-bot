package commands

import (
	"log"
	"strings"

	"github.com/jchambless/cacao/framework"
	"github.com/jchambless/cacao/util"
)

func PardonPlayerCommand(ctx framework.Context) {
	log.Printf("Pardon player command was called with args %q", ctx.Args)

	usage := "Usage: !mc pardon <player>|<ip>"
	if len(ctx.Args) <= 0 || len(ctx.Args) > 1 {
		ctx.Reply(usage)
		return
	}

	// Should probably add validation here to see if IP address is valid if the user
	// passed an IP address to the command
	player := ctx.Args[0]
	rconCommand := "/pardon " + player

	resp, err := util.RconExecutor(ctx.Conf.ServerIP, ctx.Conf.RconPort, ctx.Conf.RconPassword, rconCommand)
	if err != nil {
		return
	}

	log.Printf("Server response: %q", resp)

	if strings.Contains(resp, "Usage") {
		ctx.Reply(usage)
		return
	}

	ctx.Reply(resp)
}
