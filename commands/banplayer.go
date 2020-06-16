package commands

import (
	"log"
	"strconv"
	"strings"

	"github.com/jchambless/cacao/framework"

	"github.com/Tnze/go-mc/net"
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

	port := strconv.Itoa(ctx.Conf.RconPort)
	conn, err := net.DialRCON(ctx.Conf.ServerIP+":"+port, ctx.Conf.RconPassword)
	if err != nil {
		log.Println("Cound not connect to Minecraft server,", err)
		return
	}

	err = conn.Cmd(rconCommand)
	if err != nil {
		log.Println("Command failed to ban player,", err)
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
