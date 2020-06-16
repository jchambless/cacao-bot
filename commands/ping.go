package commands

import (
	"fmt"
	"log"

	"github.com/jchambless/cacao/framework"

	"github.com/jchambless/cacao/util"
)

func PingCommand(ctx framework.Context) {
	log.Printf("Ping command was called with args %q", ctx.Args)

	usage := "Usage: !mc ping"
	if len(ctx.Args) > 0 {
		ctx.Reply(usage)
		return
	}

	checker := util.MCServerChecker(ctx)
	msg := fmt.Sprintf("**Server**: %s\n**Online**: No\n", ctx.Conf.ServerIP)

	if checker.Online {
		msg = fmt.Sprintf("**Server**: %s\n**Online**: Yes\n**Version**: %s\n**Players**: %s out of %s players\n", ctx.Conf.ServerIP,
			checker.Version, checker.CurrentPlayers, checker.MaxPlayers)
		msg += fmt.Sprintf("**Message of the day**: %s\n", checker.Motd)
		msg += fmt.Sprintf("**Latency**: %s\n", checker.Latency)
		ctx.Reply(msg)
	} else {
		ctx.Reply(msg)
	}
}
