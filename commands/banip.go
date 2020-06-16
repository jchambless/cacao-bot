package commands

import (
	"log"
	"strconv"
	"strings"

	"github.com/jchambless/cacao/framework"

	"github.com/Tnze/go-mc/net"
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

	port := strconv.Itoa(ctx.Conf.RconPort)
	conn, err := net.DialRCON(ctx.Conf.ServerIP + ":" + port, ctx.Conf.RconPassword)
	if err != nil {
		log.Println("Could not connect to Minecraft server,", err)
		return
	}

	err = conn.Cmd(rconCommand)
	if err != nil {
		log.Println("Command failed to ban ip-address,", err)
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