package commands

import (
	"log"
	"strings"

	"github.com/jchambless/cacao/framework"
	"github.com/jchambless/cacao/util"
)

func BanIPCommand(ctx framework.Context) {
	log.Printf("Ban IP command was called with args %q", ctx.Args)

	usage := "Usage: !mc banip <ip-address>"
	if len(ctx.Args) <= 0 {
		ctx.Reply(usage)
		return
	}

	ipaddress := ctx.Args[0]
	rconCommand := "/ban-ip " + ipaddress

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
