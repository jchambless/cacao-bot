package commands

import (
	"log"

	"github.com/jchambless/cacao/framework"
)

func ShutdownCommand(ctx framework.Context) {
	log.Println("Shutdown command was called.")
}
