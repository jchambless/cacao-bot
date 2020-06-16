package commands

import (
	"log"
	"strconv"
	"strings"

	"github.com/jchambless/cacao/framework"

	"github.com/Tnze/go-mc/net"
)

func DeOperatorCommand(ctx framework.Context) {
	log.Printf("DeOperator command was called with args %q", ctx.Args)

	usage := "Usage: !mc deop <player>"
	if len(ctx.Args) <= 0 || len(ctx.Args) > 1 {
		ctx.Reply(usage)
		return
	}

	player := ctx.Args[0]
	rconCommand := "/deop " + player

	port := strconv.Itoa(ctx.Conf.RconPort)
	conn, err := net.DialRCON(ctx.Conf.ServerIP+":"+port, ctx.Conf.RconPassword)
	if err != nil {
		log.Println("Could not connect to Minecraft server,", err)
		return
	}

	err = conn.Cmd(rconCommand)
	if err != nil {
		log.Println("Command failed for de-assigning operator status to player,", err)
		return
	}

	resp, err := conn.Resp()
	if err != nil {
		log.Println("Command response was nil,", err)
		return
	}
	log.Printf("Server response: %q", resp)

	if strings.Contains(resp, "Usage") {
		ctx.Reply(usage)
		return
	}

	ctx.Reply(resp)
}
