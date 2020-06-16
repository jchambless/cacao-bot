package commands

import (
	"log"
	"strings"

	"github.com/jchambless/cacao/framework"

	"github.com/jchambless/cacao/util"
)

func OperatorCommand(ctx framework.Context) {
	log.Printf("Operator command was called with args %q", ctx.Args)

	usage := "Usage: !mc op <player>"
	if len(ctx.Args) <= 0 || len(ctx.Args) > 1 {
		ctx.Reply(usage)
		return
	}

	player := ctx.Args[0]
	rconCommand := "/op " + player

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
