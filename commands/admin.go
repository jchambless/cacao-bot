package commands

import (
	"log"

	"github.com/jchambless/cacao/framework"
)

func BanPlayerCommand(ctx framework.Context) {
	log.Println("Ban player command was called")
}

func BanIPCommand(ctx framework.Context) {
	log.Println("Ban IP command was called.")
}

func BanListCommand(ctx framework.Context) {
	log.Println("Ban List command was called.")
}

func DeopCommand(ctx framework.Context) {
	log.Println("Deop command was called.")
}

func OpCommand(ctx framework.Context) {
	log.Println("OP command was called.")
}
