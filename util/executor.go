package util

import (
	"log"
	"strconv"

	"github.com/Tnze/go-mc/net"
)

func RconExecutor(server string, serverPort int, password string, command string) (string, error) {
	port := strconv.Itoa(serverPort)
	conn, err := net.DialRCON(server+":"+port, password)
	if err != nil {
		log.Println("Cound not connect to Minecraft server,", err)
		return "", err
	}

	err = conn.Cmd(command)
	if err != nil {
		log.Println("Command "+command+" failed ,", err)
		return "", err
	}

	resp, err := conn.Resp()
	if err != nil {
		return "", err
	}

	return resp, nil
}
