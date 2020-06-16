package commands

import (
	"log"
	"strings"

	"github.com/jchambless/cacao/framework"
	"github.com/jchambless/cacao/util"
)

func BanPlayerCommand(ctx framework.Context) {
	log.Printf("Ban player command was called with args %q", ctx.Args)

	usage := "Usage: !mc ban <player> [reason...]"
	if len(ctx.Args) <= 0 {
		ctx.Reply(usage)
		return
	}

	player := ctx.Args[0]
	reason := make([]string, 0)

	if len(ctx.Args) > 1 {
		reason = ctx.Args[1:]
	}

	rconCommand := "/ban " + player + " " + strings.Join(reason, " ")

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
