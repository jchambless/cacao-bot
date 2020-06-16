package commands

import (
	"log"
	"strconv"
	"strings"

	"github.com/jchambless/cacao/framework"

	"github.com/Tnze/go-mc/net"
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

	port := strconv.Itoa(ctx.Conf.RconPort)
	conn, err := net.DialRCON(ctx.Conf.ServerIP+":"+port, ctx.Conf.RconPassword)
	if err != nil {
		log.Println("Cound not connect to Minecraft server,", err)
		return
	}

	err = conn.Cmd(rconCommand)
	if err != nil {
		log.Println("Command failed to pardon player or ip,", err)
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
