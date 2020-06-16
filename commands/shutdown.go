package commands

import (
	"log"
	"strconv"
	"strings"

	"github.com/Tnze/go-mc/net"
	"github.com/jchambless/cacao/framework"
)

func ShutdownCommand(ctx framework.Context) {
	log.Printf("Shutdown command was called with args %q", ctx.Args)

	usage := "Usage: !mc stop"
	if len(ctx.Args) > 0 {
		ctx.Reply(usage)
		return
	}

	rconCommand := "/stop"

	port := strconv.Itoa(ctx.Conf.RconPort)
	conn, err := net.DialRCON(ctx.Conf.ServerIP+":"+port, ctx.Conf.RconPassword)
	if err != nil {
		log.Println("Cound not connect to Minecraft server,", err)
		return
	}

	err = conn.Cmd(rconCommand)
	if err != nil {
		log.Println("Command failed to shutdown server,", err)
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
