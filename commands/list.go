package commands

import (
	"log"
	"strconv"

	"github.com/jchambless/cacao/framework"

	"github.com/Tnze/go-mc/net"
)

func PlayerListCommand(ctx framework.Context) {
	log.Println("Player list command was called.")

	port := strconv.Itoa(ctx.Conf.RconPort)
	conn, err := net.DialRCON(ctx.Conf.ServerIP+":"+port, ctx.Conf.RconPassword)
	if err != nil {
		log.Println("Could not connect to Minecraft server,", err)
		return
	}

	err = conn.Cmd("/list")
	if err != nil {
		log.Println("Command failed for Player list,", err)
		return
	}

	resp, err := conn.Resp()
	if err != nil {
		log.Println("Command response was nil,", err)
		return
	}
	log.Printf("Server response: %q", resp)

	ctx.Reply(resp)
}
